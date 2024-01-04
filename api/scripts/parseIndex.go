package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/davidbmaier/dbm/types"
	"github.com/dlclark/regexp2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Entry struct {
	workID       string
	artistName   string
	birthYear    int
	deathYear    int
	name         string
	creationInfo string
	material     string
	size         string
	owner        string
	source       string
}

var entryRegex = regexp2.MustCompile(`^\s*(?P<workID>[^\s]+): (?P<artistName>[^:(]+[^\s(:])(?:(?: \([^\d]*(?P<birthYear>\d+)(?:,[^\d]*)?(?:-(?P<deathYear>\d+)[^\d]*)?\))?:?(?: (?P<name>[^(\.]+(?:\((?!o\.[DJ]\.)[^\d]+\))?))?(?: \((?P<creationInfo>[^\)]*\d+[^\)]*|o\.D\.|o\.J\.)\)(?: \(.*\))?)?\.(?: (?:(?P<materialAlt>[^\d\.]+)\.)| (?:(?P<sizeAlt>[^\d,\.]*[\d,]* x [\d,]*[^\.]*)\.)| (?:(?P<material>[^\d\.]+)), (?:(?P<size>[^\d\.]*[\d,]* x [\d,]*[^\.]*))\.)?(?: (?P<owner>[^\d\.]+)\.?)?(?: Q: (?P<source>.+))?)?$`, regexp2.RE2)
var env map[string]string

func findNamedMatches(regex *regexp2.Regexp, str string) map[string]string {
	match, err := regex.FindStringMatch(str)
	check(err)

	results := map[string]string{}
	if match != nil {
		groups := match.Groups()
		for _, group := range groups {
			if group.Name != "0" {
				results[group.Name] = group.String()
			}
		}
	}

	return results
}

func main() {
	var err error
	env, err = godotenv.Read("../.env")
	check(err)

	db, err := gorm.Open(postgres.Open(env["DATABASE_STRING"]), &gorm.Config{})
	check(err)
	db.AutoMigrate(&types.Artist{})
	db.AutoMigrate(&types.Work{})

	data, err := os.ReadFile("./index.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	errorLines := []string{}
	lastWorkIDName := ""
	for _, line := range lines {
		matches := findNamedMatches(entryRegex, line)

		numberWorkIDRegex := regexp.MustCompile(`^\s*\d+[a-z]`)
		if len(matches) == 0 {
			errorLines = append(errorLines, line)

			// if line is not empty and the workID is not just a sub-entry (whitespace + number + letter), update the lastWorkIDName for future potential sub-entries
			if len(line) > 0 && !numberWorkIDRegex.Match([]byte(line)) {
				line = strings.Replace(line, "’", "'", -1) // replace different apostrophes in IDs
				lastWorkIDName = regexp.MustCompile(`^[^\d:]+`).FindString(line)
			}
		} else {
			entry := Entry{}

			matches["workID"] = strings.Replace(matches["workID"], "’", "'", -1) // replace different apostrophes in IDs

			if numberWorkIDRegex.MatchString(matches["workID"]) {
				entry.workID = lastWorkIDName + matches["workID"]
			} else {
				entry.workID = matches["workID"]
				lastWorkIDName = regexp.MustCompile(`^[^\d]+`).FindString(matches["workID"])
			}

			// go through possible fields
			if matches["artistName"] != "" {
				entry.artistName = matches["artistName"]
			}
			if matches["birthYear"] != "" {
				entry.birthYear, err = strconv.Atoi(matches["birthYear"])
				check(err)
			}
			if matches["deathYear"] != "" {
				entry.deathYear, err = strconv.Atoi(matches["deathYear"])
				check(err)
			}
			if matches["name"] != "" {
				entry.name = matches["name"]
			}
			if matches["creationInfo"] != "" {
				entry.creationInfo = matches["creationInfo"]
			}
			if matches["material"] != "" {
				entry.material = matches["material"]
			}
			if matches["size"] != "" {
				cleanSize := strings.Replace(matches["size"], "x", "×", -1)
				entry.size = cleanSize
			}
			if matches["sizeAlt"] != "" && matches["materialAlt"] != "" {
				entry.size = matches["sizeAlt"]
				entry.material = matches["materialAlt"]
			}
			if matches["owner"] != "" {
				entry.owner = matches["owner"]
			}
			if matches["source"] != "" {
				entry.source = matches["source"]
			}

			handleParsedEntry(entry, db)
		}
	}

	fmt.Print("Errors: ", len(errorLines), "/", len(lines))
	os.WriteFile("indexErrors.txt", []byte(strings.Join(errorLines, "\n")), 0777)
}

func handleParsedEntry(entry Entry, db *gorm.DB) {
	artist := types.Artist{
		Name:      entry.artistName,
		BirthYear: &entry.birthYear,
		DeathYear: &entry.deathYear,
	}
	db.Where(types.Artist{Name: entry.artistName}).FirstOrCreate(&artist)

	if (artist.BirthYear != &entry.birthYear && entry.birthYear != 0) || (artist.DeathYear != &entry.deathYear && entry.deathYear != 0) {
		// birth/death have to be updated
		db.Model(&artist).Updates(types.Artist{BirthYear: &entry.birthYear, DeathYear: &entry.deathYear})
	}

	work := types.Work{
		WorkID:       entry.workID,
		Name:         &entry.name,
		CreationInfo: &entry.creationInfo,
		Material:     &entry.material,
		Size:         &entry.size,
		Owner:        &entry.owner,
		Source:       &entry.source,
		Artist:       artist,
	}

	db.Create(&work)
}

// TODO: go through artist names and deduplicate - current progress: Salvador Dalí
// TODO: go through duplicate workIDs as logged by the script's output

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var entryRegex = regexp.MustCompile(`^\s*(?P<workID>[^\s]+): (?P<artistName>[^:(]+[^\s(:])(?:(?: \([^\d]*(?P<birthYear>\d+)(?:,[^\d]*)?(?:-(?P<deathYear>\d+)[^\d]*)?\))?:?(?: (?P<name>[^(\.]+))?(?: \((?P<creationInfo>[^\)]+)\)(?: \(.*\))?)?\.(?: (?:(?P<materialAlt>[^\d\.]+)\.)| (?:(?P<sizeAlt>[\d,]* x [\d,]*[^\.]*)\.)| (?:(?P<material>[^\d\.]+)), (?:(?P<size>[\d,]* x [\d,]*[^\.]*))\.)?(?: (?P<owner>[^\d\.]+)\.)?(?: Q: (?P<source>.+))?)?$`)

func findNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)

	results := map[string]string{}
	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}
	return results
}

func main() {
	data, err := os.ReadFile("./index.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	errorCounter := 0
	for _, line := range lines {
		matches := findNamedMatches(entryRegex, line)
		//fmt.Print(line, "\n")
		var elements = []string{
			"workID", "artistName", "birthYear", "deathYear", "name", "creationInfo", "material", "size", "materialAlt", "sizeAlt", "owner", "source",
		}
		for _, element := range elements {
			if matches[element] != "" {
				//fmt.Print(element, ": ", matches[element], "\n")
			}
		}
		if len(matches) == 0 {
			//fmt.Print("ERROR!\n\n")
			fmt.Print(line, "\n")
			errorCounter++
		}
		//fmt.Print("-----\n")
	}

	fmt.Print("Errors: ", errorCounter, "/", len(lines))
}

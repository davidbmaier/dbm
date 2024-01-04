package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/davidbmaier/dbm/types"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var env map[string]string

func main() {
	var err error
	env, err = godotenv.Read("../../.env")
	check(err)

	db, err := gorm.Open(postgres.Open(env["DATABASE_STRING"]), &gorm.Config{})
	check(err)

	works := []types.Work{}
	db.Model(&types.Work{}).Find(&works)

	errorLines := []string{}

	for _, work := range works {

		firstCharacter, _ := utf8.DecodeRuneInString(work.WorkID)

		switch firstCharacter {
		case 'Ä':
			firstCharacter = 'A'
		case 'Ö':
			firstCharacter = 'O'
		case 'Ü':
			firstCharacter = 'U'
		}

		possiblePaths := []string{
			env["FILES_PATH"] + "/" + string(firstCharacter) + "/" + work.WorkID + ".jpg",
			env["FILES_PATH"] + "/" + string(firstCharacter) + "/" + work.WorkID + ".JPG",
		}

		found := false
		for _, path := range possiblePaths {
			_, err = os.Open(path)
			if err == nil {
				found = true
			}
		}

		if !found {
			errorLines = append(errorLines, work.WorkID)
			fmt.Print(work.WorkID, "\n")
		}
	}

	fmt.Print("Errors: ", len(errorLines), "/", len(works))
	os.WriteFile("missingImages.txt", []byte(strings.Join(errorLines, "\n")), 0777)
}

// TODO: go through works and resolve issues

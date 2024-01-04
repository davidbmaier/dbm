package routes

import (
	"net/url"
	"os"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
)

func GetFile(c *fiber.Ctx, env map[string]string) error {
	workID := c.Params("workID", "")

	decodedWorkID, err := url.PathUnescape(workID)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	firstCharacter, _ := utf8.DecodeRuneInString(decodedWorkID)

	switch firstCharacter {
	case 'Ä':
		firstCharacter = 'A'
	case 'Ö':
		firstCharacter = 'O'
	case 'Ü':
		firstCharacter = 'U'
	}

	possiblePaths := []string{
		env["FILES_PATH"] + "/" + string(firstCharacter) + "/" + decodedWorkID + ".jpg",
		env["FILES_PATH"] + "/" + string(firstCharacter) + "/" + decodedWorkID + ".JPG",
	}

	for _, path := range possiblePaths {
		_, err = os.Open(path)
		if err == nil {
			c.Set("Cache-Control", "max-age=604800")
			return c.SendFile(path)
		}
	}

	// return 404 if we haven't found a valid path
	return c.SendStatus(fiber.StatusNotFound)
}

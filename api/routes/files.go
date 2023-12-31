package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetFile(c *fiber.Ctx, env map[string]string) error {
	workID := c.Params("workID", "")
	firstCharacter := string(workID[0])
	path := env["FILES_PATH"] + "/" + firstCharacter + "/" + workID + ".jpg"

	_, err := os.Open(path)
	if err != nil {
		return c.SendStatus(404)
	}

	return c.SendFile(path)
}

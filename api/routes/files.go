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
		c.Set("Content-Type", "application/json")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Not found",
		})
	}

	return c.SendFile(path)
}

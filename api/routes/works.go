package routes

import (
	database "github.com/davidbmaier/dbm/db"
	"github.com/gofiber/fiber/v2"
)

func GetWorks(c *fiber.Ctx, env map[string]string) error {
	search := c.Query("search")
	page := c.QueryInt("page")
	pageSize := c.QueryInt("pageSize")
	artistID := c.QueryInt("artistID")

	if pageSize == 0 {
		pageSize = 20 // default size
	} else if pageSize < 10 {
		pageSize = 10
	} else if pageSize > 50 {
		pageSize = 50
	}

	works := database.RetrieveWorks(search, artistID, page, pageSize)

	c.Set("Content-Type", "application/json")
	return c.JSON(works)
}

func GetWork(c *fiber.Ctx, env map[string]string) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	work := database.RetrieveWork(id)

	c.Set("Content-Type", "application/json")
	return c.JSON(work)
}

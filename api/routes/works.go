package routes

import (
	database "github.com/davidbmaier/dbm/db"
	"github.com/gofiber/fiber/v2"
)

func GetWorks(c *fiber.Ctx, env map[string]string) error {
	search := c.Query("search")
	page := c.QueryInt("page")
	artistID := c.QueryInt("artistID")

	works := database.RetrieveWorks(search, artistID, page)

	c.Set("Content-Type", "application/json")
	return c.JSON(works)
}

func GetWork(c *fiber.Ctx, env map[string]string) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	work := database.RetrieveWork(id)

	c.Set("Content-Type", "application/json")
	return c.JSON(work)
}

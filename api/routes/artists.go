package routes

import (
	database "github.com/davidbmaier/dbm/db"
	"github.com/gofiber/fiber/v2"
)

func GetArtists(c *fiber.Ctx, env map[string]string) error {
	search := c.Query("search")
	page := c.QueryInt("page")
	pageSize := c.QueryInt("pageSize")

	if pageSize == 0 {
		pageSize = 20 // default size
	} else if pageSize < 10 {
		pageSize = 10
	} else if pageSize > 50 {
		pageSize = 50
	}

	artists := database.RetrieveArtists(search, page, pageSize)

	c.Set("Content-Type", "application/json")
	return c.JSON(artists)
}

func GetArtist(c *fiber.Ctx, env map[string]string) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	artist := database.RetrieveArtist(id)

	c.Set("Content-Type", "application/json")
	return c.JSON(artist)
}

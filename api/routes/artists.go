package routes

import (
	database "github.com/davidbmaier/dbm/db"
	"github.com/gofiber/fiber/v2"
)

func GetArtists(c *fiber.Ctx, env map[string]string) error {
	search := c.Query("search")
	page := c.QueryInt("page")

	artists := database.RetrieveArtists(search, page)

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

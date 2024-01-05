package routes

import (
	database "github.com/davidbmaier/dbm/db"
	"github.com/davidbmaier/dbm/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx, env map[string]string) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	admin := claims["admin"].(bool)
	if !admin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}

	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Admin    bool   `json:"admin"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	database.CreateUser(types.User{
		Username:       payload.Username,
		HashedPassword: string(hashedPassword),
		Admin:          payload.Admin,
	})

	c.Set("Content-Type", "application/json")
	return c.JSON(fiber.Map{"success": true})
}

func UpdateUserPassword(c *fiber.Ctx, env map[string]string) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	payload := struct {
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	if len([]byte(payload.Password)) > 72 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request, password too long",
		})
	}

	if len(payload.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request, password too short",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	database.UpdateUserPassword(types.User{
		Username:       username,
		HashedPassword: string(hashedPassword),
	})
	c.Set("Content-Type", "application/json")
	return c.JSON(fiber.Map{"success": true})
}

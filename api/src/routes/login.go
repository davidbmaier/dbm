package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx, env map[string]string) error {
	payload := struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	// Throws Unauthorized error
	if !(payload.User == env["ADMIN_USER"] && payload.Password == env["ADMIN_PASS"]) && !(payload.User == env["GUEST_USER"] && payload.Password == env["GUEST_PASS"]) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(env["JWT_SECRET"]))
	if err != nil {
		log.Warn(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	c.Set("Content-Type", "application/json")
	cookieHeader := "token=" + t + "; HttpOnly; Max-Age=259200; Path=/"
	if env["DEPLOY_MODE"] == "prod" {
		// only add Secure to cookie if this is production
		cookieHeader = cookieHeader + "; Secure"
	}
	c.Set("Set-Cookie", cookieHeader)
	return c.JSON(fiber.Map{"token": t})
}

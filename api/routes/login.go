package routes

import (
	"time"

	database "github.com/davidbmaier/dbm/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx, env map[string]string) error {
	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	user := database.RetrieveUser(payload.Username)

	if user.ID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"admin":    user.Admin,
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
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

func Logout(c *fiber.Ctx, env map[string]string) error {
	cookieHeader := "token=deleted; HttpOnly; Max-Age=-1; Path=/"
	if env["DEPLOY_MODE"] == "prod" {
		// only add Secure to cookie if this is production
		cookieHeader = cookieHeader + "; Secure"
	}
	c.Set("Set-Cookie", cookieHeader)
	return c.SendStatus(200)
}

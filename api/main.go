package main

import (
	"crypto/tls"
	"fmt"

	database "github.com/davidbmaier/dbm/db"
	"github.com/davidbmaier/dbm/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/joho/godotenv"
)

var env map[string]string

func main() {
	var err error
	env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.InitDatabaseConnection(env)

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "",
		AppName:       "DBM API",
	})

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}]: ${locals:requestid} - ${method} ${path} - ${status}\n",
		TimeFormat: "2006/02/01 15:04:05",
		TimeZone:   "Europe/Berlin",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,https://datenbankmaier.de",
		AllowCredentials: true,
	}))

	apiRouter := app.Group("/api")
	// Login route
	apiRouter.Post("/login", func(c *fiber.Ctx) error { return routes.Login(c, env) })
	apiRouter.Post("/logout", func(c *fiber.Ctx) error { return routes.Logout(c, env) })

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(env["JWT_SECRET"])},
		TokenLookup: "cookie:token",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set("Content-Type", "application/json")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	worksRouter := apiRouter.Group("/works")
	worksRouter.Get("/", func(c *fiber.Ctx) error { return routes.GetWorks(c, env) })
	worksRouter.Get("/:id", func(c *fiber.Ctx) error { return routes.GetWork(c, env) })

	artistsRouter := apiRouter.Group("/artists")
	artistsRouter.Get("/", func(c *fiber.Ctx) error { return routes.GetArtists(c, env) })
	artistsRouter.Get("/:id", func(c *fiber.Ctx) error { return routes.GetArtist(c, env) })

	usersRouter := apiRouter.Group("/users")
	usersRouter.Post("/", func(c *fiber.Ctx) error { return routes.CreateUser(c, env) })
	usersRouter.Patch("/", func(c *fiber.Ctx) error { return routes.UpdateUserPassword(c, env) })

	filesRouter := app.Group("/files")
	filesRouter.Get("/:workID", func(c *fiber.Ctx) error { return routes.GetFile(c, env) })

	if env["DEPLOY_MODE"] == "prod" {
		log.Info("Starting in production mode")
		// create tls certificate
		cer, err := tls.LoadX509KeyPair(env["CERT_PATH"], env["KEY_PATH"])
		if err != nil {
			log.Fatal(err)
		}

		config := &tls.Config{Certificates: []tls.Certificate{cer}}
		// create custom listener
		ln, err := tls.Listen("tcp", fmt.Sprintf(":%s", env["PORT"]), config)
		if err != nil {
			panic(err)
		}

		app.Listener(ln)
	} else {
		log.Info("Starting in development mode")
		app.Listen(fmt.Sprintf(":%s", env["PORT"]))
	}

}

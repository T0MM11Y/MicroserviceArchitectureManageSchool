package main

import (
	"pengajar/database"
	"pengajar/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db := database.Connect()

	app := fiber.New()
	// Enable CORS for all routes
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.SetupPengajar(app)

	app.Listen(":3006")
}

package main

import (
	"pengumuman/database"
	"pengumuman/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db := database.Connect()

	app := fiber.New()

	// Enable CORS for all routes
	app.Use(cors.New())

	app.Static("/api/pengumuman/uploads/", "./uploads/")

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.SetupPengumuman(app)

	app.Listen(":3003")
}

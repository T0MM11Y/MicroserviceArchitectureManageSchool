package main

import (
	"RegisterSiswa/database"
	"RegisterSiswa/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db := database.Connect()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Static("/api/siswa/uploads/", "./uploads/")

	// Setup siswa routes
	routes.SetupSiswa(app)

	app.Listen(":3008")
}

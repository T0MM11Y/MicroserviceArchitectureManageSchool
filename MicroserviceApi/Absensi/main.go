package main

import (
	"Absensi/database"
	"Absensi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))

	// Setup absensi routes
	routes.SetupAbsensi(app)

	app.Listen(":2004")
}

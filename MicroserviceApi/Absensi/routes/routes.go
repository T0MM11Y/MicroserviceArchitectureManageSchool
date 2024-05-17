package routes

import (
	"Absensi/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAbsensi(app *fiber.App) {
	app.Get("/api/absensi", controllers.GetAllAbsensi)
	app.Post("/api/absensi", controllers.NewAbsensi)
	app.Get("/api/absensi/:id", controllers.GetAbsensi)
	app.Put("/api/absensi/:id", controllers.UpdateAbsensi)
	app.Delete("/api/absensi/:id", controllers.DeleteAbsensi)
	app.Get("/api/absensi/user/:id", controllers.GetAbsensiByUser)
	app.Get("/api/absensi/history/:id", controllers.GetAbsensiHistoryByUser)
	app.Get("/api/absensi/user/:id/hariini", controllers.GetAbsensiByUserhariini)
}

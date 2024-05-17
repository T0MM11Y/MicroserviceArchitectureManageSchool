package routes

import (
	"kelas/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupKelas(app *fiber.App) {

	app.Post("/api/kelas", controllers.CreateKelas)
	app.Get("/api/kelas", controllers.GetAllKelas)
	app.Get("/api/kelas/:id", controllers.GetKelasById)
	app.Put("/api/kelas/:id", controllers.UpdateKelas)
	app.Delete("/api/kelas/:id", controllers.DeleteKelas)

}

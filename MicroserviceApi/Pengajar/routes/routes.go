package routes

import (
	"pengajar/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPengajar(app *fiber.App) {

	app.Get("/api/pengajar/:id", controllers.GetPengajarById)
	app.Get("/api/pengajar", controllers.GetAllPengajar)
	app.Post("/api/pengajar", controllers.CreatePengajar)
	app.Put("/api/pengajar/:id", controllers.UpdatePengajar)
	app.Delete("/api/pengajar/:id", controllers.DeletePengajar)

}

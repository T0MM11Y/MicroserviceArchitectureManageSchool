package routes

import (
	"pengumuman/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPengumuman(app *fiber.App) {

	app.Post("/api/pengumuman", controllers.CreatePengumuman)
	app.Get("/api/pengumuman", controllers.GetAllPengumuman)
	app.Get("/api/pengumuman/:id", controllers.GetPengumumanById)
	app.Put("/api/pengumuman/:id", controllers.UpdatePengumuman)
	app.Delete("/api/pengumuman/:id", controllers.DeletePengumuman)

}

package routes

import (
	"RegisterSiswa/controllers"

	"github.com/gofiber/fiber/v2"
)

// SiswaRoutes ...
func SetupSiswa(app *fiber.App) {
	siswaGroup := app.Group("/api/siswa")
	siswaGroup.Get("/", controllers.GetAllSiswa)
	siswaGroup.Post("/register", controllers.RegisterSiswa)
	siswaGroup.Get("/:id", controllers.GetSiswaById)
	siswaGroup.Put("/:id", controllers.UpdateSiswa)
	siswaGroup.Delete("/:id", controllers.DeleteSiswa)
}

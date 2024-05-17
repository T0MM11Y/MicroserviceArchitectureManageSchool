package routes

import (
	"roster/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoster(app *fiber.App) {

	app.Get("/api/roster/:id", controllers.GetRosterById)
	app.Get("/api/roster", controllers.GetAllRosters)
	app.Post("/api/roster", controllers.CreateRoster)
	app.Put("/api/roster/:id", controllers.UpdateRoster)
	app.Delete("/api/roster/:id", controllers.DeleteRoster)

}

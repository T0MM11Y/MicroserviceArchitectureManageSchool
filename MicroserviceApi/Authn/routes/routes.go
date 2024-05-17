package routes

import (
	"Authn/controllers"

	"github.com/gofiber/fiber/v2"
)

// UserRoutes ...
func Setup(app *fiber.App) {
	app.Post("/api/login", controllers.LoginUser)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logoutUser", controllers.LogoutUser)
}

// AdminRoutes ...
func SetupAdmin(app *fiber.App) {
	app.Post("/api/loginAdmin", controllers.LoginAdmin)
	app.Get("/api/admin", controllers.Admin)
	app.Post("/api/logoutAdmin", controllers.LogoutAdmin)
}

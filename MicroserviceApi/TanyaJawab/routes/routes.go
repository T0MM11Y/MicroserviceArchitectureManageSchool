package routes

import (
	"tanyajawab/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupTanyaJawab(app *fiber.App) {
	app.Post("/api/tanyajawab", controllers.AskQuestion)
	app.Get("/api/tanyajawab", controllers.GetQuestions)
	app.Get("/api/tanyajawab/:id", controllers.GetQuestionByID)
	app.Post("/api/tanyajawab/:id", controllers.AnswerQuestion)

}

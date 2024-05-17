package controllers

import (
	"fmt"
	"tanyajawab/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type QuestionRequest struct {
	User       string `json:"User"`
	Pertanyaan string `json:"Pertanyaan"`
}

func AskQuestion(c *fiber.Ctx) error {
	var q QuestionRequest

	if err := c.BodyParser(&q); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	if q.User != "" && q.Pertanyaan != "" {
		db := c.Locals("db").(*gorm.DB)

		var lastQuestion models.TanyaJawab
		result := db.Where("user = ?", q.User).Last(&lastQuestion)

		if result.Error != nil {
			if result.Error != gorm.ErrRecordNotFound {
				// An actual error occurred
				fmt.Println(result.Error)
				return c.Status(500).SendString("Internal Server Error")
			}
		} else {
			tanggalTanya, err := time.Parse("2006-01-02 15:04:05", lastQuestion.TanggalTanya)
			if err != nil {
				fmt.Println(err)
				return c.Status(500).SendString("Internal Server Error")
			}

			if tanggalTanya.Add(2 * time.Hour).After(time.Now()) {
				return c.Status(429).SendString("You can only ask a question once every 2 hours")
			}
		}

		question := models.TanyaJawab{
			Pertanyaan:   q.Pertanyaan,
			User:         q.User,
			TanggalTanya: time.Now().Format("2006-01-02 15:04:05"),
		}

		db.Create(&question)

		return c.Status(200).SendString("Question saved successfully")
	} else {
		return c.Status(400).SendString("Bad Request")
	}
}

func GetQuestions(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var questions []models.TanyaJawab
	db.Find(&questions)

	return c.JSON(questions)
}

func GetQuestionByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := c.Locals("db").(*gorm.DB)

	var question models.TanyaJawab
	result := db.First(&question, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("Question not found")
		} else {
			fmt.Println(result.Error)
			return c.Status(500).SendString("Internal Server Error")
		}
	}

	return c.JSON(question)
}

func AnswerQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	var answer models.TanyaJawab

	if err := c.BodyParser(&answer); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	db := c.Locals("db").(*gorm.DB)

	var question models.TanyaJawab
	result := db.First(&question, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("Question not found")
		} else {
			fmt.Println(result.Error)
			return c.Status(500).SendString("Internal Server Error")
		}
	}

	if question.Jawaban != "" {
		return c.Status(400).SendString("Question has already been answered")
	}

	question.Jawaban = answer.Jawaban
	question.Admin = "admin"
	question.TanggalJawab = time.Now().Format("2006-01-02 15:04:05")
	db.Save(&question)

	return c.Status(200).SendString("Answer saved successfully")
}

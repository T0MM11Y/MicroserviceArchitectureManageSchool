package controllers

import (
	"fmt"
	"path/filepath"
	"pengumuman/database"
	"pengumuman/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreatePengumuman(c *fiber.Ctx) error {
	db := database.DB
	pengumuman := new(models.Pengumuman)

	// Get the fields from the form data
	pengumuman.Judul = c.FormValue("judul")
	pengumuman.Isi = c.FormValue("isi")

	// Get the file from the form data
	file, err := c.FormFile("urlphoto")

	// Initialize the URL to an empty string
	url := ""

	// If a file was uploaded, save it and create the URL
	if err == nil {
		// Save the file to a directory
		dir := "./uploads/"
		dst := filepath.Join(dir, file.Filename)
		if err := c.SaveFile(file, dst); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Could not save file",
			})
		}

		// Create the URL for the file
		url = fmt.Sprintf("http://%s/api/pengumuman/uploads/%s", c.Hostname(), file.Filename)
	}

	// Format the current time to a string
	now := time.Now().Format(time.RFC3339)

	pengumuman.Created_At = now
	pengumuman.Urlphoto = url // Use the file URL here

	db.Create(
		&pengumuman,
	)
	return c.JSON(fiber.Map{
		"message": "Successfully created announcement",
		"data":    pengumuman,
	})
}

func GetAllPengumuman(c *fiber.Ctx) error {
	db := database.DB
	var pengumuman []models.Pengumuman
	db.Find(&pengumuman)
	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    pengumuman,
	})
}

func GetPengumumanById(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var pengumuman models.Pengumuman
	db.First(&pengumuman, id)
	if pengumuman.Judul == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Pengumuman not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    pengumuman,
	})
}

func UpdatePengumuman(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var pengumuman models.Pengumuman
	db.First(&pengumuman, id)
	if pengumuman.Judul == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Pengumuman not found",
		})
	}
	pengumuman.Judul = c.FormValue("judul")
	pengumuman.Isi = c.FormValue("isi")
	file, err := c.FormFile("urlphoto")
	if err == nil {
		dir := "./uploads/"
		dst := filepath.Join(dir, file.Filename)
		if err := c.SaveFile(file, dst); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Could not save file",
			})
		}
		url := fmt.Sprintf("http://%s/api/pengumuman/uploads/%s", c.Hostname(), file.Filename)
		pengumuman.Urlphoto = url
	}
	pengumuman.Updated_At = time.Now().Format(time.RFC3339)
	db.Save(&pengumuman)
	return c.JSON(fiber.Map{
		"message": "Successfully updated announcement",
		"data":    pengumuman,
	})
}

func DeletePengumuman(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var pengumuman models.Pengumuman
	db.First(&pengumuman, id)
	if pengumuman.Judul == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Pengumuman not found",
		})
	}
	db.Delete(&pengumuman)
	return c.JSON(fiber.Map{
		"message": "Successfully deleted announcement",
	})
}

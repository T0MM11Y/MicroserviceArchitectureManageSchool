package controllers

import (
	"fmt"
	"roster/database"
	"strconv"

	"roster/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllRosters(c *fiber.Ctx) error {
	db := database.DB
	var rosters []models.Roster

	db.Preload("Pengajar").Find(&rosters)
	if len(rosters) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No rosters found",
		})
	}
	return c.JSON(rosters)
}

func GetRosterById(c *fiber.Ctx) error {
	db := database.DB
	var roster models.Roster

	rosterID, _ := strconv.Atoi(c.Params("id"))    // Get roster ID from route parameters
	db.Preload("Pengajar").Find(&roster, rosterID) // Find roster by ID

	if roster.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No roster found with given ID",
		})
	}

	return c.JSON(roster)

}
func CreateRoster(c *fiber.Ctx) error {
	db := database.DB
	roster := new(models.Roster)

	if err := c.BodyParser(roster); err != nil {
		fmt.Println(err) // Print the error to the console
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	result := db.Create(&roster)
	if result.Error != nil {
		fmt.Println(result.Error) // Print the error to the console
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create roster",
		})
	}

	return c.JSON(roster)
}
func UpdateRoster(c *fiber.Ctx) error {
	db := database.DB
	roster := new(models.Roster)

	rosterID, _ := strconv.Atoi(c.Params("id")) // Get roster ID from route parameters
	db.Find(&roster, rosterID)                  // Find roster by ID

	if roster.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No roster found with given ID",
		})
	}

	if err := c.BodyParser(roster); err != nil {
		fmt.Println(err) // Print the error to the console
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	result := db.Save(&roster)
	if result.Error != nil {
		fmt.Println(result.Error) // Print the error to the console
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update roster",
		})
	}

	return c.JSON(roster)
}

func DeleteRoster(c *fiber.Ctx) error {
	db := database.DB
	var roster models.Roster
	id := c.Params("id")

	db.Find(&roster, id)
	if roster.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No roster found with given ID",
		})
	}

	db.Delete(&roster)
	return c.JSON(fiber.Map{
		"message": "Roster deleted successfully",
	})
}

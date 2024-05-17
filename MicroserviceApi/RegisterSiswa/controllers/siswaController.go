package controllers

import (
	"RegisterSiswa/database"
	"RegisterSiswa/models"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetSiswaByNISN(c *fiber.Ctx) error {
	db := database.DB
	var siswa models.User
	NISN := c.Params("NISN")

	db.Find(&siswa, "NISN = ?", NISN)
	if siswa.NISN == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No student found with given NISN",
		})
	}
	return c.JSON(siswa)
}
func GetAllSiswa(c *fiber.Ctx) error {

	db := database.DB
	var siswas []models.User
	db.Find(&siswas)

	if len(siswas) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No students found",
		})
	}
	return c.JSON(siswas)
}
func UpdateSiswa(c *fiber.Ctx) error {
	db := database.DB
	var siswa models.User
	id := c.Params("id")

	db.Find(&siswa, id)
	if siswa.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No student found with given ID",
		})
	}

	// Update the student's data
	siswa.Nama_Depan = c.FormValue("Nama_Depan")
	siswa.Nama_Belakang = c.FormValue("Nama_Belakang")
	siswa.Kelas = c.FormValue("Kelas")
	siswa.Agama = c.FormValue("Agama")
	siswa.Email = c.FormValue("Email")
	siswa.Alamat = c.FormValue("Alamat")
	siswa.Tanggal_Lahir = c.FormValue("Tanggal_Lahir")
	siswa.Jenis_Kelamin = c.FormValue("Jenis_Kelamin")

	result := db.Save(&siswa)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error when saving to database",
			"error":   result.Error,
		})
	}

	return c.JSON(siswa)
}

func DeleteSiswa(c *fiber.Ctx) error {
	db := database.DB
	var siswa models.User
	id := c.Params("id")

	db.Find(&siswa, id)
	if siswa.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No student found with given ID",
		})
	}

	db.Delete(&siswa)

	return c.JSON(fiber.Map{
		"message": "Student deleted successfully",
	})
}
func RegisterSiswa(c *fiber.Ctx) error {

	password, _ := bcrypt.GenerateFromPassword([]byte("siswa123"), 14)
	nisn, err := strconv.ParseInt(c.FormValue("NISN"), 10, 32)
	fmt.Println("NISN:", c.FormValue("NISN"))
	fmt.Println("Error:", err)
	if err != nil {
		// Handle parsing error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid NISN format",
		})
	}
	var existingUser models.User
	database.DB.Where("NISN = ? OR Email = ?", nisn, c.FormValue("Email")).First(&existingUser, &models.User{})
	if existingUser.NISN != 0 {
		// User dengan NISN atau email yang sama sudah ada
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "NISN or email already in use",
		})
	}

	user := models.User{
		NISN:          int(nisn),
		Nama_Depan:    c.FormValue("Nama_Depan"),
		Nama_Belakang: c.FormValue("Nama_Belakang"),
		Kelas:         c.FormValue("Kelas"),
		Agama:         c.FormValue("Agama"),
		Email:         c.FormValue("Email"),
		Alamat:        c.FormValue("Alamat"),
		Tanggal_Lahir: c.FormValue("Tanggal_Lahir"),
		Jenis_Kelamin: c.FormValue("Jenis_Kelamin"),
		Password:      password,
	}

	// Save the user to the database
	result := database.DB.Create(&user)
	if result.Error != nil {
		// Log the error
		log.Println("Failed to save user to database:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Could not save user to database",
		})
	}

	log.Println("Saved user to database:", user)

	// Return appropriate response, for example, user ID
	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"toast":   "Siswa telah berhasil ditambahkan",
	})

}

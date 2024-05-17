package controllers

import (
	"fmt"
	"pengajar/database"
	"pengajar/models"
	"strconv"
	"unicode"

	"github.com/gofiber/fiber/v2"
)

func GetPengajarById(c *fiber.Ctx) error {
	db := database.DB
	var pengajar models.Pengajar
	id := c.Params("id")

	db.Find(&pengajar, id)
	if pengajar.NIP == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No teacher found with given ID",
		})
	}
	return c.JSON(pengajar)
}

func GetAllPengajar(c *fiber.Ctx) error {
	db := database.DB
	var pengajars []models.Pengajar
	db.Find(&pengajars)

	if len(pengajars) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No teachers found",
		})
	}
	return c.JSON(pengajars)
}

func UpdatePengajar(c *fiber.Ctx) error {
	db := database.DB
	var pengajar models.Pengajar
	id := c.Params("id")

	db.Find(&pengajar, id)
	if pengajar.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No teacher found with given ID",
		})
	}

	// Update the teacher's data
	pengajar.Nama_Depan = c.FormValue("Nama_Depan")
	pengajar.Nama_Belakang = c.FormValue("Nama_Belakang")
	pengajar.Bidang = c.FormValue("Bidang")
	pengajar.Agama = c.FormValue("Agama")
	pengajar.Email = c.FormValue("Email")
	pengajar.Alamat = c.FormValue("Alamat")
	pengajar.Tanggal_Lahir = c.FormValue("Tanggal_Lahir")
	pengajar.Jenis_Kelamin = c.FormValue("Jenis_Kelamin")

	db.Save(&pengajar)

	return c.JSON(pengajar)
}

func DeletePengajar(c *fiber.Ctx) error {
	db := database.DB
	var pengajar models.Pengajar
	id := c.Params("id")

	db.Find(&pengajar, id)
	if pengajar.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No teacher found with given ID",
		})
	}

	db.Delete(&pengajar)

	return c.JSON(fiber.Map{
		"message": "Teacher deleted successfully",
	})
}

func CreatePengajar(c *fiber.Ctx) error {

	// Get the other fields from the form data
	nipStr := c.FormValue("NIP")
	if nipStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "NIP is required",
		})
	}

	// Validate that the NIP string only contains digits
	for _, char := range nipStr {
		if !unicode.IsDigit(char) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid NIP format",
			})
		}
	}

	nip, err := strconv.ParseInt(nipStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid NIP format",
		})
	}
	fmt.Println("NIP:", c.FormValue("NIP"))
	fmt.Println("Error:", err)
	if err != nil {
		// Handle parsing error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid NIP format",
		})
	}
	var existingPengajar models.Pengajar
	database.DB.Where("NIP = ? OR Email = ?", nip, c.FormValue("Email")).First(&existingPengajar)
	if existingPengajar.NIP != 0 {
		// User dengan NIP atau email yang sama sudah ada
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "NIP or email already in use",
		})
	}

	pengajar := models.Pengajar{
		NIP:           int(nip),
		Nama_Depan:    c.FormValue("Nama_Depan"),
		Nama_Belakang: c.FormValue("Nama_Belakang"),
		Bidang:        c.FormValue("Bidang"),
		Agama:         c.FormValue("Agama"),
		Email:         c.FormValue("Email"),
		Alamat:        c.FormValue("Alamat"),
		Tanggal_Lahir: c.FormValue("Tanggal_Lahir"),
		Jenis_Kelamin: c.FormValue("Jenis_Kelamin"),
	}
	database.DB.Create(&pengajar)

	// Return appropriate response, for example, user ID
	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"toast":   "Pengajar telah berhasil ditambahkan",
		"user_id": pengajar.NIP,
	})

}

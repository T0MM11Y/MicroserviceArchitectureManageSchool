package database

import (
	"Authn/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/microParmaksianschool"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = connection
	connection.AutoMigrate(&models.Admin{})
	connection.AutoMigrate(&models.User{})

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password")
	}

	// Create a new admin user
	admin := models.Admin{
		Username: "admin",
		Password: hashedPassword,
	}

	// Check if admin user already exists
	var existingAdmin models.Admin
	if err := DB.First(&existingAdmin, "username = ?", admin.Username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Admin user does not exist, create a new one
			DB.Create(&admin)
		} else {
			// An error occurred while checking for the admin user
			panic("Failed to check for existing admin user")
		}
	}
}

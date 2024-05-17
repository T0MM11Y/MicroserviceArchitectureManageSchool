package database

import (
	"tanyajawab/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	connection, err := gorm.Open(mysql.Open("root:@/microparmaksianschool"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = connection
	connection.AutoMigrate(&models.TanyaJawab{})

	return DB
}

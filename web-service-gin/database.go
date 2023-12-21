package database

import (
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"web-service-gin/models"
)

var db *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open(), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Albums{})
}

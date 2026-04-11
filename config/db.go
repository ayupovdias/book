package db

import (
	"book/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=bookstore port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(&models.User{}, &models.Author{}, &models.Book{}, &models.Category{}, &models.FavoriteBook{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = database
	fmt.Println("Database connected and migrated!")
}

package config

import (
	"fmt"
	"log"
	"os"

	"go-rest-api/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set in .env")
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// migration
	err = DB.AutoMigrate(
		&models.User{},
		&models.Post{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database connection established and migration successfully.")
}

package database

import (
	"fmt"
	"log"

	"github.com/pithandev/fuzzy-octo-meme/internal/database/models"
	"github.com/pithandev/fuzzy-octo-meme/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	// 1) load configurations
	cfg := config.Load()
	// 2) build connection string
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// 3) open connection with gorm.Open()
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	// 4) check errors
	if err != nil {
		fmt.Println("Error trying to connect with database.")
	}

	// 5) put in global DB variable
	DB = db
	log.Println("✅ Database connected successfully")

	if err := AutoMigrate(); err != nil {
		return fmt.Errorf("migrations failed: %w", err)
	}

	return nil //to avoid error mensages
}

// migrations
func AutoMigrate() error {

	db := GetDB()

	err := db.AutoMigrate(&models.User{}, &models.Task{})

	if err != nil {
		return fmt.Errorf("migration failed")
	}

	fmt.Println("✅ Database migrations applied successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

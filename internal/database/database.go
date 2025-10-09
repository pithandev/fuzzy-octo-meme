package database

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	// 1) load configurations

	// 2) build connection string

	// 3) open connection with gorm.Open()

	// 4) check errors

	// 5) put in global DB variable

	return nil //to avoid error mensages
}

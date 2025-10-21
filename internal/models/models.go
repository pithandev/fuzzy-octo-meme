package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Task struct {
	gorm.Model
	ID          uint
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

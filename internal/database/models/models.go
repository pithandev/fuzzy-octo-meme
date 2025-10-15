package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Task struct {
	ID          uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Completed   bool   `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

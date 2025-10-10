package models

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Task struct {
	ID          uint
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

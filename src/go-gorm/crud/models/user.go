package models

import (
	"time"
)

type Model struct {
	UUID      string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Model    Model  `gorm:"embedded"`
	Name     string `gorm:"not null"`
	Age      uint8
	Birthday time.Time
	Email    string
	Tel      string `gorm:"column:mobile"`
}

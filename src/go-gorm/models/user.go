package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      uint8
	Birthday time.Time
	Email    string
}

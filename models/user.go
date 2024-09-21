package models

import "gorm.io/gorm"

// Model User
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
}

package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Name string `json:"name"`
}

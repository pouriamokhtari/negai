package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Email          string `gorm:"unique;not null;uniqueIndex"`
	PasswordDigest string `json:"-" gorm:"not null"` // don't include when marshaling
}

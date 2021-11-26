package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Email          string
	PasswordDigest string `json:"-"` // don't include when marshaling
}

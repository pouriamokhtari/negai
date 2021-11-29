package models

import "gorm.io/gorm"

const (
	Member = 0
	Admin  = 1
)

// User model
type User struct {
	gorm.Model
	FullName       string
	Role           byte   `gorm:"not null;default=0"`
	Email          string `gorm:"unique;not null;uniqueIndex"`
	PasswordDigest string `json:"-" gorm:"not null"` // don't include when marshaling
}

func NewRoleFromString(role string) byte {
	switch role {
	case "admin", "Admin", "ADMIN":
		return Admin
	}
	return Member
}

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
	Role           int    `gorm:"not null;default=0"`
	Email          string `gorm:"unique;not null;uniqueIndex"`
	PasswordDigest string `json:"-" gorm:"not null"` // don't include when marshaling
}

func RoleFromString(role string) int {
	switch role {
	case "admin", "Admin", "ADMIN":
		return Admin
	}
	return Member
}

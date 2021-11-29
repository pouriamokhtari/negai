package models

import (
	"negai/database"

	"gorm.io/gorm"
)

const (
	Member = 0
	Admin  = 1
)

// User model
type User struct {
	gorm.Model
	FullName       string
	Role           uint   `gorm:"not null;default=0"`
	Email          string `gorm:"unique;not null;uniqueIndex"`
	PasswordDigest string `json:"-" gorm:"not null"` // don't include when marshaling
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := database.Connection.Find(&users)
	return users, result.Error
}

func (u *User) Find(id uint) error {
	result := database.Connection.First(u, id)
	return result.Error
}

func (u *User) Create() error {
	result := database.Connection.Create(u)
	return result.Error
}

func (u *User) Update(newUser User) error {
	result := database.Connection.Model(&u).Updates(newUser)
	return result.Error
}

func (u *User) Delete() error {
	result := database.Connection.Delete(u)
	return result.Error
}

func RoleFromString(role string) uint {
	switch role {
	case "admin", "Admin", "ADMIN":
		return Admin
	}
	return Member
}

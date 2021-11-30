package models

import (
	"negai/database"
	"negai/helpers"

	"gorm.io/gorm"
)

const (
	Member = "member"
	Admin  = "admin"
)

// User model
type User struct {
	gorm.Model
	FullName       string
	Role           string `gorm:"not null;default=0"`
	Email          string `gorm:"unique;not null;uniqueIndex"`
	PasswordDigest string `json:"-" gorm:"not null"` // don't include when marshaling
	Password       string `json:"-" gorm:"-"`
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := database.Connection.Find(&users)
	return users, result.Error
}

func (u *User) Find(id uint) error {
	return database.Connection.First(u, id).Error
}

func (u *User) Create() error {
	if len(u.Password) != 0 {
		passwordDigest, err := helpers.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.PasswordDigest = passwordDigest
	}
	return database.Connection.Create(u).Error
}

func (u *User) Update(newUser User) error {
	return database.Connection.Model(&u).Updates(newUser).Error

}

func (u *User) Delete() error {
	return database.Connection.Delete(u).Error
}

package models

import "gorm.io/gorm"

// List model
type List struct {
	gorm.Model
	Title string `gorm:"unique;not null;uniqueIndex"`
	Items string `json:"-" gorm:"not null"` // don't include when marshaling
}

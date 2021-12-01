package models

import "negai/database"

type Model interface {
	Find(id uint) error
	Create() error
	Delete() error
	Update(Model) error
}

func AutoMigrateModels() {
	database.Connection.AutoMigrate(&User{})
	database.Connection.AutoMigrate(&List{})
}

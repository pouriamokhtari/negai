package models

import "negai/database"

type Model interface {
	Find()
	Create()
	Delete()
	Update(*Model)
}

func AutoMigrateModels() {
	database.Connection.AutoMigrate(&User{})
	database.Connection.AutoMigrate(&List{})
}

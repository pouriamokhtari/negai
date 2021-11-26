package database

import (
	"negai/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Connection *gorm.DB
)

func Connect() {
	var err error

	Connection, err = gorm.Open(postgres.Open("host=localhost user=postgres password= dbname=negai port=5432 sslmode=disable"))
	if err != nil {
		panic("Couldn't connect to the database")
	}
}

func AutoMigrateModels() {
	Connection.AutoMigrate(&models.User{})
}

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Connection *gorm.DB
)

func Connect(connectionString string) {
	var err error

	Connection, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic("Couldn't connect to the database")
	}
}

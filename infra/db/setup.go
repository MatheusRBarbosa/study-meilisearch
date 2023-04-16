package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"squad10x.com.br/boilerplate/crosscutting"
)

var ctx *gorm.DB

func Context() *gorm.DB {
	return ctx
}

func ConnectDatabase() {
	connectionString := crosscutting.GetConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	ctx = db
}

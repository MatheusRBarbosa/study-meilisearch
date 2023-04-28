package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"squad10x.com.br/boilerplate/crosscutting"
)

var ctx *gorm.DB
var dbLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,        // Don't include params in the SQL log
		Colorful:                  true,        // Disable color
	},
)

func Context() *gorm.DB {
	return ctx
}

func ConnectDatabase() {
	connectionString := crosscutting.GetConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 dbLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	ctx = db
}

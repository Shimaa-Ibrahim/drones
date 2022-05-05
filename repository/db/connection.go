package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ConnectToDB(envDB string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv(envDB)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "drones.",
		},
	})
	return db, err
}

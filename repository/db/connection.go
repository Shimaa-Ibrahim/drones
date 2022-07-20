package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ConnectToDatabase(databaseConnection string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseConnection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "drones.",
		},
	})
	return db, err
}

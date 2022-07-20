package repository

import (
	"fmt"
	"github/Shimaa-Ibrahim/drones/repository/db"
	"log"
	"os"
	"os/exec"

	"gorm.io/gorm"
)

const MAIN_DB = "MAIN_DB"

func InitializeTestDatabase(database string, schema string, dbConnection string) (*gorm.DB, error) {
	dbClient, err := db.ConnectToDatabase(os.Getenv(MAIN_DB))
	if err != nil {
		log.Fatalf("[Error] connection to db: %v", err)
	}
	if err := dbClient.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS test_%s;", database)).Error; err != nil {
		log.Fatalf("[Error] test db drop: %v", err)
	}
	if err = dbClient.Exec(fmt.Sprintf("CREATE DATABASE test_%s;", database)).Error; err != nil {
		log.Fatalf("[Error] test db creation: %v", err)
	}

	dbClient, err = db.ConnectToDatabase(dbConnection)
	if err != nil {
		log.Fatalf("[Error] connection to db: %v", err)
	}
	if err = dbClient.Exec(fmt.Sprintf("CREATE SCHEMA %s", schema)).Error; err != nil {
		log.Fatalf("[Error] schema creation: %v", err)
	}
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("~/go/bin/gorm-goose -env=test -path=./db -pgschema=%s up", schema))
	err = cmd.Run()
	if err != nil {
		log.Fatalf("[Error] migratrion: %v", err)
	}
	return dbClient, nil
}

func TruncateDB(dbClient *gorm.DB, schema string, tables []string) {
	for _, table := range tables {
		dbClient.Exec(fmt.Sprintf(`TRUNCATE TABLE %s.%s CASCADE;`, schema, table))
	}
}

func DestroyTestDataBase(dbClient *gorm.DB, database string) {
	sqlDB, err := dbClient.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
	dbClient.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS test_%s;", database))
}

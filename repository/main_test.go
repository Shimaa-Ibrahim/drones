package repository

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"gorm.io/gorm"
)

const TEST_DRONE_DATABASE = "TEST_DRONE_DATABASE"

var dbClient *gorm.DB

func GenerateRandomText(n int) string {
	const letterBytes = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func clearDataBase() {
	TruncateDB(dbClient, "drones", []string{"drones", "drone_medications", "medications"})
}

func setup(database string, schema string, dbConnection string) *gorm.DB {
	dbClient, err := InitializeTestDatabase(database, schema, dbConnection)
	if err != nil {
		log.Fatalf("[Error] connection to db: %v", err)
	}
	return dbClient
}

func shutdown(dbClient *gorm.DB, database string) {
	fmt.Println("destroy test database ...")
	DestroyTestDataBase(dbClient, database)
}

func TestMain(m *testing.M) {
	dbClient = setup("drones", "drones", os.Getenv(TEST_DRONE_DATABASE))
	code := m.Run()
	shutdown(dbClient, "drones")
	os.Exit(code)
}

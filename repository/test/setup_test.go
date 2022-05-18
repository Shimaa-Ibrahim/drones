package test

import (
	"fmt"
	"log"
	"math/rand"

	"gorm.io/gorm"
)

const TEST_DRONE_DATABASE = "TEST_DRONE_DATABASE"
const DB_SCHEME = "drones"

func TruncateDB(db *gorm.DB) {
	tables := []string{"drone_states", "drone_models", "drones", "medications", "battery_levels"}
	for _, table := range tables {
		log.Printf("removing all records in scheme %s database table %s ", DB_SCHEME, table)
		db.Exec(fmt.Sprintf(`TRUNCATE TABLE %s.%s CASCADE;`, DB_SCHEME, table))
	}
}

func generateRandomText(n int) string {
	const letterBytes = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

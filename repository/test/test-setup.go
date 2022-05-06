package test

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

const TEST_DRONE_DATABASE = "TEST_DRONE_DATABASE"
const DB_SCHEME = "drones"

func TruncateDB(db *gorm.DB) {
	tables := []string{"drones", "medications", "battery_levels"}
	for _, table := range tables {
		log.Printf("removing all records in scheme %s database table %s ", DB_SCHEME, table)
		db.Exec(fmt.Sprintf(`TRUNCATE TABLE %s.%s CASCADE;`, DB_SCHEME, table))
	}
}

package main

import (
	"github/Shimaa-Ibrahim/drones/repository/entity"

	"gorm.io/gorm"
)

// Up is executed when this migration is applied
func Up_20220717003712(txn *gorm.DB) {
	type BatteryLog struct {
		entity.DBModel
		BatteryLevel float64 `json:"battery"`
		DroneID      int     `json:"drone_id"`
		Drone        entity.Drone
	}
	txn.AutoMigrate(&BatteryLog{})
}

// Down is executed when this migration is rolled back
func Down_20220717003712(txn *gorm.DB) {
	txn.Migrator().DropTable("battery_logs")
}

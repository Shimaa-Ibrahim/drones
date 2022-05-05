package main

import (
	"gorm.io/gorm"
)

// Up is executed when this migration is applied
func Up_20220505181515(txn *gorm.DB) {
	type DroneModel struct {
		ID   uint   `gorm:"primaryKey"`
		Name string `gorm:"unique;not null"`
	}

	type DroneState struct {
		ID   uint   `gorm:"primaryKey"`
		Name string `gorm:"unique;not null"`
	}

	type Medication struct {
		gorm.Model
		Name      string `gorm:"not null"`
		Code      string `gorm:"not null"`
		Weight    string `gorm:"not null"`
		ImagePath string
		DroneID   uint
	}

	type BatteryLevels struct {
		gorm.Model
		DroneID      uint    `gorm:"not null"`
		BatteryLevel float64 `gorm:"not null"`
	}

	type Drone struct {
		gorm.Model
		SerielNumber    string     `gorm:"size:100;unique;not null"`
		DroneModelID    uint       `gorm:"not null"`
		DroneModel      DroneModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		WeightLimit     uint64     `gorm:"not null;check:weight_limit<=500"`
		BatteryCapacity uint64     `gorm:"not null;check:battery_capacity<=100"`
		DroneStateID    uint
		DroneState      DroneState `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		BatteryLevels   []BatteryLevels
		Medications     []Medication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}

	txn.AutoMigrate(&DroneModel{})
	txn.AutoMigrate(&DroneState{})
	txn.AutoMigrate(&Drone{})
	txn.AutoMigrate(&BatteryLevels{})
	txn.AutoMigrate(&Medication{})
}

// Down is executed when this migration is rolled back
func Down_20220505181515(txn *gorm.DB) {
	txn.Migrator().DropTable("drones")
	txn.Migrator().DropTable("drone_models")
	txn.Migrator().DropTable("drone_states")
	txn.Migrator().DropTable("battery_levels")
	txn.Migrator().DropTable("medications")

}

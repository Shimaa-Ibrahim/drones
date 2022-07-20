package main

import (
	"github/Shimaa-Ibrahim/drones/repository/entity"
	"log"

	"gorm.io/gorm"
)

// Up is executed when this migration is applied
func Up_20220711225831(txn *gorm.DB) {
	type Medication struct {
		entity.DBModel
		Name      string
		Weight    float64
		Code      string `gorm:"unique"`
		ImagePath string
	}

	type Drone struct {
		entity.DBModel
		SerialNumber    string `gorm:"unique"`
		Model           string
		WeightLimit     float64
		BatteryCapacity float64
		State           string
		Medications     []Medication `gorm:"many2many:drone_medications;"`
	}

	type DroneMedication struct {
		entity.DBModel
		DroneID      uint `gorm:"primaryKey"`
		MedicationID uint `gorm:"primaryKey"`
		Loaded       bool `gorm:"default:true"`
	}

	err := txn.SetupJoinTable(&Drone{}, "Medications", &DroneMedication{})
	if err != nil {
		log.Fatal(err)
	}
	txn.AutoMigrate(&Drone{}, &Medication{}, &DroneMedication{})
}

// Down is executed when this migration is rolled back
func Down_20220711225831(txn *gorm.DB) {
	txn.Migrator().DropTable("drone_medications")
	txn.Migrator().DropTable("medications")
	txn.Migrator().DropTable("drones")
}

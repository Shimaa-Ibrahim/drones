package entity

import "gorm.io/gorm"

type BatteryLevels struct {
	gorm.Model
	DroneID      uint    `gorm:"not null"`
	BatteryLevel float64 `gorm:"not null"`
}

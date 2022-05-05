package entity

import "gorm.io/gorm"

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

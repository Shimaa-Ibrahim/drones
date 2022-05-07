package entity

import "gorm.io/gorm"

type Drone struct {
	gorm.Model
	SerielNumber    string     `json:"serial_number" gorm:"size:100;unique;not null"`
	DroneModelID    uint       `json:"drone_model_id" gorm:"not null"`
	DroneModel      DroneModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WeightLimit     uint64     `json:"weight_limit" gorm:"not null;check:weight_limit<=500"`
	BatteryCapacity uint64     `json:"battery_capacity" gorm:"not null;check:battery_capacity<=100"`
	DroneStateID    uint       `json:"drone_state_id"`
	DroneState      DroneState `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BatteryLevels   []BatteryLevels
	Medications     []Medication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

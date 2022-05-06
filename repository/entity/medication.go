package entity

import "gorm.io/gorm"

type Medication struct {
	gorm.Model
	Name      string  `gorm:"not null"`
	Code      string  `gorm:"not null"`
	Weight    float64 `gorm:"not null"`
	ImagePath string
	DroneID   uint `gorm:"default:null"`
}

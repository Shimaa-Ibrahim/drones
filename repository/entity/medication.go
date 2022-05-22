package entity

import "gorm.io/gorm"

type Medication struct {
	gorm.Model
	Name      string  `json:"name" gorm:"not null"`
	Code      string  `json:"code" gorm:"not null"`
	Weight    float64 `json:"weight" gorm:"not null"`
	ImagePath string
	DroneID   uint //`gorm:"default:null"`
}

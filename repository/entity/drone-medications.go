package entity

type DroneMedication struct {
	DBModel
	DroneID      uint `gorm:"primaryKey"`
	MedicationID uint `gorm:"primaryKey;unique"`
	Loaded       bool `gorm:"default:true"`
}

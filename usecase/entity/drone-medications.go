package entity

type DroneMedications struct {
	ID            uint   `json:"id"`
	MedicatonsIDs []uint `json:"medications_ids"`
}

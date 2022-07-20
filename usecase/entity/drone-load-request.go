package entity

type DroneLoadRequest struct {
	DroneID      uint `json:"drone_id" valid:"required~drone id is required"`
	MedicationID uint `json:"medication_id" valid:"required~medication id is required"`
	Loaded       bool `json:"loaded"`
}

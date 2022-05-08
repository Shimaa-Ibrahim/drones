package entity

type BatteryLevelResponse struct {
	DroneID      uint   `json:"drone_id"`
	BatteryLevel uint64 `json:"battery_level"`
}

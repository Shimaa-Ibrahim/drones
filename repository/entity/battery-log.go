package entity

type BatteryLog struct {
	DBModel
	BatteryLevel float64 `json:"battery"`
	DroneID      uint    `json:"drone_id"`
	Drone        Drone
}

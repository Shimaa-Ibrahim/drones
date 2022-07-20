package entity

type Drone struct {
	DBModel
	SerialNumber    string       `json:"serial_number" gorm:"unique" valid:"required~serial number is required,maxstringlength(100)~serial number cannot be more than 100 characters"`
	Model           Model        `json:"model"`
	WeightLimit     float64      `json:"weight_limit" valid:"required~weight limit is required,range(0|500)~weight limit must be between 0 and 500gr"`
	BatteryCapacity float64      `json:"battery_capacity" valid:"required~battery capacity is required,range(0|100)~battery capacity must be between 0 and 100 percent"`
	State           State        `json:"state" `
	Medications     []Medication `json:"medications" gorm:"many2many:drone_medications;"`
}

type Model string
type State string

const (
	Lightweight   Model = "Lightweight"
	Middleweight  Model = "Middleweight"
	Cruiserweight Model = "Cruiserweight"
	Heavyweight   Model = "Heavyweight"
)

const (
	IDLE       State = "IDLE"
	LOADING    State = "LOADING"
	LOADED     State = "LOADED"
	DELIVERING State = "DELIVERING"
	DELIVERED  State = "DELIVERED"
	RETURNING  State = "RETURNING"
)

var DroneModels = []Model{Lightweight, Middleweight, Cruiserweight, Heavyweight}
var DroneStates = []State{IDLE, LOADING, LOADED, DELIVERING, DELIVERED, RETURNING}

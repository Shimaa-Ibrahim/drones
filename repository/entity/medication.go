package entity

type Medication struct {
	DBModel
	Name      string  `json:"name" valid:"required~name is required,matches(^[a-zA-Z0-9\\-\\_]+$)~invalid name"`
	Weight    float64 `json:"weight" valid:"required~weight is required"`
	Code      string  `json:"code" gorm:"unique" valid:"required~code is required,matches(^[a-zA-Z0-9\\_]+$)~invalid code"`
	ImagePath string  `json:"image_path"`
	Drones    []Drone `json:"-" gorm:"many2many:drone_medications"`
}

package entity

import "time"

type DBModel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

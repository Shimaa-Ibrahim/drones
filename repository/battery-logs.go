package repository

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository/entity"

	"gorm.io/gorm"
)

type BatteryLogRepo interface {
	Create(ctx context.Context, logs []entity.BatteryLog) error
}

type BatteryLogRepository struct {
	db *gorm.DB
}

func NewBatteryLogRepository(db *gorm.DB) BatteryLogRepo {
	return &BatteryLogRepository{db}
}

func (r BatteryLogRepository) Create(ctx context.Context, logs []entity.BatteryLog) error {
	err := r.db.Create(&logs).Error
	return err
}

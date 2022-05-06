package repository

import (
	"context"
	"github/Shimaa-Ibrahim/grones/repository/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DroneRepoProto interface {
	Create(ctx context.Context, drone entity.Drone) (*entity.Drone, error)
	GetByID(ctx context.Context, id uint) (*entity.Drone, error)
}

type DroneRepo struct {
	client *gorm.DB
}

func NewDroneRepository(client *gorm.DB) DroneRepoProto {
	return &DroneRepo{client: client}
}

func (ddb DroneRepo) Create(ctx context.Context, drone entity.Drone) (*entity.Drone, error) {
	result := ddb.client.WithContext(ctx).Create(&drone)
	return &drone, result.Error
}

func (ddb DroneRepo) GetByID(ctx context.Context, id uint) (*entity.Drone, error) {
	drone := &entity.Drone{}
	result := ddb.client.WithContext(ctx).Preload(clause.Associations).First(drone, id)
	return drone, result.Error
}

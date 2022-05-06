package repository

import (
	"context"
	"github/Shimaa-Ibrahim/grones/repository/entity"

	"gorm.io/gorm"
)

type DroneRepoProto interface {
	Create(ctx context.Context, drone entity.Drone) (*entity.Drone, error)
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

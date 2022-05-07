package mocks

import (
	"context"
	"github/Shimaa-Ibrahim/grones/repository"
	"github/Shimaa-Ibrahim/grones/repository/entity"
)

type MokedDroneRepo struct{}

func NewMockedDroneRepository() repository.DroneRepoProto {
	return &MokedDroneRepo{}
}

func (ddb MokedDroneRepo) Create(ctc context.Context, drone entity.Drone) (entity.Drone, error) {
	return drone, nil
}

func (ddb MokedDroneRepo) GetByID(ctx context.Context, id uint) (entity.Drone, error) {
	return entity.Drone{
		SerielNumber:    "ldrefmweoflmj956flfrv2",
		DroneModel:      entity.DroneModel{Name: "Heavyweight"},
		WeightLimit:     500,
		BatteryCapacity: 25,
		DroneState:      entity.DroneState{Name: "IDLE"},
	}, nil
}

func (ddb MokedDroneRepo) GetDronesAvailableForLoading(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{
		{
			SerielNumber:    "sca43tge56u76ybed2cw",
			DroneModel:      entity.DroneModel{Name: "Lightweight"},
			WeightLimit:     100,
			BatteryCapacity: 100,
			DroneState:      entity.DroneState{Name: "RETURNING"},
		},
		{
			SerielNumber:    "ldrefmweoflmj956flfrv2",
			DroneModel:      entity.DroneModel{Name: "Heavyweight"},
			WeightLimit:     500,
			BatteryCapacity: 50,
			DroneState:      entity.DroneState{Name: "IDLE"},
		},
	}, nil
}

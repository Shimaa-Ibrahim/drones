package mocks

import (
	"context"
	"github/Shimaa-Ibrahim/grones/repository"
	"github/Shimaa-Ibrahim/grones/repository/entity"

	"gorm.io/gorm"
)

type MokedDroneRepo struct{}

func NewMockedDroneRepository() repository.DroneRepoProto {
	return &MokedDroneRepo{}
}

func (ddb MokedDroneRepo) Create(ctc context.Context, drone entity.Drone) (entity.Drone, error) {
	return drone, nil
}

func (ddb MokedDroneRepo) GetByID(ctx context.Context, id uint) (entity.Drone, error) {
	switch id {
	case 0:
		return entity.Drone{}, gorm.ErrRecordNotFound
	case 1:
		return entity.Drone{
			SerielNumber:    "ldrefmweoflmj956flfrv2",
			DroneModel:      entity.DroneModel{Name: "Heavyweight"},
			WeightLimit:     500,
			BatteryCapacity: 80,
			DroneState:      entity.DroneState{Name: "IDLE"},
		}, nil
	case 2:
		return entity.Drone{
			SerielNumber:    "ldrefmweoflmj956flfrv2",
			DroneModel:      entity.DroneModel{Name: "Heavyweight"},
			WeightLimit:     500,
			BatteryCapacity: 80,
			DroneState:      entity.DroneState{Name: "LOADING"},
			Medications: []entity.Medication{
				{
					Name:   "med1",
					Code:   "veop45tgem",
					Weight: 50,
				},
				{
					Name:      "med2",
					Code:      "veop45tgem",
					Weight:    50,
					ImagePath: "vklitjr0g5enrbif x",
				},
			},
		}, nil
	default:
		return entity.Drone{
			SerielNumber:    "ldrefmweoflmj956flfrv2",
			DroneModel:      entity.DroneModel{Name: "Heavyweight"},
			WeightLimit:     300,
			BatteryCapacity: 20,
			DroneState:      entity.DroneState{Name: "LOADING"},
		}, nil
	}
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

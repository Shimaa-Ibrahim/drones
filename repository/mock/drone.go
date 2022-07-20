package mock

import (
	"context"
	"errors"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/entity"

	"gorm.io/gorm"
)

type MockedDrone struct{}

type SuccessMockedDrone struct {
	MockedDrone
}

type FailureMockedDrone struct {
	MockedDrone
}

type UnavailableStateMockedDrone struct {
	MockedDrone
}

type BatteryLessThan25MockedDrone struct {
	MockedDrone
}

func NewSuccessMockedDrone() repository.DroneRepo {
	return SuccessMockedDrone{}
}
func NewFailureMockedDrone() repository.DroneRepo {
	return FailureMockedDrone{}
}

func NewUnavailableStateMockedDrone() repository.DroneRepo {
	return UnavailableStateMockedDrone{}
}

func NewBatteryLessThan25MockedDrone() repository.DroneRepo {
	return BatteryLessThan25MockedDrone{}
}
func (MockedDrone) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return drone, nil
}

func (FailureMockedDrone) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	return drone, errors.New("serial number already exists")
}

func (SuccessMockedDrone) Get(ctx context.Context, id uint) (entity.Drone, error) {
	return entity.Drone{SerialNumber: "SN", Model: entity.Cruiserweight, State: entity.IDLE, BatteryCapacity: 50, WeightLimit: 100}, nil
}

func (FailureMockedDrone) Get(ctx context.Context, id uint) (entity.Drone, error) {
	return entity.Drone{}, gorm.ErrRecordNotFound
}

func (UnavailableStateMockedDrone) Get(ctx context.Context, id uint) (entity.Drone, error) {
	return entity.Drone{SerialNumber: "SN", Model: entity.Cruiserweight, State: entity.LOADED, BatteryCapacity: 50, WeightLimit: 200}, nil
}
func (BatteryLessThan25MockedDrone) Get(ctx context.Context, id uint) (entity.Drone, error) {
	return entity.Drone{SerialNumber: "SN", Model: entity.Cruiserweight, State: entity.IDLE, BatteryCapacity: 20, WeightLimit: 200}, nil
}

func (MockedDrone) GetAvailableDrones(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{{SerialNumber: "SN", Model: entity.Lightweight, State: entity.DELIVERED, BatteryCapacity: 70, WeightLimit: 40}}, nil
}

func (MockedDrone) Load(ctx context.Context, droneID uint, medicationID uint, state entity.State) (entity.Drone, error) {
	return entity.Drone{State: state}, nil
}

func (MockedDrone) GetAll(ctx context.Context) ([]entity.Drone, error) {
	return []entity.Drone{
		{SerialNumber: "SN1", Model: entity.Lightweight, State: entity.IDLE, BatteryCapacity: 50, WeightLimit: 50},
		{SerialNumber: "SN2", Model: entity.Lightweight, State: entity.DELIVERED, BatteryCapacity: 100, WeightLimit: 100},
	}, nil
}

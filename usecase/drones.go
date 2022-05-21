package usecase

import (
	"context"
	"encoding/json"
	"github/Shimaa-Ibrahim/drones/repository"
	repoEntity "github/Shimaa-Ibrahim/drones/repository/entity"
	"github/Shimaa-Ibrahim/drones/usecase/entity"
)

type DroneUseCaseProto interface {
	RegisterDrone(ctx context.Context, request []byte) ([]byte, error)
	CheckDroneLoadedItem(ctx context.Context, id uint) ([]byte, error)
	GetDronesAvailableForLoading(context.Context) ([]byte, error)
	CheckDroneBatteryLevel(context.Context, uint) ([]byte, error)
}

type DroneUseCase struct {
	droneRepository repository.DroneRepoProto
}

func NewDroneUseCase(droneRepository repository.DroneRepoProto) DroneUseCaseProto {
	return DroneUseCase{droneRepository: droneRepository}
}

func (d DroneUseCase) RegisterDrone(ctx context.Context, request []byte) ([]byte, error) {
	drone := repoEntity.Drone{}
	if err := json.Unmarshal(request, &drone); err != nil {
		return []byte{}, err
	}
	drone, err := d.droneRepository.Create(ctx, drone)
	if err != nil {
		return []byte{}, err
	}

	return json.Marshal(drone)
}

func (d DroneUseCase) CheckDroneLoadedItem(ctx context.Context, id uint) ([]byte, error) {
	drone, err := d.droneRepository.GetByID(ctx, id)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(drone)
}

func (d DroneUseCase) GetDronesAvailableForLoading(ctx context.Context) ([]byte, error) {
	drone, err := d.droneRepository.GetDronesAvailableForLoading(ctx)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(drone)

}

func (d DroneUseCase) CheckDroneBatteryLevel(ctx context.Context, id uint) ([]byte, error) {
	drone, err := d.droneRepository.GetByID(ctx, id)
	if err != nil {
		return []byte{}, err
	}
	batteryLevel := entity.BatteryLevelResponse{
		DroneID:      drone.ID,
		BatteryLevel: drone.BatteryCapacity,
	}
	return json.Marshal(batteryLevel)

}

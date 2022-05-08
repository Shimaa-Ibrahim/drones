package usecase

import (
	"context"
	"encoding/json"
	"github/Shimaa-Ibrahim/grones/repository"
	"github/Shimaa-Ibrahim/grones/repository/entity"
)

type DroneUseCaseProto interface {
	RegisterDrone(ctx context.Context, request []byte) ([]byte, error)
	CheckDroneLoadedItem(ctx context.Context, id uint) ([]byte, error)
	GetDronesAvailableForLoading(context.Context) ([]byte, error)
}

type DroneUseCase struct {
	droneRepository repository.DroneRepoProto
}

func NewDroneUseCase(droneRepository repository.DroneRepoProto) DroneUseCaseProto {
	return DroneUseCase{droneRepository: droneRepository}
}

func (d DroneUseCase) RegisterDrone(ctx context.Context, request []byte) ([]byte, error) {
	drone := entity.Drone{}
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

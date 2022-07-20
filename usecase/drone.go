package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github/Shimaa-Ibrahim/drones/repository"
	repoEntity "github/Shimaa-Ibrahim/drones/repository/entity"
	"github/Shimaa-Ibrahim/drones/usecase/entity"
	"log"

	"github.com/asaskevich/govalidator"
	"golang.org/x/exp/slices"
)

type IDroneUsecase interface {
	Create(ctx context.Context, request []byte) ([]byte, error)
	Get(ctx context.Context, id uint) ([]byte, error)
	GetAvailableDrones(ctx context.Context) ([]byte, error)
	Load(ctx context.Context, request []byte) ([]byte, error)
}

type DroneUsecase struct {
	droneRepository      repository.DroneRepo
	medicationRepository repository.MedicationRepo
}

func NewDroneUsecase(droneRepository repository.DroneRepo, medicationRepository repository.MedicationRepo) IDroneUsecase {
	return DroneUsecase{droneRepository: droneRepository, medicationRepository: medicationRepository}
}

func (d DroneUsecase) Create(ctx context.Context, request []byte) ([]byte, error) {
	drone := repoEntity.Drone{}
	err := json.Unmarshal(request, &drone)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	validStruct, err := govalidator.ValidateStruct(drone)
	if err != nil || !validStruct {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	if !slices.Contains(repoEntity.DroneModels, drone.Model) {
		err = errors.New("invalid drone model")
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	if !slices.Contains(repoEntity.DroneStates, drone.State) {
		err = errors.New("invalid drone state")
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	if drone, err = d.droneRepository.Create(ctx, drone); err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	return json.Marshal(drone)
}

func (d DroneUsecase) Get(ctx context.Context, id uint) ([]byte, error) {
	drone, err := d.droneRepository.Get(ctx, id)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	return json.Marshal(drone)
}

func (d DroneUsecase) GetAvailableDrones(ctx context.Context) ([]byte, error) {
	drones, err := d.droneRepository.GetAvailableDrones(ctx)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	return json.Marshal(drones)
}

func (d DroneUsecase) Load(ctx context.Context, request []byte) ([]byte, error) {
	state := repoEntity.LOADING
	req := entity.DroneLoadRequest{}
	if err := json.Unmarshal(request, &req); err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	validStruct, err := govalidator.ValidateStruct(req)
	if err != nil || !validStruct {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	drone, err := d.droneRepository.Get(ctx, req.DroneID)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	if drone.State != repoEntity.IDLE && drone.State != repoEntity.LOADING || drone.BatteryCapacity < 25 {
		err = errors.New("drone is not available")
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	medication, err := d.medicationRepository.Get(ctx, req.MedicationID)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	if len(medication.Drones) > 0 {
		err = errors.New("medication is already loaded")
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	totalWeight := calculateLoadedWeight(drone) + medication.Weight
	if totalWeight > drone.WeightLimit {
		err = errors.New("drone weight limit exceeded")
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	if totalWeight == drone.WeightLimit || req.Loaded {
		state = repoEntity.LOADED
	}
	drone, err = d.droneRepository.Load(ctx, req.DroneID, req.MedicationID, state)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	return json.Marshal(drone)
}

func calculateLoadedWeight(drone repoEntity.Drone) float64 {
	loadedWeight := 0.0
	for _, medication := range drone.Medications {
		loadedWeight += medication.Weight
	}
	return loadedWeight
}

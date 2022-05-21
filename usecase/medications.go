package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/usecase/entity"

	"gorm.io/gorm"
)

type MedicationUseCaseProto interface {
	LoadDroneWithMedicationItems(ctx context.Context, request []byte) error
}

type MedicationUseCase struct {
	medicationRepo repository.MedicationRepoProto
	droneRepo      repository.DroneRepoProto
}

func NewMedicationUseCase(medicationRepo repository.MedicationRepoProto, droneRepo repository.DroneRepoProto) MedicationUseCaseProto {
	return &MedicationUseCase{
		medicationRepo: medicationRepo,
		droneRepo:      droneRepo,
	}
}

func (m MedicationUseCase) LoadDroneWithMedicationItems(ctx context.Context, request []byte) error {
	droneMedicationsData := entity.DroneMedications{}
	if err := json.Unmarshal(request, &droneMedicationsData); err != nil {
		return err
	}

	drone, err := m.droneRepo.GetByID(ctx, droneMedicationsData.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("The drone does not exist")
		}
		return err
	}
	if drone.DroneState.Name != "LOADING" {
		return errors.New("The drone is not available")
	}

	if drone.BatteryCapacity < 25 {
		return errors.New("the drone's battery level below 25")
	}
	var medsNotFound []uint
	var loadedMedictions []uint
	totalWeight := 0

	for _, medID := range droneMedicationsData.MedicatonsIDs {
		med, err := m.medicationRepo.GetByID(ctx, medID)
		if err != nil {
			medsNotFound = append(medsNotFound, medID)
			continue
		}
		if med.DroneID != 0 {
			loadedMedictions = append(loadedMedictions, medID)
		} else {
			totalWeight += int(med.Weight)
		}
	}

	if len(medsNotFound) != 0 {
		return errors.New(fmt.Sprintf("These medications do not exist: %v", medsNotFound))
	}

	if len(loadedMedictions) != 0 {
		return errors.New(fmt.Sprintf("These medications are already loaded: %v", loadedMedictions))
	}

	if totalWeight > int(drone.WeightLimit) {
		return errors.New("The medications weight exceeded the drone's weight limit")
	}

	err = m.medicationRepo.UpdateMedicationsWithDroneID(ctx, droneMedicationsData.ID, droneMedicationsData.MedicatonsIDs)
	return err
}

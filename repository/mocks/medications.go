package mocks

import (
	"context"
	"errors"
	"github/Shimaa-Ibrahim/grones/repository"
	"github/Shimaa-Ibrahim/grones/repository/entity"
)

type MockedMedicationRepo struct{}

func NewMockedMedicationRepository() repository.MedicationRepoProto {
	return &MockedMedicationRepo{}
}

func (mdb MockedMedicationRepo) UpdateMedicationsWithDroneID(ctx context.Context, droneID uint, medicationsIDs []uint) error {
	return nil
}

func (mdb MockedMedicationRepo) GetByID(ctx context.Context, id uint) (entity.Medication, error) {
	switch id {
	case 0:
		err := errors.New("Not Found")
		return entity.Medication{}, err
	case 1:
		return entity.Medication{
			Name:   "med1",
			Code:   "veop45tgem",
			Weight: 50,
		}, nil
	case 2:
		return entity.Medication{
			Name:    "med2",
			Code:    "veop45tgem",
			Weight:  70,
			DroneID: 1,
		}, nil
	default:
		return entity.Medication{
			Name:   "med3",
			Code:   "veop45tgem",
			Weight: 3000,
		}, nil

	}
}

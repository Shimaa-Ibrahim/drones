package mock

import (
	"context"
	"errors"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/entity"

	"gorm.io/gorm"
)

type MockedMedication struct{}

type FailureMockedMedication struct {
	MockedMedication
}
type LoadedMockedMedication struct {
	MockedMedication
}

type OverWeightMockedMedication struct {
	MockedMedication
}

func NewMockedMedication() repository.MedicationRepo {
	return &MockedMedication{}
}

func NewFailureMockedMedication() repository.MedicationRepo {
	return &FailureMockedMedication{}
}

func NewLoadedMockedMedication() repository.MedicationRepo {
	return &LoadedMockedMedication{}
}

func NewOverWeightMockedMedication() repository.MedicationRepo {
	return &OverWeightMockedMedication{}
}

func (MockedMedication) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return medication, nil
}

func (FailureMockedMedication) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	return medication, errors.New("code already exists")
}

func (MockedMedication) Get(ctx context.Context, id uint) (entity.Medication, error) {
	return entity.Medication{Name: "Medication", Weight: 50, Code: "code", ImagePath: "image_path"}, nil
}

func (FailureMockedMedication) Get(ctx context.Context, id uint) (entity.Medication, error) {
	return entity.Medication{}, gorm.ErrRecordNotFound
}

func (OverWeightMockedMedication) Get(ctx context.Context, id uint) (entity.Medication, error) {
	return entity.Medication{Name: "Medication", Weight: 5000, Code: "code", ImagePath: "image_path"}, nil
}

func (LoadedMockedMedication) Get(ctx context.Context, id uint) (entity.Medication, error) {
	return entity.Medication{Name: "Medication", Weight: 50, Code: "code", ImagePath: "image_path", Drones: []entity.Drone{
		{SerialNumber: "SN", Model: entity.Lightweight, State: entity.DELIVERED, BatteryCapacity: 70, WeightLimit: 40},
	}}, nil
}

func (MockedMedication) GetAll(ctx context.Context) ([]entity.Medication, error) {
	return []entity.Medication{{Name: "Medication", Weight: 50, Code: "code", ImagePath: "image_path"}}, nil
}

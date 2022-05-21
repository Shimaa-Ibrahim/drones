package repository

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository/entity"

	"gorm.io/gorm"
)

type MedicationRepoProto interface {
	UpdateMedicationsWithDroneID(ctx context.Context, droneID uint, medicationsIDs []uint) error
	GetByID(ctx context.Context, id uint) (entity.Medication, error)
	CreateMedication(ctx context.Context, medication entity.Medication) (entity.Medication, error)
}

type MedicationRepo struct {
	client *gorm.DB
}

func NewMedicationRepository(client *gorm.DB) MedicationRepoProto {
	return &MedicationRepo{client: client}
}

func (mdb MedicationRepo) UpdateMedicationsWithDroneID(ctx context.Context, droneID uint, medicationsIDs []uint) error {
	result := mdb.client.WithContext(ctx).Model(&entity.Medication{}).Where("id IN ?", medicationsIDs).Update("drone_id", droneID)
	return result.Error
}

func (mdb MedicationRepo) GetByID(ctx context.Context, id uint) (entity.Medication, error) {
	medication := entity.Medication{}
	err := mdb.client.First(&medication, id).Error
	return medication, err
}

func (mdb MedicationRepo) CreateMedication(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	result := mdb.client.WithContext(ctx).Omit("DroneID").Create(&medication)
	return medication, result.Error
}

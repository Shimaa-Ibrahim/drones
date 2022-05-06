package repository

import (
	"context"
	"github/Shimaa-Ibrahim/grones/repository/entity"

	"gorm.io/gorm"
)

type MedicationRepoProto interface {
	UpdateMedicationsWithDroneID(ctx context.Context, droneID uint, medicationsIDs []uint) error
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

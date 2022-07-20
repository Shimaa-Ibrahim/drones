package repository

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository/entity"

	"gorm.io/gorm"
)

type MedicationRepo interface {
	Create(ctx context.Context, medication entity.Medication) (entity.Medication, error)
	Get(ctx context.Context, id uint) (entity.Medication, error)
	GetAll(ctx context.Context) ([]entity.Medication, error)
}

type MedicationRepository struct {
	db *gorm.DB
}

func NewMedicationRepository(db *gorm.DB) MedicationRepo {
	return MedicationRepository{db: db}
}

func (m MedicationRepository) Create(ctx context.Context, medication entity.Medication) (entity.Medication, error) {
	err := m.db.WithContext(ctx).Create(&medication).Error
	return medication, err
}

func (m MedicationRepository) Get(ctx context.Context, id uint) (entity.Medication, error) {
	medication := entity.Medication{}
	err := m.db.WithContext(ctx).
		Preload("Drones", func(db *gorm.DB) *gorm.DB {
			return db.Joins("LEFT JOIN drones.drone_medications dm ON drones.drones.id = dm.drone_id").Where("loaded = ?", true)
		}).
		First(&medication, id).Error
	return medication, err
}

// use pagination TODO
func (m MedicationRepository) GetAll(ctx context.Context) ([]entity.Medication, error) {
	medications := []entity.Medication{}
	err := m.db.WithContext(ctx).Find(&medications).Error
	return medications, err
}

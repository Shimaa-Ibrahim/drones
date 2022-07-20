package repository

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository/entity"

	"gorm.io/gorm"
)

type DroneRepo interface {
	Create(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	Get(ctx context.Context, id uint) (entity.Drone, error)
	GetAvailableDrones(ctx context.Context) ([]entity.Drone, error)
	Load(ctx context.Context, droneID uint, medicationID uint, state entity.State) (entity.Drone, error)
	GetAll(ctx context.Context) ([]entity.Drone, error)
}

type DroneRepository struct {
	db *gorm.DB
}

func NewDroneRepository(db *gorm.DB) DroneRepo {
	return DroneRepository{db: db}
}

func (d DroneRepository) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	err := d.db.WithContext(ctx).Create(&drone).Error
	return drone, err
}

func (d DroneRepository) Get(ctx context.Context, id uint) (entity.Drone, error) {
	drone := entity.Drone{}
	err := d.db.WithContext(ctx).
		Preload("Medications", func(db *gorm.DB) *gorm.DB {
			return db.Joins("LEFT JOIN drones.drone_medications dm ON drones.medications.id = dm.medication_id").Where("loaded = ?", true)
		}).
		First(&drone, id).Error
	return drone, err
}

func (d DroneRepository) GetAvailableDrones(ctx context.Context) ([]entity.Drone, error) {
	drones := []entity.Drone{}
	err := d.db.WithContext(ctx).
		Joins("LEFT JOIN drones.drone_medications dm ON drones.drones.id = dm.drone_id").
		Joins("LEFT JOIN drones.medications m ON m.id = dm.medication_id").
		Having("(SUM(weight) < weight_limit AND loaded = true) OR SUM(weight) IS NULL").
		Group("drones.drones.id, dm.loaded").
		Order("drones.drones.id").
		Preload("Medications", func(db *gorm.DB) *gorm.DB {
			return db.Joins("LEFT JOIN drones.drone_medications dm ON drones.medications.id = dm.medication_id").Where("loaded = ?", true)
		}).
		Find(&drones, "battery_capacity >= ? AND state IN (?)", 25, []string{"IDLE", "LOADING"}).Error
	return drones, err
}

func (d DroneRepository) Load(ctx context.Context, droneID uint, medicationID uint, state entity.State) (entity.Drone, error) {
	drone := entity.Drone{DBModel: entity.DBModel{ID: droneID}}
	err := d.db.WithContext(ctx).Model(&drone).Update("state", state).Association("Medications").Append(&entity.Medication{DBModel: entity.DBModel{ID: medicationID}})
	return drone, err
}

func (d DroneRepository) GetAll(ctx context.Context) ([]entity.Drone, error) {
	drones := []entity.Drone{}
	err := d.db.WithContext(ctx).Find(&drones).Error
	return drones, err
}

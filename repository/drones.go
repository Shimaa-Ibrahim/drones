package repository

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DroneRepoProto interface {
	Create(ctx context.Context, drone entity.Drone) (entity.Drone, error)
	GetByID(ctx context.Context, id uint) (entity.Drone, error)
	GetDronesAvailableForLoading(ctx context.Context) ([]entity.Drone, error)
	Get(ctx context.Context) ([]entity.Drone, error)
	LogDronesBatteryLevel() ([]entity.BatteryLevels, error)
}

type DroneRepo struct {
	client *gorm.DB
}

func NewDroneRepository(client *gorm.DB) DroneRepoProto {
	return &DroneRepo{client: client}
}

func (ddb DroneRepo) Create(ctx context.Context, drone entity.Drone) (entity.Drone, error) {
	result := ddb.client.WithContext(ctx).Create(&drone)
	return drone, result.Error
}

func (ddb DroneRepo) GetByID(ctx context.Context, id uint) (entity.Drone, error) {
	drone := entity.Drone{}
	result := ddb.client.WithContext(ctx).Preload(clause.Associations).First(&drone, id)
	return drone, result.Error
}

func (ddb DroneRepo) GetDronesAvailableForLoading(ctx context.Context) ([]entity.Drone, error) {
	var drones []entity.Drone
	subQuery := ddb.client.Model(&entity.Medication{}).Select("SUM(weight) as loaded_weight, drone_id").Group("drone_id")
	result := ddb.client.WithContext(ctx).
		Joins("LEFT JOIN (?) sumResult ON sumResult.drone_id = drones.drones.id", subQuery).
		Joins(`JOIN drones.drone_states ON drones.drone_states.id = drones.drones.drone_state_id AND drones.drone_states.name = ?`, "LOADING").
		Preload(clause.Associations).
		Order("id").
		Find(&drones, "battery_capacity >= ? AND (drones.drones.weight_limit > sumResult.loaded_weight OR loaded_weight IS NULL)", 25)
	return drones, result.Error
}

func (ddb DroneRepo) Get(ctx context.Context) ([]entity.Drone, error) {
	var drones []entity.Drone
	result := ddb.client.WithContext(ctx).Find(&drones)
	return drones, result.Error
}

func (ddb DroneRepo) LogDronesBatteryLevel() ([]entity.BatteryLevels, error) {
	drones := []entity.Drone{}
	logsRecord := []entity.BatteryLevels{}
	if err := ddb.client.Find(&drones).Error; err != nil {
		return logsRecord, err
	}

	err := ddb.client.Transaction(func(tx *gorm.DB) error {
		for _, drone := range drones {
			batteryLevelRecord := entity.BatteryLevels{BatteryLevel: float64(drone.BatteryCapacity)}
			if err := ddb.client.Model(&drone).Association("BatteryLevels").Append(&batteryLevelRecord); err != nil {
				return err
			}
			logsRecord = append(logsRecord, batteryLevelRecord)
		}
		return nil
	})
	return logsRecord, err
}

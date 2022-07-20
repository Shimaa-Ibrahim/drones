package mock

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/entity"
)

type MockedBatteryLogsRepository struct{}

func NewMockedBatteryLogsRepository() repository.BatteryLogRepo {
	return MockedBatteryLogsRepository{}
}

func (MockedBatteryLogsRepository) Create(ctx context.Context, batteryLogs []entity.BatteryLog) error {
	return nil
}

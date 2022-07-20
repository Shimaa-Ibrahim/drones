package repository

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLog(t *testing.T) {
	batteryLogRepository := NewBatteryLogRepository(dbClient)
	drones := []entity.Drone{
		{SerialNumber: "SNN1", Model: "M1", WeightLimit: 10, BatteryCapacity: 100, State: "IDLE"},
		{SerialNumber: "SNN2", Model: "M2", WeightLimit: 50, BatteryCapacity: 50, State: "LOADING"},
	}
	if err := dbClient.Create(&drones).Error; err != nil {
		t.Errorf("[Error] drone creation: %v", err)
	}
	type args struct {
		ctx  context.Context
		logs []entity.BatteryLog
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.BatteryLog
		wantErr string
	}{
		{
			name: "[Test] create battery logs with valid data",
			args: args{context.Background(), []entity.BatteryLog{
				{BatteryLevel: 10, DroneID: drones[0].ID},
				{BatteryLevel: 20, DroneID: drones[1].ID},
			}},
			want: []entity.BatteryLog{
				{BatteryLevel: 10, DroneID: drones[0].ID},
				{BatteryLevel: 20, DroneID: drones[1].ID},
			},
		},
		{
			name:    "[Test] create battery logs should return error if drone id is invalid",
			args:    args{context.Background(), []entity.BatteryLog{{BatteryLevel: 10, DroneID: 0}}},
			wantErr: `ERROR: insert or update on table "battery_logs" violates foreign key constraint "fk_battery_logs_drone" (SQLSTATE 23503)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := batteryLogRepository.Create(tt.args.ctx, tt.args.logs)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			got := []entity.BatteryLog{}
			if err := dbClient.Find(&got).Error; err != nil {
				t.Errorf("[Error] battery log retrieval: %v", err)
			}
			assert.Equal(t, len(tt.want), len(got))
			for i, batteryLog := range got {
				checkBatteryLogEquality(t, tt.want[i], batteryLog)
			}
		})
	}
}

func checkBatteryLogEquality(t *testing.T, want entity.BatteryLog, got entity.BatteryLog) {
	assert.NotEmpty(t, got.ID)
	assert.Equal(t, want.DroneID, got.DroneID)
	assert.Equal(t, want.BatteryLevel, got.BatteryLevel)
}

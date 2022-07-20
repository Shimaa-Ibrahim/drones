package repository

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDroneCreation(t *testing.T) {
	clearDataBase()
	droneRepository := NewDroneRepository(dbClient)
	if err := dbClient.Create(&entity.Drone{SerialNumber: "don't duplicate me"}).Error; err != nil {
		t.Errorf("[Error] drone creation: %v", err)
	}
	type args struct {
		ctx   context.Context
		drone entity.Drone
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Drone
		wantErr string
	}{
		{
			name: "[Test] create drone with valid data",
			args: args{context.Background(), entity.Drone{SerialNumber: "SN", Model: "M", WeightLimit: 10, BatteryCapacity: 100, State: "IDLE"}},
			want: entity.Drone{SerialNumber: "SN", Model: "M", WeightLimit: 10, BatteryCapacity: 100, State: "IDLE"},
		},
		{
			name:    "[Test] create drone should return error if drone serial number is duplicate",
			args:    args{context.Background(), entity.Drone{SerialNumber: "don't duplicate me"}},
			wantErr: `ERROR: duplicate key value violates unique constraint "drones_serial_number_key" (SQLSTATE 23505)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := droneRepository.Create(tt.args.ctx, tt.args.drone)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			got := entity.Drone{}
			if err := dbClient.Find(&got, "serial_number", tt.args.drone.SerialNumber).Error; err != nil {
				t.Errorf("[Error] drone retrieval: %v", err)
			}
			checkDroneEquality(t, tt.want, got)
		})
	}
}

func TestGetAvailableDrones(t *testing.T) {
	clearDataBase()
	droneRepository := NewDroneRepository(dbClient)
	drones := []entity.Drone{
		{SerialNumber: "SN1", Model: "M1", WeightLimit: 10, BatteryCapacity: 100, State: "IDLE"},
		{SerialNumber: "SN2", Model: "M2", WeightLimit: 20, BatteryCapacity: 20, State: "LOADING"},
		{SerialNumber: "SN3", Model: "M3", WeightLimit: 20, BatteryCapacity: 10, State: "LOADED"},
		{SerialNumber: "SN4", Model: "M4", WeightLimit: 20, BatteryCapacity: 25, State: "LOADING", Medications: []entity.Medication{{Code: "C1", Weight: 10}, {Code: "C3", Weight: 10}}},
		{SerialNumber: "SN5", Model: "M5", WeightLimit: 50, BatteryCapacity: 50, State: "LOADING", Medications: []entity.Medication{{Code: "C2", Weight: 10}, {Code: "C4", Weight: 10}}},
	}
	if err := dbClient.Create(&drones).Error; err != nil {
		t.Errorf("[Error] drone creation: %v", err)
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Drone
		wantErr string
	}{
		{
			name: "[Test] get available drones",
			args: args{context.Background()},
			want: []entity.Drone{
				{SerialNumber: "SN1", Model: "M1", WeightLimit: 10, BatteryCapacity: 100, State: "IDLE"},
				{SerialNumber: "SN5", Model: "M5", WeightLimit: 50, BatteryCapacity: 50, State: "LOADING", Medications: []entity.Medication{{Code: "C2", Weight: 10}, {Code: "C4", Weight: 10}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := droneRepository.GetAvailableDrones(tt.args.ctx)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, len(tt.want), len(got))
			for i, drone := range got {
				checkDroneEquality(t, tt.want[i], drone)
			}
		})
	}
}

func TestGetDrone(t *testing.T) {
	clearDataBase()
	droneRepository := NewDroneRepository(dbClient)
	drones := []entity.Drone{
		{SerialNumber: "SNN1", Model: "M1", WeightLimit: 10, BatteryCapacity: 100, State: "IDLE"},
		{SerialNumber: "SNN2", Model: "M2", WeightLimit: 50, BatteryCapacity: 50, State: "LOADING", Medications: []entity.Medication{{Code: "C1", Weight: 10}, {Code: "C3", Weight: 10}}},
	}
	if err := dbClient.Create(&drones).Error; err != nil {
		t.Errorf("[Error] drone creation: %v", err)
	}
	if err := dbClient.Model(&entity.DroneMedication{}).Where("drone_id=? AND medication_id=?", drones[1].ID, drones[1].Medications[0].ID).Update("loaded", false).Error; err != nil {
		t.Errorf("[Error] drone medication update: %v", err)
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Drone
		wantErr string
	}{
		{
			name: "[Test] get drone with valid data",
			args: args{context.Background(), drones[0].ID},
			want: drones[0],
		},
		{
			name: "[Test] get drone should return sec medication only as first is not loaded anymore",
			args: args{context.Background(), drones[1].ID},
			want: entity.Drone{SerialNumber: "SNN2", Model: "M2", WeightLimit: 50, BatteryCapacity: 50, State: "LOADING", Medications: []entity.Medication{{Code: "C3", Weight: 10}}},
		},
		{
			name:    "[Test] get drone should return error if drone id is invalid",
			args:    args{context.Background(), 0},
			wantErr: "record not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := droneRepository.Get(tt.args.ctx, tt.args.id)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			checkDroneEquality(t, tt.want, got)
		})
	}
}

func TestLoadDrone(t *testing.T) {
	clearDataBase()
	droneRepository := NewDroneRepository(dbClient)
	medications := []entity.Medication{
		{Code: "Code1", Weight: 10},
		{Code: "Code2", Weight: 20},
		{Code: "Code3", Weight: 30},
	}
	if err := dbClient.Create(&medications).Error; err != nil {
		t.Errorf("[Error] medications creation: %v", err)
	}
	drones := []entity.Drone{
		{SerialNumber: "SSNN1", Model: "M1", WeightLimit: 400, BatteryCapacity: 100, State: "IDLE"},
		{SerialNumber: "SSNN2", Model: "M2", WeightLimit: 500, BatteryCapacity: 50, State: "LOADING"},
	}
	if err := dbClient.Create(&drones).Error; err != nil {
		t.Errorf("[Error] drone creation: %v", err)
	}
	type args struct {
		ctx          context.Context
		droneID      uint
		medicationID uint
		state        entity.State
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Drone
		wantErr string
	}{
		{
			name: "[Test] load drone with med and change state to LOADING",
			args: args{context.Background(), drones[0].ID, medications[0].ID, entity.LOADING},
			want: entity.Drone{SerialNumber: "SSNN1", Model: "M1", WeightLimit: 400, BatteryCapacity: 100, State: "LOADING", Medications: []entity.Medication{medications[0]}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := droneRepository.Load(tt.args.ctx, tt.args.droneID, tt.args.medicationID, tt.args.state)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			got := entity.Drone{}
			if err := dbClient.Preload("Medications").Find(&got, "serial_number", drones[0].SerialNumber).Error; err != nil {
				t.Errorf("[Error] drone retrieval: %v", err)
			}
			checkDroneEquality(t, tt.want, got)
		})
	}

}

func TestAllDrones(t *testing.T) {
	clearDataBase()
	droneRepository := NewDroneRepository(dbClient)
	drones := []entity.Drone{
		{SerialNumber: "SN1", Model: "M1", WeightLimit: 10, BatteryCapacity: 100, State: "IDLE"},
		{SerialNumber: "SN2", Model: "M2", WeightLimit: 20, BatteryCapacity: 20, State: "LOADING"},
		{SerialNumber: "SN3", Model: "M3", WeightLimit: 20, BatteryCapacity: 10, State: "LOADED"},
	}
	if err := dbClient.Create(&drones).Error; err != nil {
		t.Errorf("[Error] drone creation: %v", err)
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Drone
		wantErr string
	}{
		{
			name: "[Test] get All drones",
			args: args{context.Background()},
			want: drones,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := droneRepository.GetAll(tt.args.ctx)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, len(tt.want), len(got))
			for i, drone := range got {
				checkDroneEquality(t, tt.want[i], drone)
			}
		})
	}
}

func checkDroneEquality(t *testing.T, want entity.Drone, got entity.Drone) {
	t.Helper()
	assert.NotEmpty(t, got.ID)
	assert.Equal(t, want.SerialNumber, got.SerialNumber)
	assert.Equal(t, want.Model, got.Model)
	assert.Equal(t, want.WeightLimit, got.WeightLimit)
	assert.Equal(t, want.BatteryCapacity, got.BatteryCapacity)
	assert.Equal(t, want.State, got.State)
	assert.Equal(t, len(want.Medications), len(got.Medications))
	for i, medication := range got.Medications {
		assert.NotEmpty(t, medication.ID)
		assert.Equal(t, want.Medications[i].Code, medication.Code)
	}
}

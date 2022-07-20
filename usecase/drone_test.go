package usecase

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDrone(t *testing.T) {
	type args struct {
		ctx context.Context
		req []byte
	}
	type mocks struct {
		droneRepository      repository.DroneRepo
		medicationRepository repository.MedicationRepo
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
		mocks   mocks
	}{
		{
			name:  "[Test] succssful drone creation",
			args:  args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "IDLE", "battery_capacity": 50, "weight_limit": 200}`)},
			want:  `{"id":0,"serial_number":"SN","model":"Cruiserweight","weight_limit":200,"battery_capacity":50,"state":"IDLE","medications":null}`,
			mocks: mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] drone creation should return error if serial number is already taken",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "IDLE", "battery_capacity": 50, "weight_limit": 200}`)},
			wantErr: "serial number already exists",
			mocks:   mocks{droneRepository: mock.NewFailureMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] drone creation should return error if serial number is not provided",
			args:    args{context.Background(), []byte(`{"model": "Cruiserweight", "state": "IDLE", "battery_capacity": 50, "weight_limit": 200}`)},
			wantErr: "serial number is required",
		},
		{
			name:    "[Test] drone creation should return error if serial number is more than 100 characters",
			args:    args{context.Background(), []byte(`{"serial_number": "sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss", "model": "Cruiserweight", "state": "IDLE", "battery_capacity": 50, "weight_limit": 200}`)},
			wantErr: "serial number cannot be more than 100 characters",
		},
		{
			name:    "[Test] drone creation should return error if model is not provided",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "state": "IDLE", "battery_capacity": 50, "weight_limit": 200}`)},
			wantErr: "invalid drone model",
		},
		{
			name:    "[Test] drone creation should return error if model is invalid",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "invalid", "state": "IDLE", "battery_capacity": 50, "weight_limit": 200}`)},
			wantErr: "invalid drone model",
		},
		{
			name:    "[Test] drone creation should return error if battery capacity is not provided",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "IDLE", "weight_limit": 200}`)},
			wantErr: "battery capacity is required",
		},
		{
			name:    "[Test] drone creation should return error if battery capacity is invalid",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "IDLE", "battery_capacity": -1, "weight_limit": 200}`)},
			wantErr: "battery capacity must be between 0 and 100 percent",
		},
		{
			name:    "[Test] drone creation should return error if battery capacity is invalid",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "IDLE", "battery_capacity": 101, "weight_limit": 200}`)},
			wantErr: "battery capacity must be between 0 and 100 percent",
		},
		{
			name:    "[Test] drone creation should return error if weight limit is not provided",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "IDLE", "battery_capacity": 50}`)},
			wantErr: "weight limit is required",
		},
		{
			name:    "[Test] drone creation should return error if weight limit is invalid",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "IDLE", "battery_capacity": 50, "weight_limit": -1}`)},
			wantErr: "weight limit must be between 0 and 500gr",
		},
		{
			name:    "[Test] drone creation should return error if weight limit is invalid",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "IDLE", "battery_capacity": 50, "weight_limit": 505}`)},
			wantErr: "weight limit must be between 0 and 500gr",
		},
		{
			name:    "[Test] drone creation should return error if state is not provided",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "battery_capacity": 50, "weight_limit": 200}`)},
			wantErr: "invalid drone state",
		},
		{
			name:    "[Test] drone creation should return error if state is invalid",
			args:    args{context.Background(), []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "invalid", "battery_capacity": 50, "weight_limit": 200}`)},
			wantErr: "invalid drone state",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := NewDroneUsecase(tt.mocks.droneRepository, tt.mocks.medicationRepository)
			got, err := droneUsecase.Create(tt.args.ctx, tt.args.req)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestGetDrone(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint
	}
	type mocks struct {
		droneRepository      repository.DroneRepo
		medicationRepository repository.MedicationRepo
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
		mocks   mocks
	}{
		{
			name:  "[Test] sussessful drone retrieval",
			args:  args{context.Background(), 1},
			want:  `{"id":0,"serial_number":"SN","model":"Cruiserweight","weight_limit":100,"battery_capacity":50,"state":"IDLE","medications":null}`,
			mocks: mocks{droneRepository: mock.NewSuccessMockedDrone()},
		},
		{
			name:    "[Test] drone retrieval should return error if drone does not exist",
			args:    args{context.Background(), 2},
			wantErr: "record not found",
			mocks:   mocks{droneRepository: mock.NewFailureMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := NewDroneUsecase(tt.mocks.droneRepository, tt.mocks.medicationRepository)
			got, err := droneUsecase.Get(tt.args.ctx, tt.args.id)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestGetAvailabeDrones(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type mocks struct {
		droneRepository      repository.DroneRepo
		medicationRepository repository.MedicationRepo
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
		mocks   mocks
	}{
		{
			name:  "[Test] sussessful drone retrieval",
			args:  args{context.Background()},
			want:  `[{"id":0,"serial_number":"SN","model":"Lightweight","weight_limit":40,"battery_capacity":70,"state":"DELIVERED","medications":null}]`,
			mocks: mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := NewDroneUsecase(tt.mocks.droneRepository, tt.mocks.medicationRepository)
			got, err := droneUsecase.GetAvailableDrones(tt.args.ctx)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestLoadDrone(t *testing.T) {
	type args struct {
		ctx context.Context
		req []byte
	}
	type mocks struct {
		droneRepository      repository.DroneRepo
		medicationRepository repository.MedicationRepo
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
		mocks   mocks
	}{
		{
			name:  "[Test] sussessful drone loading with loading state",
			args:  args{context.Background(), []byte(`{"drone_id":1,"medication_id":1, "loaded": false}`)},
			want:  `{"id":0,"serial_number":"","model":"","weight_limit":0,"battery_capacity":0,"state":"LOADING","medications":null}`,
			mocks: mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:  "[Test] sussessful drone loading with loaded state",
			args:  args{context.Background(), []byte(`{"drone_id":1,"medication_id":1, "loaded": true}`)},
			want:  `{"id":0,"serial_number":"","model":"","weight_limit":0,"battery_capacity":0,"state":"LOADED","medications":null}`,
			mocks: mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] drone loading should return error if drone does not exist",
			args:    args{context.Background(), []byte(`{"drone_id":2,"medication_id":1, "loaded": true}`)},
			wantErr: "record not found",
			mocks:   mocks{droneRepository: mock.NewFailureMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] drone loading should return error if medication does not exist",
			args:    args{context.Background(), []byte(`{"drone_id":1,"medication_id":2, "loaded": true}`)},
			wantErr: "record not found",
			mocks:   mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewFailureMockedMedication()},
		},
		{
			name:    "[Test] drone loading should return error if medication is already loaded",
			args:    args{context.Background(), []byte(`{"drone_id":1,"medication_id":1, "loaded": true}`)},
			wantErr: "medication is already loaded",
			mocks:   mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewLoadedMockedMedication()},
		},
		{
			name:    "[Test] drone loading should return error if drone id is not provided",
			args:    args{context.Background(), []byte(`{"medication_id":1, "loaded": true}`)},
			wantErr: "drone id is required",
			mocks:   mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] drone loading should return error if medication id is not provided",
			args:    args{context.Background(), []byte(`{"drone_id":1, "loaded": true}`)},
			wantErr: "medication id is required",
			mocks:   mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] drone loading should return error if drone state not LOADING or IDLE",
			args:    args{context.Background(), []byte(`{"drone_id":1,"medication_id":1, "loaded": true}`)},
			wantErr: "drone is not available",
			mocks:   mocks{droneRepository: mock.UnavailableStateMockedDrone{}, medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] drone loading should return error if drone battery capacity is less than 25",
			args:    args{context.Background(), []byte(`{"drone_id":1,"medication_id":1, "loaded": true}`)},
			wantErr: "drone is not available",
			mocks:   mocks{droneRepository: mock.BatteryLessThan25MockedDrone{}, medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] drone loading should return error if weight exceeds drone weight limit",
			args:    args{context.Background(), []byte(`{"drone_id":1,"medication_id":1, "loaded": true}`)},
			wantErr: "drone weight limit exceeded",
			mocks:   mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewOverWeightMockedMedication()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := NewDroneUsecase(tt.mocks.droneRepository, tt.mocks.medicationRepository)
			got, err := droneUsecase.Load(tt.args.ctx, tt.args.req)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}

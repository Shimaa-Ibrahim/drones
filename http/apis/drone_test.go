package apis

import (
	"bytes"
	"context"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/mock"
	"github/Shimaa-Ibrahim/drones/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRegisterDrone(t *testing.T) {
	type mocks struct {
		droneRepository      repository.DroneRepo
		medicationRepository repository.MedicationRepo
	}
	tests := []struct {
		name       string
		ctx        context.Context
		reqBody    []byte
		want       string
		statusCode int
		mocks      mocks
	}{
		{
			name:       "[Test] successful register drone",
			ctx:        context.Background(),
			reqBody:    []byte(`{"serial_number": "SN", "model": "Cruiserweight", "state": "IDLE", "battery_capacity": 50, "weight_limit": 200}`),
			want:       `{"id":0,"serial_number":"SN","model":"Cruiserweight","weight_limit":200,"battery_capacity":50,"state":"IDLE","medications":null}`,
			statusCode: http.StatusCreated,
			mocks:      mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:       "[Test] register drone should return error if serial number is not provided",
			ctx:        context.Background(),
			reqBody:    []byte(`{"model": "Cruiserweight", "state": "IDLE", "battery_capacity": 50, "weight_limit": 200}`),
			want:       `{"error":"serial number is required"}`,
			statusCode: http.StatusBadRequest,
			mocks:      mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := usecase.NewDroneUsecase(tt.mocks.droneRepository, tt.mocks.medicationRepository)
			droneAPIs := NewDroneAPIs(droneUsecase)
			req, err := http.NewRequest("POST", "", bytes.NewBuffer(tt.reqBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(droneAPIs.Register)
			handler.ServeHTTP(rr, req.WithContext(tt.ctx))
			result := rr.Body.String()
			assert.Equal(t, rr.Code, tt.statusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGetDrone(t *testing.T) {
	type mocks struct {
		droneRepository      repository.DroneRepo
		medicationRepository repository.MedicationRepo
	}
	tests := []struct {
		name       string
		ctx        context.Context
		url        string
		want       string
		statusCode int
		mocks      mocks
	}{
		{
			name:       "[Test] successful drone retrieval",
			ctx:        context.Background(),
			url:        "/1",
			want:       `{"id":0,"serial_number":"SN","model":"Cruiserweight","weight_limit":100,"battery_capacity":50,"state":"IDLE","medications":null}`,
			statusCode: http.StatusOK,
			mocks:      mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:       "[Test] get drone should return error if drone id is invalid",
			ctx:        context.Background(),
			url:        "/string",
			want:       `{"error":"invalid ID"}`,
			statusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := usecase.NewDroneUsecase(tt.mocks.droneRepository, tt.mocks.medicationRepository)
			droneAPIs := NewDroneAPIs(droneUsecase)
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/{id}", droneAPIs.Get).Methods("GET")
			router.ServeHTTP(rr, req.WithContext(tt.ctx))
			result := rr.Body.String()
			assert.Equal(t, rr.Code, tt.statusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGetAvailableDrones(t *testing.T) {
	type mocks struct {
		droneRepository      repository.DroneRepo
		medicationRepository repository.MedicationRepo
	}
	tests := []struct {
		name       string
		ctx        context.Context
		want       string
		statusCode int
		mocks      mocks
	}{
		{
			name:       "[Test] successful available drones retrieval",
			ctx:        context.Background(),
			want:       `[{"id":0,"serial_number":"SN","model":"Lightweight","weight_limit":40,"battery_capacity":70,"state":"DELIVERED","medications":null}]`,
			statusCode: http.StatusOK,
			mocks:      mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := usecase.NewDroneUsecase(tt.mocks.droneRepository, tt.mocks.medicationRepository)
			droneAPIs := NewDroneAPIs(droneUsecase)
			req, err := http.NewRequest("GET", "", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(droneAPIs.GetAvailableDrones)
			handler.ServeHTTP(rr, req.WithContext(tt.ctx))
			result := rr.Body.String()
			assert.Equal(t, rr.Code, tt.statusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestLoadDrone(t *testing.T) {
	type mocks struct {
		droneRepository      repository.DroneRepo
		medicationRepository repository.MedicationRepo
	}
	tests := []struct {
		name       string
		ctx        context.Context
		reqBody    []byte
		want       string
		statusCode int
		mocks      mocks
	}{
		{
			name:       "[Test] successful drone load",
			ctx:        context.Background(),
			reqBody:    []byte(`{"drone_id":1,"medication_id":1, "loaded": false}`),
			want:       `{"id":0,"serial_number":"","model":"","weight_limit":0,"battery_capacity":0,"state":"LOADING","medications":null}`,
			statusCode: http.StatusOK,
			mocks:      mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:       "[Test] drone load should return error if drone id is not provided",
			ctx:        context.Background(),
			reqBody:    []byte(`{"medication_id":1, "loaded": false}`),
			want:       `{"error":"drone id is required"}`,
			statusCode: http.StatusBadRequest,
			mocks:      mocks{droneRepository: mock.NewSuccessMockedDrone(), medicationRepository: mock.NewMockedMedication()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := usecase.NewDroneUsecase(tt.mocks.droneRepository, tt.mocks.medicationRepository)
			droneAPIs := NewDroneAPIs(droneUsecase)
			req, err := http.NewRequest("POST", "", bytes.NewBuffer(tt.reqBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(droneAPIs.Load)
			handler.ServeHTTP(rr, req.WithContext(tt.ctx))
			result := rr.Body.String()
			assert.Equal(t, rr.Code, tt.statusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}

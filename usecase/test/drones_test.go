package test

import (
	"context"
	"encoding/json"
	"github/Shimaa-Ibrahim/grones/repository/entity"
	"github/Shimaa-Ibrahim/grones/repository/mocks"
	"github/Shimaa-Ibrahim/grones/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDroneSuccessfulRegisteration(t *testing.T) {
	droneBytes := []byte(`{
		"seriel_number": "ldrefmweoflmj956flfrv2",
		"drone_model_id": 1,
		"weight_limit": 500,
		"battery_capacity": 25,
		"drone_state_id": 1
		
	}`)

	mockedDroneRepository := mocks.NewMockedDroneRepository()
	droneUseCase := usecase.NewDroneUseCase(mockedDroneRepository)
	t.Run("Test successful drone registeration", func(t *testing.T) {
		retrievedDroneBytes, err := droneUseCase.RegisterDrone(context.Background(), droneBytes)
		if err != nil {
			t.Errorf("[Error] Cannot create drone: %v", err)
		}
		var drone entity.Drone
		if err := json.Unmarshal(droneBytes, &drone); err != nil {
			t.Fatalf("[Error] Cannot  Unmarshal: %v", err)
		}

		var registeredDrone entity.Drone
		if err := json.Unmarshal(retrievedDroneBytes, &registeredDrone); err != nil {
			t.Fatalf("[Error] Cannot  Unmarshal: %v", err)
		}
		assert.Equal(t, registeredDrone, drone)

	})
}

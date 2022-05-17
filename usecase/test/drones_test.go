package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github/Shimaa-Ibrahim/grones/repository"
	repoEntity "github/Shimaa-Ibrahim/grones/repository/entity"
	"github/Shimaa-Ibrahim/grones/repository/mocks"
	"github/Shimaa-Ibrahim/grones/usecase"
	"github/Shimaa-Ibrahim/grones/usecase/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockedDroneRepository repository.DroneRepoProto
var droneUseCase usecase.DroneUseCaseProto

func init() {
	mockedDroneRepository = mocks.NewMockedDroneRepository()
	droneUseCase = usecase.NewDroneUseCase(mockedDroneRepository)
}

func TestDroneSuccessfulRegisteration(t *testing.T) {
	droneBytes := []byte(`{
		"seriel_number": "ldrefmweoflmj956flfrv2",
		"drone_model_id": 1,
		"weight_limit": 500,
		"battery_capacity": 25,
		"drone_state_id": 1
		
	}`)
	t.Run("Test successful drone registeration", func(t *testing.T) {
		retrievedDroneBytes, err := droneUseCase.RegisterDrone(context.Background(), droneBytes)
		if err != nil {
			t.Errorf("[Error] Cannot create drone: %v", err)
		}
		var drone repoEntity.Drone
		if err := json.Unmarshal(droneBytes, &drone); err != nil {
			t.Fatalf("[Error] Cannot  Unmarshal: %v", err)
		}

		var registeredDrone repoEntity.Drone
		if err := json.Unmarshal(retrievedDroneBytes, &registeredDrone); err != nil {
			t.Fatalf("[Error] Cannot  Unmarshal: %v", err)
		}
		assert.Equal(t, registeredDrone, drone)

	})
}

func TestCheckingLoadedItemsForGivenDrone(t *testing.T) {
	dronesIDs := []uint{1, 2}
	var expectDrones []repoEntity.Drone
	for _, id := range dronesIDs {
		drone, _ := mockedDroneRepository.GetByID(context.Background(), id)
		expectDrones = append(expectDrones, drone)
	}
	for index, id := range dronesIDs {
		t.Run(fmt.Sprintf("Test retrive right medication data for given drone of id %v", id), func(t *testing.T) {
			retievedDrone, err := droneUseCase.CheckDroneLoadedItem(context.Background(), id)
			if err != nil {
				t.Errorf("[Error] Cannot retrieve drone data: %v", err)
			}
			drone := repoEntity.Drone{}
			if err := json.Unmarshal(retievedDrone, &drone); err != nil {
				t.Fatalf("[Error] Cannot unmarshal drone data: %v", err)
			}
			assert.Equal(t, drone, expectDrones[index])
		})
	}

}

func TestCheckingAvailableDronesForLoading(t *testing.T) {
	expectAvailableDrones, _ := mockedDroneRepository.GetDronesAvailableForLoading(context.Background())
	retrievedDrones, err := droneUseCase.GetDronesAvailableForLoading(context.Background())
	if err != nil {
		t.Errorf("[Error] Cannot retrieve available drones: %v", err)
	}
	drones := []repoEntity.Drone{}
	if err := json.Unmarshal(retrievedDrones, &drones); err != nil {
		t.Fatalf("[Error] Cannot unmarshal drone data: %v", err)
	}
	assert.Equal(t, drones, expectAvailableDrones)

}

func TestCheckingDroneBatteryLevel(t *testing.T) {
	expectedBatteryLevels := []uint64{80, 90, 20}
	dronesIDs := []uint{1, 2, 3}
	for index, id := range dronesIDs {
		t.Run(fmt.Sprintf("Test retrive battery level for given drone of id %v", id), func(t *testing.T) {
			retrivedBatteryLevel, err := droneUseCase.CheckDroneBatteryLevel(context.Background(), id)
			if err != nil {
				t.Errorf("[Error] Cannot retrieve batteryLevel: %v", err)
			}
			batteryLevel := entity.BatteryLevelResponse{}
			if err := json.Unmarshal(retrivedBatteryLevel, &batteryLevel); err != nil {
				t.Fatalf("[Error] Cannot unmarshal batteryLevel data: %v", err)
			}
			assert.Equal(t, batteryLevel.BatteryLevel, expectedBatteryLevels[index])
		})
	}

}

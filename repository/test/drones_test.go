package test

import (
	"context"
	"github/Shimaa-Ibrahim/grones/repository"
	"github/Shimaa-Ibrahim/grones/repository/db"
	"github/Shimaa-Ibrahim/grones/repository/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDroneSuccessfulCreation(t *testing.T) {
	dbClient, err := db.ConnectToDB(TEST_DRONE_DATABASE)
	if err != nil {
		t.Skipf("[Error] failed to connect to database: %v", err)
	}
	TruncateDB(dbClient)
	droneModels := []entity.DroneModel{{Name: "modelOne"}, {Name: "modelTwo"}}
	if err := dbClient.Create(&droneModels).Error; err != nil {
		t.Errorf("[Error] Cannot create droneModels: %v", err)
	}
	droneStates := []entity.DroneState{{Name: "stateOne"}, {Name: "stateTwo"}}
	if err := dbClient.Create(&droneStates).Error; err != nil {
		t.Errorf("[Error] Cannot create droneStates: %v", err)
	}
	droneRepository := repository.NewDroneRepository(dbClient)
	type args struct {
		ctx   context.Context
		drone entity.Drone
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "[Test 1] test drone creation with correct data when calling dronerepo create method",
			args: args{
				ctx: context.Background(),
				drone: entity.Drone{
					SerielNumber:    generateRandomText(14),
					DroneModelID:    droneModels[0].ID,
					WeightLimit:     400,
					BatteryCapacity: 60,
					DroneStateID:    droneStates[0].ID,
				},
			},
		},
		{
			name: "[Test 2] test drone creation with correct data when calling dronerepo create method",
			args: args{
				ctx: context.Background(),
				drone: entity.Drone{
					SerielNumber:    generateRandomText(14),
					DroneModelID:    droneModels[1].ID,
					WeightLimit:     500,
					BatteryCapacity: 90,
					DroneStateID:    droneStates[1].ID,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdDrone, err := droneRepository.Create(tt.args.ctx, tt.args.drone)
			if err != nil {
				t.Errorf("[Error] Cannot create drone: %v", err)
			}
			drone := entity.Drone{}
			dbClient.Find(&drone, "id=?", createdDrone.ID)
			assert.Equal(t, drone.SerielNumber, tt.args.drone.SerielNumber)
			assert.Equal(t, drone.DroneModelID, tt.args.drone.DroneModelID)
			assert.Equal(t, drone.WeightLimit, tt.args.drone.WeightLimit)
			assert.Equal(t, drone.BatteryCapacity, tt.args.drone.BatteryCapacity)
			assert.Equal(t, drone.DroneStateID, tt.args.drone.DroneStateID)
		})
	}

}

func TestRightDroneDataRetrivalUsingID(t *testing.T) {
	dbClient, err := db.ConnectToDB(TEST_DRONE_DATABASE)
	if err != nil {
		t.Skipf("[Error] failed to connect to database: %v", err)
	}
	TruncateDB(dbClient)
	droneModels := []entity.DroneModel{{Name: "modelOne"}, {Name: "modelTwo"}}
	if err := dbClient.Create(&droneModels).Error; err != nil {
		t.Errorf("[Error] Cannot create droneModels: %v", err)
	}
	droneStates := []entity.DroneState{{Name: "stateOne"}, {Name: "stateTwo"}}
	if err := dbClient.Create(&droneStates).Error; err != nil {
		t.Errorf("[Error] Cannot create droneStates: %v", err)
	}
	drones := []entity.Drone{
		{
			SerielNumber:    generateRandomText(20),
			DroneModelID:    droneModels[0].ID,
			WeightLimit:     400,
			BatteryCapacity: 60,
			DroneStateID:    droneStates[0].ID,
		},
		{
			SerielNumber:    generateRandomText(21),
			DroneModelID:    droneModels[1].ID,
			WeightLimit:     500,
			BatteryCapacity: 90,
			DroneStateID:    droneStates[1].ID,
		},
	}
	if err := dbClient.Create(&drones).Error; err != nil {
		t.Errorf("[Error] Cannot create drones: %v", err)
	}
	var medications = []entity.Medication{
		{
			Name:    "med1",
			Code:    generateRandomText(12),
			Weight:  50,
			DroneID: drones[0].ID,
		},
		{
			Name:    "med2",
			Code:    generateRandomText(12),
			Weight:  100,
			DroneID: drones[0].ID,
		},
		{
			Name:    "med3",
			Code:    generateRandomText(12),
			Weight:  150,
			DroneID: drones[1].ID,
		},
		{
			Name:    "med4",
			Code:    generateRandomText(12),
			Weight:  200,
			DroneID: drones[1].ID,
		},
	}

	if result := dbClient.Create(&medications); result.Error != nil {
		t.Errorf("[Error] Cannot create medications: %v", err)
	}
	var medicationsIDs []uint
	for _, med := range medications {
		medicationsIDs = append(medicationsIDs, med.ID)
	}
	droneRepository := repository.NewDroneRepository(dbClient)
	type args struct {
		ctx context.Context
		ID  uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "[Test 1] test retrieve correct drone data when calling dronerepo GetByID method",
			args: args{
				ctx: context.Background(),
				ID:  drones[0].ID,
			},
		},
		{
			name: "[Test 2] test retrieve correct drone data when calling dronerepo GetByID method",
			args: args{
				ctx: context.Background(),
				ID:  drones[1].ID,
			},
		},
	}

	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drone, err := droneRepository.GetByID(tt.args.ctx, tt.args.ID)
			if err != nil {
				t.Errorf("[Error] Cannot retrive drone data: %v", err)
			}
			assert.Equal(t, drone.ID, drones[index].ID)
			assert.Equal(t, drone.SerielNumber, drones[index].SerielNumber)
			assert.Equal(t, drone.DroneModelID, drones[index].DroneModelID)
			assert.Equal(t, drone.WeightLimit, drones[index].WeightLimit)
			assert.Equal(t, drone.BatteryCapacity, drones[index].BatteryCapacity)
			assert.Equal(t, drone.DroneStateID, drones[index].DroneStateID)
			// deepEqual
			assert.Equal(t, drone.DroneModel, droneModels[index])
			assert.Equal(t, drone.DroneState, droneStates[index])
			var expextedMeds []entity.Medication
			dbClient.Find(&expextedMeds, medicationsIDs[index*2:index*2+2])
			assert.Equal(t, drone.Medications, expextedMeds)

		})
	}

}

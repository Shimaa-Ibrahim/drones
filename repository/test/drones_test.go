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
	droneModels := []entity.DroneModel{{Name: generateRandomText(20)}, {Name: generateRandomText(20)}}
	droneStates := []entity.DroneState{{Name: generateRandomText(20)}, {Name: generateRandomText(20)}}
	dbClient, err := db.ConnectToDB(TEST_DRONE_DATABASE)
	if err != nil {
		t.Skipf("[Error] failed to connect to database: %v", err)
	}
	TruncateDB(dbClient)
	if err := dbClient.Create(&droneModels).Error; err != nil {
		t.Errorf("[Error] Cannot create droneModels: %v", err)
	}
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
					SerielNumber:    "khiehbe0473b84n",
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
					SerielNumber:    "sdhi8453jble9y0",
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

package test

import (
	"context"
	"fmt"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/db"
	"github/Shimaa-Ibrahim/drones/repository/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadingDroneWithMedicationsBatchUpdates(t *testing.T) {
	dbClient, err := db.ConnectToDB(TEST_DRONE_DATABASE)
	if err != nil {
		t.Fatalf("[Error] failed to connect to database: %v", err)
	}
	TruncateDB(dbClient)
	drone := entity.Drone{
		SerielNumber:    generateRandomText(20),
		DroneModel:      entity.DroneModel{Name: generateRandomText(20)},
		WeightLimit:     200,
		BatteryCapacity: 40,
		DroneState:      entity.DroneState{Name: generateRandomText(20)},
	}

	if err := dbClient.Create(&drone).Error; err != nil {
		t.Errorf("[Error] Cannot create drone: %v", err)
	}
	var medications = []entity.Medication{
		{
			Name:   "med1",
			Code:   generateRandomText(12),
			Weight: 50,
		},
		{
			Name:   "med2",
			Code:   generateRandomText(12),
			Weight: 100,
		},
		{
			Name:   "med3",
			Code:   generateRandomText(12),
			Weight: 150,
		},
		{
			Name:   "med4",
			Code:   generateRandomText(12),
			Weight: 200,
		},
	}

	if result := dbClient.Omit("DroneID").Create(&medications); result.Error != nil {
		t.Errorf("[Error] Cannot create medications: %v", err)
	}

	var medicationsIDs []uint
	for _, med := range medications {
		medicationsIDs = append(medicationsIDs, med.ID)
	}
	medicationRepository := repository.NewMedicationRepository(dbClient)

	t.Run("Test update selected medications with given drone id", func(t *testing.T) {
		err := medicationRepository.UpdateMedicationsWithDroneID(context.Background(), drone.ID, medicationsIDs[1:])
		if err != nil {
			t.Errorf("[Error] Cannot update medications: %v", err)
		}

		retrievedMeds := []entity.Medication{}
		dbClient.Find(&retrievedMeds, medicationsIDs)

		assert.Equal(t, retrievedMeds[0].DroneID, uint(0))
		for index, med := range retrievedMeds[1:] {
			assert.Equal(t, med.ID, medicationsIDs[index+1])
			assert.Equal(t, med.DroneID, drone.ID)
		}

	})

}

func TestRightMedicationDataRetrivalUsingID(t *testing.T) {
	dbClient, err := db.ConnectToDB(TEST_DRONE_DATABASE)
	if err != nil {
		t.Fatalf("[Error] failed to connect to database: %v", err)
	}
	TruncateDB(dbClient)
	var medications = []entity.Medication{
		{
			Name:   "med1",
			Code:   generateRandomText(12),
			Weight: 50,
		},
		{
			Name:   "med2",
			Code:   generateRandomText(12),
			Weight: 100,
		},
		{
			Name:   "med3",
			Code:   generateRandomText(12),
			Weight: 150,
		},
		{
			Name:   "med4",
			Code:   generateRandomText(12),
			Weight: 200,
		},
	}

	if result := dbClient.Omit("DroneID").Create(&medications); result.Error != nil {
		t.Errorf("[Error] Cannot create medications: %v", err)
	}

	var medicationsIDs []uint
	for _, med := range medications {
		medicationsIDs = append(medicationsIDs, med.ID)
	}

	medicationRepository := repository.NewMedicationRepository(dbClient)
	var expextedMeds []entity.Medication
	dbClient.Find(&expextedMeds, medicationsIDs)

	for index, medication := range expextedMeds {
		t.Run(fmt.Sprintf("Test right data retrival for medication %v", index), func(t *testing.T) {
			retrievedMedication, err := medicationRepository.GetByID(context.Background(), medication.ID)
			if err != nil {
				t.Errorf("[Error] Cannot retrive medication data: %v", err)
			}
			assert.Equal(t, retrievedMedication, medication)

		})
	}

}

func TestMedicationSuccessfulCreation(t *testing.T) {
	dbClient, err := db.ConnectToDB(TEST_DRONE_DATABASE)
	if err != nil {
		t.Fatalf("[Error] failed to connect to database: %v", err)
	}
	TruncateDB(dbClient)
	medicationRepository := repository.NewMedicationRepository(dbClient)
	var medications = []entity.Medication{
		{
			Name:   "med1",
			Code:   generateRandomText(12),
			Weight: 50,
		},
		{
			Name:   "med2",
			Code:   generateRandomText(12),
			Weight: 100,
		},
	}

	for index, medication := range medications {
		t.Run(fmt.Sprintf("Test right data retrival for medication %v", index), func(t *testing.T) {
			createdMedication, err := medicationRepository.CreateMedication(context.Background(), medication)
			if err != nil {
				t.Errorf("[Error] Cannot Create medication: %v", err)
			}
			retrievedMedication := entity.Medication{}
			assert.NotEmpty(t, createdMedication.ID)
			if err := dbClient.Find(&retrievedMedication, createdMedication.ID).Error; err != nil {
				t.Errorf("[Error] Cannot retrive medication data: %v", err)
			}
			assert.Equal(t, retrievedMedication.Name, createdMedication.Name)
			assert.Equal(t, retrievedMedication.Code, createdMedication.Code)
			assert.Equal(t, retrievedMedication.Weight, createdMedication.Weight)
		})
	}

}

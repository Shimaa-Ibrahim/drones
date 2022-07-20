package repository

import (
	"context"
	"github/Shimaa-Ibrahim/drones/repository/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMedication(t *testing.T) {
	clearDataBase()
	medicationRepository := NewMedicationRepository(dbClient)
	medication := entity.Medication{Code: "don't duplicate me"}
	if err := dbClient.Create(&medication).Error; err != nil {
		t.Errorf("[Error] medication creation: %v", err)
	}
	type args struct {
		ctx        context.Context
		medication entity.Medication
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Medication
		wantErr string
	}{
		{
			name: "[Test] succssful medication creation",
			args: args{context.Background(), entity.Medication{Code: "C-20", Name: "Medication 1", Weight: 10, ImagePath: "image_path"}},
			want: entity.Medication{Code: "C-1", Name: "Medication 1", Weight: 10, ImagePath: "image_path"},
		},
		{
			name:    "[Test] create medication should fail if code is duplicate",
			args:    args{context.Background(), entity.Medication{Code: "don't duplicate me"}},
			wantErr: `ERROR: duplicate key value violates unique constraint "medications_code_key" (SQLSTATE 23505)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := medicationRepository.Create(tt.args.ctx, tt.args.medication)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			got := entity.Medication{}
			err = dbClient.Find(&got, "code", tt.args.medication.Code).Error
			if err != nil {
				t.Errorf("[Error] medication retrieval: %v", err)
			}
			checkMedicationsEquality(t, tt.args.medication, got)
		})
	}
}

func TestGetMedication(t *testing.T) {
	clearDataBase()
	medicationRepository := NewMedicationRepository(dbClient)
	medications := []entity.Medication{
		{Code: "C-1", Name: "Medication 1", Weight: 10, ImagePath: "image_path"},
		{Code: "C-2", Name: "Medication 2", Weight: 20, ImagePath: "image_path", Drones: []entity.Drone{{SerialNumber: "S-1"}}},
		{Code: "C-3", Name: "Medication 3", Weight: 40, ImagePath: "image_path", Drones: []entity.Drone{{SerialNumber: "S-2"}}},
	}
	if err := dbClient.Create(&medications).Error; err != nil {
		t.Errorf("[Error] medication creation: %v", err)
	}
	if err := dbClient.Model(&entity.DroneMedication{}).Where("drone_id=? AND medication_id=?", medications[2].Drones[0].ID, medications[2].ID).Update("loaded", false).Error; err != nil {
		t.Errorf("[Error] drone medication update: %v", err)
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Medication
		wantErr string
	}{
		{
			name: "[Test] succssful medication retrieval",
			args: args{context.Background(), medications[1].ID},
			want: medications[1],
		},
		{
			name: "[Test] get medication should return medication with empty drones",
			args: args{context.Background(), medications[2].ID},
			want: entity.Medication{Code: "C-3", Name: "Medication 3", Weight: 40, ImagePath: "image_path", Drones: []entity.Drone{}},
		},
		{
			name:    "[Test] get medication should fail if medication doesn't exist",
			args:    args{context.Background(), 0},
			wantErr: "record not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := medicationRepository.Get(tt.args.ctx, tt.args.id)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			checkMedicationsEquality(t, tt.want, got)
		})
	}
}

func TestGetAllMedication(t *testing.T) {
	clearDataBase()
	medicationRepository := NewMedicationRepository(dbClient)
	medications := []entity.Medication{
		{Code: "C-1", Name: "Medication 1", Weight: 10, ImagePath: "image_path"},
		{Code: "C-2", Name: "Medication 2", Weight: 20, ImagePath: "image_path"},
		{Code: "C-3", Name: "Medication 3", Weight: 30, ImagePath: "image_path"},
	}
	if err := dbClient.Create(&medications).Error; err != nil {
		t.Errorf("[Error] medication creation: %v", err)
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Medication
		wantErr string
	}{
		{
			name: "[Test] succssful medication retrieval",
			args: args{context.Background()},
			want: medications,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := medicationRepository.GetAll(tt.args.ctx)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			for i, medication := range got {
				checkMedicationsEquality(t, tt.want[i], medication)
			}
		})
	}
}

func checkMedicationsEquality(t *testing.T, want entity.Medication, got entity.Medication) {
	t.Helper()
	assert.NotEmpty(t, got.ID)
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Code, got.Code)
	assert.Equal(t, want.Weight, got.Weight)
	assert.Equal(t, want.ImagePath, got.ImagePath)
	assert.Equal(t, len(want.Drones), len(got.Drones))
	for i, drone := range got.Drones {
		assert.NotEmpty(t, drone.ID)
		assert.Equal(t, want.Drones[i].SerialNumber, drone.SerialNumber)
	}
}

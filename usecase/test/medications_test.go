package test

import (
	"context"
	"encoding/json"
	"github/Shimaa-Ibrahim/drones/repository/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailureToLoadDroneWithMedicationItems(t *testing.T) {
	type args struct {
		ctx context.Context
		req []byte
	}
	tests := []struct {
		name      string
		args      args
		wantedErr string
	}{
		{
			name: "Test drone loading failure when drone does not exist",
			args: args{
				ctx: context.Background(),
				req: []byte(`{"id": 0, "medications_ids": [1,2]}`),
			},
			wantedErr: "The drone does not exist",
		},
		{
			name: "Test drone loading failure when the drone is not in loading state",
			args: args{
				ctx: context.Background(),
				req: []byte(`{"id": 1, "medications_ids": [1,2]}`),
			},
			wantedErr: "The drone is not available",
		},
		{
			name: "Test drone loading failure when the drone's battery level below 25",
			args: args{
				ctx: context.Background(),
				req: []byte(`{"id": 3, "medications_ids": [1,2]}`),
			},
			wantedErr: "the drone's battery level below 25",
		},
		{
			name: "Test drone loading failure when one of medications does not exist",
			args: args{
				ctx: context.Background(),
				req: []byte(`{"id": 2, "medications_ids": [0,1,2]}`),
			},
			wantedErr: "These medications do not exist: [0]",
		},
		{
			name: "Test drone loading failure when one of medications already loaded",
			args: args{
				ctx: context.Background(),
				req: []byte(`{"id": 2, "medications_ids": [1,2]}`),
			},
			wantedErr: "These medications are already loaded: [2]",
		},
		{
			name: "Test drone loading failure when the medications weight more than the drone's weight limit ",
			args: args{
				ctx: context.Background(),
				req: []byte(`{"id": 2, "medications_ids": [1,3]}`),
			},
			wantedErr: "The medications weight exceeded the drone's weight limit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := medicationUseCase.LoadDroneWithMedicationItems(tt.args.ctx, tt.args.req)
			assert.EqualError(t, err, tt.wantedErr)
		})
	}

}

func TestMedicationSuccessfulRegisteration(t *testing.T) {
	medicationBytes := []byte(`{
		"name": "med",
		"code": "brt6l545eltgsl",
		"weight": 500
	}`)
	t.Run("Test successful medication registeration", func(t *testing.T) {
		createdMedicationBytes, err := medicationUseCase.RegisterMedication(context.Background(), medicationBytes, "imagePath")
		if err != nil {
			t.Errorf("[Error] Cannot create drone: %v", err)
		}
		medication := entity.Medication{}
		if err := json.Unmarshal(createdMedicationBytes, &medication); err != nil {
			t.Fatalf("[Error] Cannot  Unmarshal: %v", err)
		}
		assert.Equal(t, medication.Name, "med")
		assert.Equal(t, medication.Code, "brt6l545eltgsl")
		assert.Equal(t, medication.Weight, float64(500))
		assert.Equal(t, medication.ImagePath, "imagePath")
	})
}

package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ErrorMsg struct {
	Error string `json:"error"`
}

func TestFailureToLoadDroneWithMedicationItemsAPI(t *testing.T) {
	apiURL := "/medications/load/"
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
				req: []byte(`{"id": 0, "medications_ids": [1,2]}`),
			},
			wantedErr: "The drone does not exist",
		},
		{
			name: "Test drone loading failure when the drone is not in loading state",
			args: args{
				req: []byte(`{"id": 1, "medications_ids": [1,2]}`),
			},
			wantedErr: "The drone is not available",
		},
		{
			name: "Test drone loading failure when the drone's battery level below 25",
			args: args{
				req: []byte(`{"id": 3, "medications_ids": [1,2]}`),
			},
			wantedErr: "the drone's battery level below 25",
		},
		{
			name: "Test drone loading failure when one of medications does not exist",
			args: args{
				req: []byte(`{"id": 2, "medications_ids": [0,1,2]}`),
			},
			wantedErr: "These medications do not exist: [0]",
		},
		{
			name: "Test drone loading failure when one of medications already loaded",
			args: args{
				req: []byte(`{"id": 2, "medications_ids": [1,2]}`),
			},
			wantedErr: "These medications are already loaded: [2]",
		},
		{
			name: "Test drone loading failure when the medications weight more than the drone's weight limit ",
			args: args{
				req: []byte(`{"id": 2, "medications_ids": [1,3]}`),
			},
			wantedErr: "The medications weight exceeded the drone's weight limit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(tt.args.req))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(medicationAPI.LoadDroneWithMedicationItems)
			handler.ServeHTTP(rr, req)
			result := rr.Body.String()
			if status := rr.Code; status != http.StatusUnprocessableEntity {
				t.Errorf("Wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
			}
			errMsg := ErrorMsg{}
			if err := json.Unmarshal([]byte(result), &errMsg); err != nil {
				t.Fatalf("[Error] Cannot  Unmarshal: %v", err)
			}

			assert.Equal(t, errMsg.Error, tt.wantedErr)
		})
	}
}

func TestSuccssfulLoadingDroneWithMedicationItemsAPI(t *testing.T) {
	apiURL := "/medications/load/"
	data := []byte(`{"id": 2, "medications_ids": [1]}`)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(medicationAPI.LoadDroneWithMedicationItems)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}
}

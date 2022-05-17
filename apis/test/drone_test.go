package test

import (
	"bytes"
	"encoding/json"
	"github/Shimaa-Ibrahim/grones/apis"
	"github/Shimaa-Ibrahim/grones/repository"
	"github/Shimaa-Ibrahim/grones/repository/entity"
	"github/Shimaa-Ibrahim/grones/repository/mocks"
	"github/Shimaa-Ibrahim/grones/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var droneRepository repository.DroneRepoProto
var droneUseCase usecase.DroneUseCaseProto
var droneAPI apis.DroneAPIs

func init() {
	droneRepository = mocks.NewMockedDroneRepository()
	droneUseCase = usecase.NewDroneUseCase(droneRepository)
	droneAPI = apis.NewDroneAPI(droneUseCase)

}

func TestDroneRegisterTroughAPI(t *testing.T) {
	apiURL := "/drone/register/"
	newDrone := []byte(`{
		"serial_number": "serial1320number",
		"drone_model_id": 15,
		"weight_limit": 500,
		"battery_capacity": 75,
		"drone_state_id": 1
	}`)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(newDrone))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(droneAPI.RegisterDrone)
	handler.ServeHTTP(rr, req)
	result := rr.Body.String()
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusCreated)
	}
	drone := entity.Drone{}
	if err := json.Unmarshal([]byte(result), &drone); err != nil {
		t.Fatalf("[Error] Cannot  Unmarshal: %v", err)
	}

	assert.Equal(t, drone.SerielNumber, "serial1320number")
	assert.Equal(t, drone.DroneModelID, uint(15))
	assert.Equal(t, drone.WeightLimit, uint64(500))
	assert.Equal(t, drone.BatteryCapacity, uint64(75))
	assert.Equal(t, drone.DroneStateID, uint(1))
}

func TestCheckDroneLoadedItemTroughAPI(t *testing.T) {
	apiURL := "/drone/checkdroneloaded/1/"
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/drone/checkdroneloaded/{id}/", droneAPI.CheckDroneLoadedItem)
	router.ServeHTTP(rr, req)
	result := rr.Body.String()
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}
	drone := entity.Drone{}
	if err := json.Unmarshal([]byte(result), &drone); err != nil {
		t.Fatalf("[Error] Cannot  Unmarshal: %v", err)
	}

	expectedDrone, err := droneRepository.GetByID(req.Context(), uint(1))
	if err := json.Unmarshal([]byte(result), &drone); err != nil {
		t.Fatalf("[Error] Cannot  retrieve expected data: %v", err)
	}

	assert.Equal(t, drone, expectedDrone)
}

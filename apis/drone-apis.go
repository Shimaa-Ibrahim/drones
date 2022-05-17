package apis

import (
	"fmt"
	"github/Shimaa-Ibrahim/grones/usecase"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DroneAPIs struct {
	droneUseCase usecase.DroneUseCaseProto
}

func NewDroneAPI(droneUseCase usecase.DroneUseCaseProto) DroneAPIs {
	return DroneAPIs{
		droneUseCase: droneUseCase,
	}
}

func (api DroneAPIs) RegisterDrone(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}

	response, err := api.droneUseCase.RegisterDrone(ctx, requestByte)
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (api DroneAPIs) CheckDroneLoadedItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	pathID := vars["id"]
	id, err := strconv.Atoi(pathID)
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}

	response, err := api.droneUseCase.CheckDroneLoadedItem(ctx, uint(id))
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

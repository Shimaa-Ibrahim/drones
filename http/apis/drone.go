package apis

import (
	"errors"
	"github/Shimaa-Ibrahim/drones/usecase"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DroneAPIs struct {
	droneUsecase usecase.IDroneUsecase
}

func NewDroneAPIs(droneUsecase usecase.IDroneUsecase) DroneAPIs {
	return DroneAPIs{droneUsecase: droneUsecase}
}

func (api DroneAPIs) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}
	response, err := api.droneUsecase.Create(ctx, requestByte)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (api DroneAPIs) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.Atoi(id)
	if err != nil || ID <= 0 {
		err := errors.New("invalid ID")
		sendError(w, http.StatusBadRequest, err)
		return
	}
	response, err := api.droneUsecase.Get(ctx, uint(ID))
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (api DroneAPIs) GetAvailableDrones(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response, err := api.droneUsecase.GetAvailableDrones(ctx)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (api DroneAPIs) Load(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}
	response, err := api.droneUsecase.Load(ctx, requestByte)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

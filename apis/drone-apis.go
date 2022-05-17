package apis

import (
	"fmt"
	"github/Shimaa-Ibrahim/grones/usecase"
	"io/ioutil"
	"net/http"
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}

	responsePayload, err := api.droneUseCase.RegisterDrone(ctx, requestByte)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responsePayload)
}

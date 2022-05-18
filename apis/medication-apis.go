package apis

import (
	"fmt"
	"github/Shimaa-Ibrahim/grones/usecase"
	"io/ioutil"
	"log"
	"net/http"
)

type MedicationAPIs struct {
	medicationUseCase usecase.MedicationUseCaseProto
}

func NewMedicationAPI(medicationUseCase usecase.MedicationUseCaseProto) MedicationAPIs {
	return MedicationAPIs{
		medicationUseCase: medicationUseCase,
	}
}

func (api MedicationAPIs) LoadDroneWithMedicationItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}

	err = api.medicationUseCase.LoadDroneWithMedicationItems(ctx, requestByte)
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusUnprocessableEntity)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
}

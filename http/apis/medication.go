package apis

import (
	"errors"
	"github/Shimaa-Ibrahim/drones/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MedicationAPIs struct {
	medicationUsecase usecase.IMedicationUsecase
}

func NewMedicationAPIs(medicationUseCase usecase.IMedicationUsecase) MedicationAPIs {
	return MedicationAPIs{medicationUsecase: medicationUseCase}
}

func (api MedicationAPIs) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(5 << 20)
	filename := ""
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		if err != http.ErrMissingFile {
			sendError(w, http.StatusInternalServerError, err)
			return
		}
	}
	if file != nil {
		defer file.Close()
		filename = fileHeader.Filename
	}
	requestBody := []byte(r.FormValue("medication"))
	response, err := api.medicationUsecase.Create(ctx, requestBody, file, filename)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (api MedicationAPIs) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.Atoi(id)
	if err != nil || ID <= 0 {
		err := errors.New("invalid ID")
		sendError(w, http.StatusBadRequest, err)
		return
	}
	response, err := api.medicationUsecase.Get(ctx, uint(ID))
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (api MedicationAPIs) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response, err := api.medicationUsecase.GetAll(ctx)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

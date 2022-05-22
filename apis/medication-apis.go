package apis

import (
	"fmt"
	"github/Shimaa-Ibrahim/drones/usecase"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

func (api MedicationAPIs) RegisterMedication(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(4 * 1024 * 1024)
	// image
	file, handle, err := r.FormFile("image")
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}
	defer file.Close()

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}

	imagePath := fmt.Sprintf("%v-%v", time.Now().Format(time.RFC3339), handle.Filename)
	dst, err := os.Create(fmt.Sprintf("./uploads/%s", imagePath))
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{[Error]: %s}`, err.Error())))
		return
	}

	requestByte := []byte(r.FormValue("medication"))
	response, err := api.medicationUseCase.RegisterMedication(r.Context(), requestByte, imagePath)
	if err != nil {
		log.Printf("[Error]: %v ---- [status code] %v", err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

package server

import (
	"github/Shimaa-Ibrahim/drones/apis"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIs struct {
	DroneAPIs      apis.DroneAPIs
	MedicationAPIs apis.MedicationAPIs
}

func StartServer(apis APIs) {
	r := mux.NewRouter()
	r = r.StrictSlash(true)
	r.HandleFunc("/drone/register/", apis.DroneAPIs.RegisterDrone).Methods("POST")
	r.HandleFunc("/drone/checkdroneloaded/{id}/", apis.DroneAPIs.CheckDroneLoadedItem).Methods("GET")
	r.HandleFunc("/drone/availabledrones/", apis.DroneAPIs.GetDronesAvailableForLoading).Methods("GET")
	r.HandleFunc("/drone/checkbattery/{id}/", apis.DroneAPIs.CheckDroneBatteryLevel).Methods("GET")
	r.HandleFunc("/medications/load/", apis.MedicationAPIs.LoadDroneWithMedicationItems).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", r))
}

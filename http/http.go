package http

import (
	"github/Shimaa-Ibrahim/drones/http/apis"
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

	droneSubrouter := r.PathPrefix("/drone").Subrouter()
	droneSubrouter.HandleFunc("/", apis.DroneAPIs.Register).Methods("POST")
	droneSubrouter.HandleFunc("/available", apis.DroneAPIs.GetAvailableDrones).Methods("GET")
	droneSubrouter.HandleFunc("/{id}", apis.DroneAPIs.Get).Methods("GET")
	droneSubrouter.HandleFunc("/load", apis.DroneAPIs.Load).Methods("POST")

	medicationSubrouter := r.PathPrefix("/medication").Subrouter()
	medicationSubrouter.HandleFunc("/", apis.MedicationAPIs.Register).Methods("POST")
	medicationSubrouter.HandleFunc("/all", apis.MedicationAPIs.GetAll).Methods("GET")
	medicationSubrouter.HandleFunc("/{id}", apis.MedicationAPIs.Get).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

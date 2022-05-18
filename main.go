package main

import (
	"fmt"
	"github/Shimaa-Ibrahim/grones/apis"
	"github/Shimaa-Ibrahim/grones/repository"
	"github/Shimaa-Ibrahim/grones/repository/db"
	"github/Shimaa-Ibrahim/grones/server"
	"github/Shimaa-Ibrahim/grones/usecase"
	"log"
)

const DEV_DRONE_DATABASE = "DEV_DRONE_DATABASE"

func main() {
	fmt.Println("Hello World!...")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// db connection
	dbClient, err := db.ConnectToDB(DEV_DRONE_DATABASE)
	if err != nil {
		log.Fatalf("[Error] failed to connect to database: %v", err)
	}

	//repos
	droneRepository := repository.NewDroneRepository(dbClient)
	medicationRepository := repository.NewMedicationRepository(dbClient)

	//usecases
	droneUseCase := usecase.NewDroneUseCase(droneRepository)
	medicationUseCase := usecase.NewMedicationUseCase(medicationRepository, droneRepository)

	//apis
	droneAPIs := apis.NewDroneAPI(droneUseCase)
	medicationAPIs := apis.NewMedicationAPI(medicationUseCase)

	apis := server.APIs{
		DroneAPIs:      droneAPIs,
		MedicationAPIs: medicationAPIs,
	}

	server.StartServer(apis)
}

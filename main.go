package main

import (
	"context"
	"fmt"
	"github/Shimaa-Ibrahim/drones/http"
	"github/Shimaa-Ibrahim/drones/http/apis"
	"github/Shimaa-Ibrahim/drones/iofile"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/db"
	"github/Shimaa-Ibrahim/drones/usecase"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
)

const DEV_DRONE_DATABASE = "DEV_DRONE_DATABASE"

func main() {
	fmt.Println("Hello World!...")
	log.SetFlags(log.LstdFlags | log.Llongfile)
	dbClient, err := db.ConnectToDatabase(os.Getenv(DEV_DRONE_DATABASE))
	if err != nil {
		log.Fatalf("[Error] failed to connect to database: %v", err)
	}
	// scheduler
	scheduler := gocron.NewScheduler(time.UTC)

	ioFile := iofile.NewIOFile()
	//repos
	droneRepository := repository.NewDroneRepository(dbClient)
	medicationRepository := repository.NewMedicationRepository(dbClient)
	batteryLogRepository := repository.NewBatteryLogRepository(dbClient)

	//usecases
	droneUsecase := usecase.NewDroneUsecase(droneRepository, medicationRepository)
	medicationUsecase := usecase.NewMedicationUsecase(medicationRepository, ioFile)
	batteryLogsUsecase := usecase.NewBatterylogUsecase(droneRepository, batteryLogRepository, ioFile)

	//apis
	droneAPIs := apis.NewDroneAPIs(droneUsecase)
	medicationAPIs := apis.NewMedicationAPIs(medicationUsecase)

	apis := http.APIs{
		DroneAPIs:      droneAPIs,
		MedicationAPIs: medicationAPIs,
	}

	logDronesBatteryLevelTask := batteryLogsUsecase.Create(context.Background())
	scheduler.Every(1).Hour().Do(logDronesBatteryLevelTask)
	// periodic task to remove log files TODO

	scheduler.StartAsync()
	http.StartServer(apis)

}

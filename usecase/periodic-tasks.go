package usecase

import (
	"github/Shimaa-Ibrahim/grones/repository"
	"log"
)

type PeriodicTasksUseCaseProto interface {
	LogDronesBatteryLevel() error
}

type PeriodicTasksUseCase struct {
	droneRepo repository.DroneRepoProto
}

func NewPeriodicTasksUseCase(droneRepo repository.DroneRepoProto) PeriodicTasksUseCaseProto {
	return PeriodicTasksUseCase{
		droneRepo: droneRepo,
	}
}

func (p PeriodicTasksUseCase) LogDronesBatteryLevel() error {
	logsRecord, err := p.droneRepo.LogDronesBatteryLevel()
	// log to file TODO
	log.Println("battery logs ..................................")
	for index, logRecord := range logsRecord {
		log.Printf("[Record %v ]: %v\n", index, logRecord)
	}
	log.Println("battery logs ..................................")

	return err
}

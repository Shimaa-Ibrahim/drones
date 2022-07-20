package usecase

import (
	"context"
	"fmt"
	"github/Shimaa-Ibrahim/drones/iofile"
	"github/Shimaa-Ibrahim/drones/repository"
	repoEntity "github/Shimaa-Ibrahim/drones/repository/entity"
	"log"
	"time"
)

type IBatteryLogs interface {
	Create(ctx context.Context) error
}

type BatteryLogs struct {
	droneRepository       repository.DroneRepo
	BatteryLogsRepository repository.BatteryLogRepo
	iofile                iofile.IOFiles
}

func NewBatterylogUsecase(droneRepository repository.DroneRepo, batteryLogsRepository repository.BatteryLogRepo, iofile iofile.IOFiles) IBatteryLogs {
	return BatteryLogs{droneRepository: droneRepository, BatteryLogsRepository: batteryLogsRepository, iofile: iofile}
}

func (b BatteryLogs) Create(ctx context.Context) error {
	drones, err := b.droneRepository.GetAll(ctx)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return err
	}
	batteryLogs := []repoEntity.BatteryLog{}
	data := []string{}
	for _, drone := range drones {
		batteryLogs = append(batteryLogs, repoEntity.BatteryLog{DroneID: drone.ID, BatteryLevel: drone.BatteryCapacity})
		data = append(data, stringfyDroneBatteryLog(drone))
	}
	err = b.iofile.Write(data, fmt.Sprintf("./logs/battery-levels/%s-battery-levels.log", time.Now().Format(time.RFC3339)))
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return err
	}
	err = b.BatteryLogsRepository.Create(ctx, batteryLogs)
	return err
}

func stringfyDroneBatteryLog(drone repoEntity.Drone) string {
	return fmt.Sprintf(`Drone "%s" : battery Level %v`, drone.SerialNumber, drone.BatteryCapacity)
}

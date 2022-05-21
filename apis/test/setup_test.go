package test

import (
	"github/Shimaa-Ibrahim/drones/apis"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/mocks"
	"github/Shimaa-Ibrahim/drones/usecase"
)

var droneRepository repository.DroneRepoProto
var droneUseCase usecase.DroneUseCaseProto
var droneAPI apis.DroneAPIs
var medicationRepository repository.MedicationRepoProto
var medicationUseCase usecase.MedicationUseCaseProto
var medicationAPI apis.MedicationAPIs

func init() {
	droneRepository = mocks.NewMockedDroneRepository()
	droneUseCase = usecase.NewDroneUseCase(droneRepository)
	droneAPI = apis.NewDroneAPI(droneUseCase)
	medicationRepository = mocks.NewMockedMedicationRepository()
	medicationUseCase = usecase.NewMedicationUseCase(medicationRepository, droneRepository)
	medicationAPI = apis.NewMedicationAPI(medicationUseCase)
}

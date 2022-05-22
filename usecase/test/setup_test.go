package test

import (
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/mocks"
	"github/Shimaa-Ibrahim/drones/usecase"
)

var mockedDroneRepository repository.DroneRepoProto
var mockedMedicationRepository repository.MedicationRepoProto
var droneUseCase usecase.DroneUseCaseProto
var medicationUseCase usecase.MedicationUseCaseProto

func init() {
	mockedDroneRepository = mocks.NewMockedDroneRepository()
	mockedMedicationRepository := mocks.NewMockedMedicationRepository()

	droneUseCase = usecase.NewDroneUseCase(mockedDroneRepository)
	medicationUseCase = usecase.NewMedicationUseCase(mockedMedicationRepository, mockedDroneRepository)

}

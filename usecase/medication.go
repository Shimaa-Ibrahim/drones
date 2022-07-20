package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github/Shimaa-Ibrahim/drones/iofile"
	"github/Shimaa-Ibrahim/drones/repository"
	repoEntity "github/Shimaa-Ibrahim/drones/repository/entity"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type IMedicationUsecase interface {
	Create(ctx context.Context, request []byte, file multipart.File, filename string) ([]byte, error)
	Get(ctx context.Context, id uint) ([]byte, error)
	GetAll(ctx context.Context) ([]byte, error)
}

type MedicationUsecase struct {
	medicationRepository repository.MedicationRepo
	ioFile               iofile.IOFiles
}

func NewMedicationUsecase(medicationRepository repository.MedicationRepo, ioFile iofile.IOFiles) IMedicationUsecase {
	return MedicationUsecase{medicationRepository: medicationRepository, ioFile: ioFile}
}

func (m MedicationUsecase) Create(ctx context.Context, request []byte, file multipart.File, filename string) ([]byte, error) {
	medication := repoEntity.Medication{}
	err := json.Unmarshal(request, &medication)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	validStruct, err := govalidator.ValidateStruct(medication)
	if err != nil || !validStruct {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	if file != nil {
		Ext := strings.Split(filename, ".")
		FileMIMEType := m.ioFile.GetMIMEType(Ext[len(Ext)-1])
		if FileMIMEType != "image/jpeg" && FileMIMEType != "image/png" {
			err := errors.New("invalid file type")
			log.Printf("[Error]: %v\n", err)
			return []byte{}, err
		}
		imagePath := fmt.Sprintf("%v-%v", time.Now().Format(time.RFC3339), filename)
		err := m.ioFile.Save(file, "./uploads/"+imagePath)
		if err != nil {
			log.Printf("[Error]: %v\n", err)
			return []byte{}, err
		}
		medication.ImagePath = imagePath
	}

	if medication, err = m.medicationRepository.Create(ctx, medication); err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	return json.Marshal(medication)
}

func (m MedicationUsecase) Get(ctx context.Context, id uint) ([]byte, error) {
	medication, err := m.medicationRepository.Get(ctx, id)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	return json.Marshal(medication)
}

func (m MedicationUsecase) GetAll(ctx context.Context) ([]byte, error) {
	medications, err := m.medicationRepository.GetAll(ctx)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return []byte{}, err
	}
	return json.Marshal(medications)
}

package usecase

import (
	"context"
	"fmt"
	"github/Shimaa-Ibrahim/drones/iofile"
	mockedFile "github/Shimaa-Ibrahim/drones/iofile/mock"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/mock"
	"mime/multipart"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateMedication(t *testing.T) {
	type args struct {
		ctx      context.Context
		request  []byte
		file     multipart.File
		filename string
	}
	type mocks struct {
		medicationRepository repository.MedicationRepo
		iofile               iofile.IOFiles
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
		mocks   mocks
	}{
		{
			name:  "[Test] successful create medication",
			args:  args{ctx: context.Background(), request: []byte(`{"name":"med-name", "code": "code_code", "weight": 50}`)},
			want:  `{"id":0,"name":"med-name","weight":50,"code":"code_code","image_path":""}`,
			mocks: mocks{medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] create medication should return error if name is not provided",
			args:    args{ctx: context.Background(), request: []byte(`{"code": "code_code", "weight": 50}`)},
			wantErr: "name is required",
		},
		{
			name:    "[Test] create medication should return error if name is not valid",
			args:    args{ctx: context.Background(), request: []byte(`{"name": "med#name", "code": "code_code", "weight": 50}`)},
			wantErr: "invalid name",
		},
		{
			name:    "[Test] create medication should return error if code is not provided",
			args:    args{ctx: context.Background(), request: []byte(`{"name": "med-name", "weight": 50}`)},
			wantErr: "code is required",
		},
		{
			name:    "[Test] create medication should return error if code is not valid",
			args:    args{ctx: context.Background(), request: []byte(`{"name": "med-name", "code": "code-code", "weight": 50}`)},
			wantErr: "invalid code",
		},
		{
			name:    "[Test] create medication should return error if weight is not provided",
			args:    args{ctx: context.Background(), request: []byte(`{"name": "med-name", "code": "code_code"}`)},
			wantErr: "weight is required",
		},
		{
			name:  "[Test] create medication with image",
			args:  args{ctx: context.Background(), request: []byte(`{"name": "med-name", "code": "code_code", "weight":100 }`), file: mockedFile.NewMockedFile(), filename: ""},
			want:  fmt.Sprintf(`{"id":0,"name":"med-name","weight":100,"code":"code_code","image_path":"%s-"}`, time.Now().Format(time.RFC3339)),
			mocks: mocks{medicationRepository: mock.NewMockedMedication(), iofile: mockedFile.NewMockedImageIOFile()},
		},
		{
			name:    "[Test] create medication with image should return error if file type is not valid",
			args:    args{ctx: context.Background(), request: []byte(`{"name": "med-name", "code": "code_code", "weight":100 }`), file: mockedFile.NewMockedFile(), filename: ""},
			wantErr: "invalid file type",
			mocks:   mocks{medicationRepository: mock.NewMockedMedication(), iofile: mockedFile.NewMockedPDFIOFile()},
		},
		{
			name:    "[Test] create medication with image should return error if error occured while saving file",
			args:    args{ctx: context.Background(), request: []byte(`{"name": "med-name", "code": "code_code", "weight":100 }`), file: mockedFile.NewMockedFile(), filename: ""},
			wantErr: "error saving file",
			mocks:   mocks{medicationRepository: mock.NewMockedMedication(), iofile: mockedFile.NewMockedImageSavingFailure()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			medicationUsecase := NewMedicationUsecase(tt.mocks.medicationRepository, tt.mocks.iofile)
			got, err := medicationUsecase.Create(tt.args.ctx, tt.args.request, tt.args.file, tt.args.filename)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestGetMedication(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint
	}
	type mocks struct {
		medicationRepository repository.MedicationRepo
		iofile               iofile.IOFiles
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
		mocks   mocks
	}{
		{
			name:  "[Test] successful medication retrieval",
			args:  args{ctx: context.Background(), id: 1},
			want:  `{"id":0,"name":"Medication","weight":50,"code":"code","image_path":"image_path"}`,
			mocks: mocks{medicationRepository: mock.NewMockedMedication()},
		},
		{
			name:    "[Test] medication retrieval should return error if medication does not exists",
			args:    args{ctx: context.Background(), id: 2},
			wantErr: "record not found",
			mocks:   mocks{medicationRepository: mock.NewFailureMockedMedication()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			medicationUsecase := NewMedicationUsecase(tt.mocks.medicationRepository, tt.mocks.iofile)
			got, err := medicationUsecase.Get(tt.args.ctx, tt.args.id)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestGetAllMedication(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type mocks struct {
		medicationRepository repository.MedicationRepo
		iofile               iofile.IOFiles
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
		mocks   mocks
	}{
		{
			name:  "[Test] successful medication retrieval",
			args:  args{ctx: context.Background()},
			want:  `[{"id":0,"name":"Medication","weight":50,"code":"code","image_path":"image_path"}]`,
			mocks: mocks{medicationRepository: mock.NewMockedMedication()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			medicationUsecase := NewMedicationUsecase(tt.mocks.medicationRepository, tt.mocks.iofile)
			got, err := medicationUsecase.GetAll(tt.args.ctx)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, tt.want, string(got))
		})
	}
}

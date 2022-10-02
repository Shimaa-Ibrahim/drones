package apis

import (
	"bytes"
	"context"
	"fmt"
	"github/Shimaa-Ibrahim/drones/iofile"
	mockedFile "github/Shimaa-Ibrahim/drones/iofile/mock"
	"github/Shimaa-Ibrahim/drones/repository"
	"github/Shimaa-Ibrahim/drones/repository/mock"
	"io"
	"path"
	"runtime"

	"github/Shimaa-Ibrahim/drones/usecase"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestMedicationRegister(t *testing.T) {
	type mocks struct {
		medicationRepository repository.MedicationRepo
		MockedIOFile         iofile.IOFiles
	}
	tests := []struct {
		name       string
		ctx        context.Context
		image      string
		reqBody    string
		want       string
		statusCode int
		mocks      mocks
	}{
		{
			name:       "[Test] successful register medication with image",
			ctx:        context.Background(),
			image:      "http/apis/test-files/go.png",
			reqBody:    `{"name":"go","code":"go_code","weight":1}`,
			want:       fmt.Sprintf(`{"id":0,"name":"go","weight":1,"code":"go_code","image_path":"%v-go.png"}`, time.Now().Format(time.RFC3339)),
			statusCode: http.StatusCreated,
			mocks:      mocks{medicationRepository: mock.NewMockedMedication(), MockedIOFile: iofile.NewIOFile()},
		},
		{
			name:       "[Test] successful register medication without image",
			ctx:        context.Background(),
			image:      "",
			reqBody:    `{"name":"go","code":"noImage","weight":1}`,
			want:       `{"id":0,"name":"go","weight":1,"code":"noImage","image_path":""}`,
			statusCode: http.StatusCreated,
			mocks:      mocks{medicationRepository: mock.NewMockedMedication(), MockedIOFile: iofile.NewIOFile()},
		},
		{
			name:       "[Test] register medication should return error if file is not image",
			ctx:        context.Background(),
			image:      "http/apis/test-files/txt.txt",
			reqBody:    `{"name":"go","code":"noImage","weight":1}`,
			want:       `{"error":"invalid file type"}`,
			statusCode: http.StatusBadRequest,
			mocks:      mocks{medicationRepository: mock.NewMockedMedication(), MockedIOFile: iofile.NewIOFile()},
		},
		{
			name:       "[Test] failed register medication with invalid request body",
			ctx:        context.Background(),
			image:      "",
			reqBody:    `{"name":"go","code":"go#code","weight":1}`,
			want:       `{"error":"invalid code"}`,
			statusCode: http.StatusBadRequest,
			mocks:      mocks{medicationRepository: mock.NewMockedMedication(), MockedIOFile: iofile.NewIOFile()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			medicationUsecase := usecase.NewMedicationUsecase(tt.mocks.medicationRepository, tt.mocks.MockedIOFile)
			medicationAPIs := NewMedicationAPIs(medicationUsecase)
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			writer.WriteField("medication", tt.reqBody)
			part, err := writer.CreateFormFile("image", tt.image)
			if err != nil {
				t.Errorf("[Error] createFromFile: %v\n", err)
			}
			if tt.image != "" {
				file, err := os.Open(tt.image)
				if err != nil {
					t.Errorf("[Error] Open file: %v\n", err)
				}
				defer file.Close()
				_, err = io.Copy(part, file)
				if err != nil {
					t.Errorf("[Error] Copy File: %v\n", err)
				}
			}
			writer.Close()
			req, err := http.NewRequest("POST", "", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(medicationAPIs.Register)
			handler.ServeHTTP(rr, req.WithContext(tt.ctx))
			result := rr.Body.String()
			assert.Equal(t, tt.statusCode, rr.Code)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGetMediction(t *testing.T) {
	type mocks struct {
		medicationRepository repository.MedicationRepo
		MockedIOFile         iofile.IOFiles
	}
	tests := []struct {
		name       string
		ctx        context.Context
		url        string
		want       string
		statusCode int
		mocks      mocks
	}{
		{
			name:       "[Test] successful medication retrieval",
			ctx:        context.Background(),
			url:        "/1",
			want:       `{"id":0,"name":"Medication","weight":50,"code":"code","image_path":"image_path"}`,
			statusCode: http.StatusOK,
			mocks:      mocks{medicationRepository: mock.NewMockedMedication(), MockedIOFile: mockedFile.NewMockedImageIOFile()},
		},
		{
			name:       "[Test] get medication should return error if drone id is invalid",
			ctx:        context.Background(),
			url:        "/string",
			want:       `{"error":"invalid ID"}`,
			statusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			medicationUsecase := usecase.NewMedicationUsecase(tt.mocks.medicationRepository, tt.mocks.MockedIOFile)
			medicationAPIs := NewMedicationAPIs(medicationUsecase)
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/{id}", medicationAPIs.Get).Methods("GET")
			router.ServeHTTP(rr, req.WithContext(tt.ctx))
			result := rr.Body.String()
			assert.Equal(t, rr.Code, tt.statusCode)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGetAllMedications(t *testing.T) {
	type mocks struct {
		medicationRepository repository.MedicationRepo
		MockedIOFile         iofile.IOFiles
	}
	tests := []struct {
		name       string
		ctx        context.Context
		want       string
		statusCode int
		mocks      mocks
	}{
		{
			name:       "[Test] successful available drones retrieval",
			ctx:        context.Background(),
			want:       `[{"id":0,"name":"Medication","weight":50,"code":"code","image_path":"image_path"}]`,
			statusCode: http.StatusOK,
			mocks:      mocks{medicationRepository: mock.NewMockedMedication(), MockedIOFile: mockedFile.NewMockedImageIOFile()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			medicationUsecase := usecase.NewMedicationUsecase(tt.mocks.medicationRepository, tt.mocks.MockedIOFile)
			medicationAPIs := NewMedicationAPIs(medicationUsecase)
			req, err := http.NewRequest("GET", "", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(medicationAPIs.GetAll)
			handler.ServeHTTP(rr, req.WithContext(tt.ctx))
			result := rr.Body.String()
			assert.Equal(t, rr.Code, tt.statusCode)
			assert.Equal(t, tt.want, result)
		})
	}

}

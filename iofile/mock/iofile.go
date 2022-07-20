package mock

import (
	"fmt"
	"github/Shimaa-Ibrahim/drones/iofile"
	"mime/multipart"
)

type MockedFile struct{}

func NewMockedFile() multipart.File {
	return MockedFile{}
}

type MockedIOFile struct{}

type MockedPDFIOFile struct {
	MockedIOFile
}

type MockedImageIOFile struct {
	MockedIOFile
}

type MockedImageSavingFailure struct {
	MockedImageIOFile
}

func NewMockedPDFIOFile() iofile.IOFiles {
	return MockedPDFIOFile{}
}

func NewMockedImageIOFile() iofile.IOFiles {
	return MockedImageIOFile{}
}

func NewMockedImageSavingFailure() iofile.IOFiles {
	return MockedImageSavingFailure{}
}

func (MockedIOFile) GetMIMEType(fileName string) string {
	return ""
}

func (MockedPDFIOFile) GetMIMEType(fileName string) string {
	return "application/pdf"
}

func (MockedImageIOFile) GetMIMEType(fileName string) string {
	return "image/jpeg"
}

func (MockedIOFile) Save(file multipart.File, filename string) error {
	return nil
}

func (MockedImageSavingFailure) Save(file multipart.File, filename string) error {
	return fmt.Errorf("error saving file")
}

func (MockedIOFile) Write(data []string, filename string) error {
	return nil
}

func (MockedFile) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (MockedFile) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, nil
}

func (MockedFile) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (MockedFile) Close() error {
	return nil
}

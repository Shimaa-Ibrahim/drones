package iofile

import (
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"os"
)

type IOFiles interface {
	GetMIMEType(fileExt string) string
	Save(file multipart.File, filename string) error
	Write(data []string, filename string) error
}

type IOFile struct{}

func NewIOFile() IOFiles {
	return IOFile{}
}

func (i IOFile) GetMIMEType(fileExt string) string {
	return mime.TypeByExtension("." + fileExt)
}

// func (i IOFile) GetSiz

func (IOFile) Save(file multipart.File, filename string) error {
	dst, err := os.Create(filename)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return err
	}
	return err
}

func (IOFile) Write(data []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("[Error]: %v\n", err)
		return err
	}
	defer file.Close()
	for _, line := range data {
		_, err = fmt.Fprintln(file, line)
		if err != nil {
			log.Printf("[Error]: %v\n", err)
			return err
		}
	}
	return nil
}

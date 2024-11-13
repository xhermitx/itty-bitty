package utils

import (
	"errors"
	"path/filepath"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrEmptyDB       = errors.New("empty database")
	ErrInvalidURL    = errors.New("invalid URL")
)

func GetTemplatePath(fileName string) string {
	//cwd, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}
	return filepath.Join("/app/internal/", fileName)
}

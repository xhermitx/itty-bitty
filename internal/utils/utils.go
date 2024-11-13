package utils

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrEmptyDB       = errors.New("empty database")
	ErrInvalidURL    = errors.New("invalid URL")
)

func GetTemplatePath(fileName string) string {
	return filepath.Join(os.Getenv("VIEWS_PATH"), fileName)
}

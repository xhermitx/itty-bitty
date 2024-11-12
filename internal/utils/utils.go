package utils

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrEmptyDB       = errors.New("empty database")
	ErrInvalidURL    = errors.New("invalid URL")
)

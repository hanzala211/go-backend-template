package store

import "errors"

var (
	ErrNotFound = errors.New("Not Found")
	ErrConflict = errors.New("Already Exists")
)

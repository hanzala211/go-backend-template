package service

import "fmt"

type AppError struct {
	Message string
	Err     error
}

func (a *AppError) Error() string {
	return fmt.Sprintf("%s: %v", a.Message, a.Err)
}

func (a *AppError) Unwrap() error {
	return a.Err
}

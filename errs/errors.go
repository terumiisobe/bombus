package errs

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (a *AppError) AsMessage() *AppError {
	return &AppError{
		Message: a.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewDatabaseError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprintf("Database error: %s", message),
	}
}

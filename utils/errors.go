package utils

import (
	"fmt"
	"log"
)

type AppError struct {
	Layer   string
	Message string
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Layer, e.Message)
}

func CreateDatabaseErrorAndLog(message string) *AppError {
	var error = AppError{Layer: "database", Message: message}
	log.Println(error.Error())
	return &error
}

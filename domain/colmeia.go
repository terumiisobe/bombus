package domain

import (
	"time"

	"github.com/terumiisobe/bombus/errs"
)

type Colmeia struct {
	ID           int
	ColmeiaID    *int    // Additional visual ID
	QRCode       *string // Can be NULL
	Species      Species
	StartingDate time.Time
	Status       Status
}

type ColmeiaRepository interface {
	FindAll(string, string) ([]Colmeia, *errs.AppError)
	ById(string) (*Colmeia, *errs.AppError)
	Create(Colmeia) *errs.AppError
}

package repository

import (
	"bombus/domain"
	"bombus/errs"
)

type SpeciesRepository interface {
	FindAll() ([]domain.Species, *errs.AppError)
	ById(string) (*domain.Species, *errs.AppError)
}

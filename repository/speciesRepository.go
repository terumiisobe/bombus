package repository

import (
	"bombus/domain"
	"bombus/errs"
)

type ColmeiaRepository interface {
	FindAll(string, string) ([]domain.Colmeia, *errs.AppError)
	ById(string) (*domain.Colmeia, *errs.AppError)
	Create(domain.Colmeia) *errs.AppError
	Count(*domain.Species, *domain.Status) (int, *errs.AppError)
	CountBySpecies() (map[string]int, *errs.AppError)
	CountBySpeciesAndStatus() (map[string]map[string]int, *errs.AppError)
}

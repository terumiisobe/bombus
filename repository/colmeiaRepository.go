package repository

import (
	"bombus/domain"
	"bombus/errs"
)

type ColmeiaRepository interface {
	FindAll(string, string) ([]domain.Colmeia, *errs.AppError)
	ById(string) (domain.Colmeia, *errs.AppError)
	Create(domain.Colmeia) (uint64, *errs.AppError)
	CountBySpecies() (map[string]int, *errs.AppError)
	CountBySpeciesAndStatus() (map[string]map[string]int, *errs.AppError)
}

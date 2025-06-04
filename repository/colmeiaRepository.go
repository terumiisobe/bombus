package repository

import (
	"bombus/domain"
	"bombus/errs"
)

type ColmeiaRepository interface {
	FindAll(string, string) ([]domain.Colmeia, *errs.AppError)
	ById(string) (*domain.Colmeia, *errs.AppError)
	Create(domain.Colmeia) *errs.AppError
	CountGroupedBySpecies() map[int]int
}

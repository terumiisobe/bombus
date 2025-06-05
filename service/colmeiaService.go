package service

import (
	"bombus/domain"
	"bombus/errs"
)

type ColmeiaService interface {
	GetAllColmeia(string, string) ([]domain.Colmeia, *errs.AppError)
	GetColmeia(string) (*domain.Colmeia, *errs.AppError)
	CreateColmeia(domain.Colmeia) *errs.AppError
	CreateBatchColmeia(int, domain.Colmeia) *errs.AppError
	CountBySpecies() (map[string]int, *errs.AppError)
	CountBySpeciesAndStatus() (map[string]map[string]int, *errs.AppError)
}

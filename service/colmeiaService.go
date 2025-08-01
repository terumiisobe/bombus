package service

import (
	"bombus/domain"
	"bombus/errs"
)

type ColmeiaService interface {
	//TODO: refactor name to GetColmeias
	GetAllColmeia(string, string) ([]domain.Colmeia, *errs.AppError)
	//TODO: refactor name to GetColmeiaById
	GetColmeia(string) (*domain.Colmeia, *errs.AppError)
	CreateColmeia(domain.Colmeia) *errs.AppError
	//CreateBatchColmeia(int, dto.Colmeia) *errs.AppError

	CountBySpecies() (map[domain.Species]int, *errs.AppError)
	CountByStatus() (map[domain.Status]int, *errs.AppError)
	CountBySpeciesAndStatus() (map[domain.Species]map[domain.Status]int, *errs.AppError)
}

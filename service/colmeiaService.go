package service

import (
	"bombus/errs"
	"bombus/service/dto"
)

type ColmeiaService interface {
	GetAllColmeia(string, string) ([]dto.Colmeia, *errs.AppError)
	GetColmeia(string) (*dto.Colmeia, *errs.AppError)
	CreateColmeia(dto.Colmeia) (dto.Colmeia, *errs.AppError)
	CreateBatchColmeia(int, dto.Colmeia) *errs.AppError

	CountBySpecies() (map[string]int, *errs.AppError)
	CountBySpeciesAndStatus() (map[string]map[string]int, *errs.AppError)
}

package service

import (
	"bombus/domain"
	"bombus/errs"
	"bombus/repository"
)

type ColmeiaService interface {
	GetAllColmeia(string, string) ([]domain.Colmeia, *errs.AppError)
	GetColmeia(string) (*domain.Colmeia, *errs.AppError)
	CreateColmeia(domain.Colmeia) *errs.AppError
	CountBySpecies() (map[string]int, *errs.AppError)
}

func NewColmeiaService(repository repository.ColmeiaRepository) ColmeiaServiceImplDefault {
	return ColmeiaServiceImplDefault{repository}
}

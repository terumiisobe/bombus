package service

import (
	"bombus/domain"
	"bombus/errs"
)

type ColmeiaService interface {
	GetAllColmeia(string, string) ([]domain.Colmeia, *errs.AppError)
	GetColmeia(string) (*domain.Colmeia, *errs.AppError)
	CreateColmeia(domain.Colmeia) *errs.AppError
}

type DefaultColmeiaService struct {
	repo domain.ColmeiaRepository
}

func NewColmeiaService(repository domain.ColmeiaRepository) DefaultColmeiaService {
	return DefaultColmeiaService{repository}
}

func (s DefaultColmeiaService) GetAllColmeia(status string, species string) ([]domain.Colmeia, *errs.AppError) {
	return s.repo.FindAll(status, species)
}

func (s DefaultColmeiaService) GetColmeia(id string) (*domain.Colmeia, *errs.AppError) {
	return s.repo.ById(id)
}

func (s DefaultColmeiaService) CreateColmeia(colmeia domain.Colmeia) *errs.AppError {
	return s.repo.Create(colmeia)
}

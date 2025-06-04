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
	CreateBatchColmeia(int, domain.Colmeia) *errs.AppError
	CountPerSpecies() map[int]int
}

type DefaultColmeiaService struct {
	repo repository.ColmeiaRepository
}

func NewColmeiaService(repository repository.ColmeiaRepository) DefaultColmeiaService {
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

// TODO: implement batch creation (id creation)
func (s DefaultColmeiaService) CreateBatchColmeia(quantity int, colmeia domain.Colmeia) *errs.AppError {
	return s.repo.Create(colmeia)
}

func (s DefaultColmeiaService) CountPerSpecies() map[int]int {
	return s.repo.CountGroupedBySpecies()
}

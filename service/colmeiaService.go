package service

import (
	"github.com/terumiisobe/bombus/domain"
	"github.com/terumiisobe/bombus/errs"
)

type ColmeiaService interface {
	GetAllColmeia(string, string) ([]domain.Colmeia, error)
	GetColmeia(string) (*domain.Colmeia, *errs.AppError)
}

type DefaultColmeiaService struct {
	repo domain.ColmeiaRepository
}

func (s DefaultColmeiaService) GetAllColmeia(status string, species string) ([]domain.Colmeia, error) {
	return s.repo.FindAll(status, species)
}

func (s DefaultColmeiaService) GetColmeia(id string) (*domain.Colmeia, *errs.AppError) {
	return s.repo.ById(id)
}

func NewColmeiaService(repository domain.ColmeiaRepository) DefaultColmeiaService {
	return DefaultColmeiaService{repository}
}

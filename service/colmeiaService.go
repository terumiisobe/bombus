package service

import (
	"github.com/terumiisobe/bombus/domain"
)

type ColmeiaService interface {
	GetAllColmeia() ([]domain.Colmeia, error)
}

type DefaultColmeiaService struct {
	repo domain.ColmeiaRepository
}

func (s DefaultColmeiaService) GetAllColmeia() ([]domain.Colmeia, error) {
	return s.repo.FindAll()
}

func NewColmeiaService(repository domain.ColmeiaRepository) DefaultColmeiaService {
	return DefaultColmeiaService{repository}
}

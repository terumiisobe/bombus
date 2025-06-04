package service

import (
	"bombus/domain"
	"bombus/errs"
	"bombus/repository"
)

type ColmeiaServiceImplDefault struct {
	repo repository.ColmeiaRepository
}

func (s ColmeiaServiceImplDefault) GetAllColmeia(status string, species string) ([]domain.Colmeia, *errs.AppError) {
	return s.repo.FindAll(status, species)
}

func (s ColmeiaServiceImplDefault) GetColmeia(id string) (*domain.Colmeia, *errs.AppError) {
	return s.repo.ById(id)
}

func (s ColmeiaServiceImplDefault) CreateColmeia(colmeia domain.Colmeia) *errs.AppError {
	return s.repo.Create(colmeia)
}

func (s ColmeiaServiceImplDefault) CountBySpecies() (map[string]int, *errs.AppError) {
	colmeias, err := s.repo.FindAll("", "")
	if err != nil {
		return nil, err
	}

	countBySpecies := make(map[string]int)
	for _, colmeia := range colmeias {
		countBySpecies[colmeia.Species.String()]++
	}
	return countBySpecies, nil
}

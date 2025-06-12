package service

import (
	"bombus/domain"
	"bombus/errs"
	"bombus/repository"
)

func NewColmeiaServiceImplDefault(repo repository.ColmeiaRepository) ColmeiaServiceImplDefault {
	return ColmeiaServiceImplDefault{repo}
}

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
func (s ColmeiaServiceImplDefault) CountBySpeciesAndStatus() (map[string]map[string]int, *errs.AppError) {
	colmeias, err := s.repo.FindAll("", "")
	if err != nil {
		return nil, err
	}

	countBySpeciesAndStatus := make(map[string]map[string]int)
	for _, colmeia := range colmeias {
		if _, ok := countBySpeciesAndStatus[colmeia.Species.String()]; !ok {
			countBySpeciesAndStatus[colmeia.Species.String()] = make(map[string]int)
		}
		countBySpeciesAndStatus[colmeia.Species.String()][colmeia.Status.String()]++
	}
	return countBySpeciesAndStatus, nil
}

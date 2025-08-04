package service

import (
	"bombus/domain"
	"bombus/errs"
	"bombus/repository"
)

type ColmeiaServiceImplDefault struct {
	colmeiaRepo repository.ColmeiaRepository
	speciesRepo repository.SpeciesRepository
}

func NewColmeiaServiceImplDefault(colmeiaRepo repository.ColmeiaRepository, speciesRepo repository.SpeciesRepository) ColmeiaServiceImplDefault {
	return ColmeiaServiceImplDefault{colmeiaRepo, speciesRepo}
}

func (s ColmeiaServiceImplDefault) GetAllColmeia(status string, species string) ([]domain.Colmeia, *errs.AppError) {
	return s.colmeiaRepo.FindAll(status, species)
}

func (s ColmeiaServiceImplDefault) GetColmeia(id string) (*domain.Colmeia, *errs.AppError) {
	return s.colmeiaRepo.ById(id)
}

func (s ColmeiaServiceImplDefault) CreateColmeia(colmeia domain.Colmeia) *errs.AppError {
	return s.colmeiaRepo.Create(colmeia)
}

func (s ColmeiaServiceImplDefault) CountBySpecies() (map[domain.Species]int, *errs.AppError) {
	allSpecies, err := s.speciesRepo.FindAll()
	if err != nil {
		return nil, err
	}

	countBySpecies := make(map[domain.Species]int)
	for _, species := range allSpecies {
		count, err := s.colmeiaRepo.Count(&species, nil)
		if err != nil {
			return nil, err
		}
		countBySpecies[species] = count
	}
	return countBySpecies, nil
}

func (s ColmeiaServiceImplDefault) CountByStatus() (map[domain.Status]int, *errs.AppError) {
	countByStatus := make(map[domain.Status]int)
	for statusNum := 1; statusNum <= domain.StatusCount; statusNum++ {
		status := domain.Status(statusNum)
		count, err := s.colmeiaRepo.Count(nil, &status)
		if err != nil {
			return nil, err
		}
		countByStatus[status] = count
	}
	return countByStatus, nil
}

func (s ColmeiaServiceImplDefault) CountBySpeciesAndStatus() (map[domain.Species]map[domain.Status]int, *errs.AppError) {
	allSpecies, err := s.speciesRepo.FindAll()
	if err != nil {
		return nil, err
	}

	countBySpeciesAndStatus := make(map[domain.Species]map[domain.Status]int)
	for _, species := range allSpecies {
		countBySpeciesAndStatus[species] = make(map[domain.Status]int)
		for statusNum := 1; statusNum <= domain.StatusCount; statusNum++ {
			status := domain.Status(statusNum)
			count, err := s.colmeiaRepo.Count(&species, &status)
			if err != nil {
				return nil, err
			}
			countBySpeciesAndStatus[species][status] = count
		}
	}
	return countBySpeciesAndStatus, nil
}

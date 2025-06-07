package repository

import (
	"bombus/domain"
	"bombus/errs"
	"strconv"
	"time"
)

type ColmeiaRepositoryImplStub struct {
	colmeias []domain.Colmeia
}

func NewColmeiaRepositoryImplStub() *ColmeiaRepositoryImplStub {
	mockTime := time.Date(2025, time.April, 15, 10, 30, 0, 0, time.UTC)
	colmeias := []domain.Colmeia{
		{uint64(123), intPtr(123), nil, domain.TetragosniscaAngustula, mockTime, domain.Developing, mockTime, mockTime},
		{uint64(456), intPtr(456), nil, domain.PlebeiaSp, mockTime, domain.Developing, mockTime, mockTime},
		{uint64(789), intPtr(789), nil, domain.MeliponaQuadrifasciata, mockTime, domain.Developing, mockTime, mockTime},
	}

	return &ColmeiaRepositoryImplStub{colmeias}
}

func NewColmeiaRepositoryImplStubCustomData(colmeias []domain.Colmeia) *ColmeiaRepositoryImplStub {
	return &ColmeiaRepositoryImplStub{colmeias}
}

func intPtr(i int) *int {
	return &i
}

func (s *ColmeiaRepositoryImplStub) FindAll(species, status string) ([]domain.Colmeia, *errs.AppError) {
	return s.colmeias, nil
}

func (s *ColmeiaRepositoryImplStub) ById(id string) (domain.Colmeia, *errs.AppError) {
	var colmeia domain.Colmeia
	colmeiaID, _ := strconv.ParseUint(id, 10, 64)
	for _, colmeia := range s.colmeias {
		if colmeia.ID == colmeiaID {
			return colmeia, nil
		}
	}
	return colmeia, errs.NewNotFoundError("Colmeia not found")
}

func (s *ColmeiaRepositoryImplStub) Create(colmeia domain.Colmeia) (uint64, *errs.AppError) {
	s.colmeias = append(s.colmeias, colmeia)
	return colmeia.ID, nil
}

func (s *ColmeiaRepositoryImplStub) CountBySpecies() (map[string]int, *errs.AppError) {
	count := make(map[string]int)
	for _, colmeia := range s.colmeias {
		key := strconv.Itoa(int(colmeia.Species))
		count[key]++
	}
	return count, nil
}

func (s *ColmeiaRepositoryImplStub) CountBySpeciesAndStatus() (map[string]map[string]int, *errs.AppError) {
	count := make(map[string]map[string]int)
	for _, colmeia := range s.colmeias {
		species := strconv.Itoa(int(colmeia.Species))
		status := strconv.Itoa(int(colmeia.Status))

		if _, exists := count[species]; !exists {
			count[species] = make(map[string]int)
		}
		count[species][status]++
	}
	return count, nil
}

func (s *ColmeiaRepositoryImplStub) GetColmeias() []domain.Colmeia {
	return s.colmeias
}

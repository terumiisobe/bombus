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

func NewColmeiaRepositoryStub() ColmeiaRepositoryImplStub {
	mockTime := time.Date(2025, time.April, 15, 10, 30, 0, 0, time.UTC)
	colmeias := []domain.Colmeia{
		{123, intPtr(123), nil, 1, mockTime, 1},
		{456, intPtr(456), nil, 2, mockTime, 2},
		{789, intPtr(789), nil, 2, mockTime, 3},
	}

	return ColmeiaRepositoryImplStub{colmeias}
}

func intPtr(i int) *int {
	return &i
}

func (s ColmeiaRepositoryImplStub) FindAll(species, status string) ([]domain.Colmeia, *errs.AppError) {
	return s.colmeias, nil
}

func (s ColmeiaRepositoryImplStub) ById(id string) (*domain.Colmeia, *errs.AppError) {
	colmeiaID, _ := strconv.Atoi(id)
	for _, colmeia := range s.colmeias {
		if colmeia.ID == colmeiaID {
			return &colmeia, nil
		}
	}
	return nil, errs.NewNotFoundError("Colmeia not found")
}

func (s ColmeiaRepositoryImplStub) Create(colmeia domain.Colmeia) *errs.AppError {
	s.colmeias = append(s.colmeias, colmeia)
	return nil
}

func (s ColmeiaRepositoryImplStub) CountGroupedBySpecies() map[int]int {
	count := make(map[int]int)
	for _, colmeia := range s.colmeias {
		key := int(colmeia.Species)
		if _, keyExists := count[key]; keyExists {
			count[key]++
		}
		count[key] = 1
	}
	return count
}

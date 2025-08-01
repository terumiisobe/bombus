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

func NewColmeiaRepositoryImplStub() ColmeiaRepositoryImplStub {
	mockTime := time.Date(2025, time.April, 15, 10, 30, 0, 0, time.UTC)
	mockSpecies := domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí")
	colmeias := []domain.Colmeia{
		{int(123), intPtr(123), nil, mockSpecies, mockTime, domain.Status(1)},
		{int(456), intPtr(456), nil, mockSpecies, mockTime, domain.Status(2)},
		{int(789), intPtr(789), nil, mockSpecies, mockTime, domain.Status(3)},
	}

	return ColmeiaRepositoryImplStub{colmeias}
}

func NewColmeiaRepositoryImplStubCustomData(colmeias []domain.Colmeia) ColmeiaRepositoryImplStub {
	return ColmeiaRepositoryImplStub{colmeias}
}

func intPtr(i int) *int {
	return &i
}

func (s ColmeiaRepositoryImplStub) FindAll(species, status string) ([]domain.Colmeia, *errs.AppError) {
	return s.colmeias, nil
}

func (s ColmeiaRepositoryImplStub) ById(id string) (*domain.Colmeia, *errs.AppError) {
	var colmeia domain.Colmeia
	colmeiaID, _ := strconv.Atoi(id)
	for _, colmeia := range s.colmeias {
		if colmeia.ID == colmeiaID {
			return &colmeia, nil
		}
	}
	return &colmeia, errs.NewNotFoundError("Colmeia not found")
}

func (s ColmeiaRepositoryImplStub) Create(colmeia domain.Colmeia) *errs.AppError {
	s.colmeias = append(s.colmeias, colmeia)
	return nil
}

func (s ColmeiaRepositoryImplStub) Count(species *domain.Species, status *domain.Status) (int, *errs.AppError) {
	count := 0
	for _, colmeia := range s.colmeias {
		if colmeia.Species == *species && colmeia.Status == *status {
			count++
		}
	}
	return count, nil
}

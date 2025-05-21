package domain

import (
	"bombus/errs"
	"strconv"
	"time"
)

type ColmeiaRepositoryStub struct {
	colmeias []Colmeia
}

func NewColmeiaRepositoryStub() ColmeiaRepositoryStub {
	mockTime := time.Date(2025, time.April, 15, 10, 30, 0, 0, time.UTC)
	colmeias := []Colmeia{
		{123, intPtr(123), nil, 1, mockTime, 1},
		{456, intPtr(456), nil, 2, mockTime, 2},
	}

	return ColmeiaRepositoryStub{colmeias}
}

func intPtr(i int) *int {
	return &i
}

func (s ColmeiaRepositoryStub) FindAll(species, status string) ([]Colmeia, *errs.AppError) {
	return s.colmeias, nil
}

func (s ColmeiaRepositoryStub) ById(id string) (*Colmeia, *errs.AppError) {
	colmeiaID, _ := strconv.Atoi(id)
	for _, colmeia := range s.colmeias {
		if colmeia.ID == colmeiaID {
			return &colmeia, nil
		}
	}
	return nil, errs.NewNotFoundError("Colmeia not found")
}

func (s ColmeiaRepositoryStub) Create(colmeia Colmeia) *errs.AppError {
	s.colmeias = append(s.colmeias, colmeia)
	return nil
}

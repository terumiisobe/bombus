package test

import (
	"time"
	"domain"
)

type ColmeiaRepositoryStub struct {
	colmeias []Colmeia
}

func (s ColmeiaRepositoryStub) FindAll() ([]Colmeia, error) {
	return s.colmeias, nil
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

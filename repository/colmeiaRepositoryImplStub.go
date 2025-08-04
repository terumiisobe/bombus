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
		{ID: int(123), ColmeiaID: intPtr(123), QRCode: nil, Species: mockSpecies, StartingDate: mockTime, Status: domain.Status(1)},
		{ID: int(456), ColmeiaID: intPtr(456), QRCode: nil, Species: mockSpecies, StartingDate: mockTime, Status: domain.Status(2)},
		{ID: int(789), ColmeiaID: intPtr(789), QRCode: nil, Species: mockSpecies, StartingDate: mockTime, Status: domain.Status(3)},
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
		if species != nil && status != nil {
			if colmeia.Species.GetId() == species.GetId() && colmeia.Status == *status {
				count++
			}
		} else if species != nil {
			if colmeia.Species.GetId() == species.GetId() {
				count++
			}
		} else if status != nil {
			if colmeia.Status == *status {
				count++
			}
		} else {
			count++
		}
	}
	return count, nil
}

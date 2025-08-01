package repository

import (
	"bombus/domain"
	"bombus/errs"
	"strconv"
)

type SpeciesRepositoryImplStub struct {
	species []domain.Species
}

func NewSpeciesRepositoryImplStub() SpeciesRepositoryImplStub {
	species := []domain.Species{
		domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí"),
		domain.NewSpecies(2, "Plebeia Sp.", "Mirim"),
		domain.NewSpecies(3, "Melipona Quadrifasciata", "Mandaçaia"),
		domain.NewSpecies(4, "Melipona Bicolor", "Uruçu"),
		domain.NewSpecies(5, "Scaptotrigona Bipunctata", "Tubuna"),
		domain.NewSpecies(6, "Scaptotrigona Depilis", "Canudo"),
	}

	return SpeciesRepositoryImplStub{species}
}

func (s SpeciesRepositoryImplStub) FindAll() ([]domain.Species, *errs.AppError) {
	return s.species, nil
}

func (s SpeciesRepositoryImplStub) ById(id string) (*domain.Species, *errs.AppError) {
	var species domain.Species
	speciesID, _ := strconv.Atoi(id)
	for _, species := range s.species {
		if species.GetId() == speciesID {
			return &species, nil
		}
	}
	return &species, errs.NewNotFoundError("species")
}

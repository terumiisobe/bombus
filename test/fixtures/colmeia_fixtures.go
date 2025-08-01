package fixtures

import (
	"bombus/domain"
	"bombus/test/builders"
	"time"
)

// ColmeiaFixtures provides common test data scenarios
type ColmeiaFixtures struct {
	baseDate time.Time
}

func NewColmeiaFixtures() *ColmeiaFixtures {
	return &ColmeiaFixtures{
		baseDate: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
}

func (f *ColmeiaFixtures) EmptyColmeia() []domain.Colmeia {
	return []domain.Colmeia{}
}

// SingleColmeia returns a single Colmeia for basic tests
func (f *ColmeiaFixtures) SingleColmeia() []domain.Colmeia {
	return []domain.Colmeia{
		builders.NewColmeiaBuilder().
			WithID(1).
			WithSpecies(domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí")).
			WithStatus(domain.Developing).
			Build(),
	}
}

// MultipleColmeiaSameSpecies returns multiple Colmeia instances of the same species
func (f *ColmeiaFixtures) MultipleColmeiaSameSpecies() []domain.Colmeia {
	return []domain.Colmeia{
		builders.NewColmeiaBuilder().WithID(1).WithStatus(domain.HoneyReady).Build(),
		builders.NewColmeiaBuilder().WithID(2).WithStatus(domain.Developing).Build(),
		builders.NewColmeiaBuilder().WithID(3).WithStatus(domain.Unknown).Build(),
	}
}

// MultipleColmeiaDifferentSpecies returns Colmeia instances with different species
func (f *ColmeiaFixtures) MultipleColmeiaDifferentSpecies() []domain.Colmeia {
	return []domain.Colmeia{
		builders.NewColmeiaBuilder().
			WithID(1).
			WithSpecies(domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí")).
			WithStatus(domain.HoneyReady).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(2).
			WithSpecies(domain.NewSpecies(2, "Plebeia Sp", "Plebeia")).
			WithStatus(domain.Developing).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(3).
			WithSpecies(domain.NewSpecies(3, "Melipona Quadrifasciata", "Melipona")).
			WithStatus(domain.Unknown).
			Build(),
	}
}

// MultipleColmeiaDifferentSpeciesAndStatus returns Colmeia instances with different species and status
func (f *ColmeiaFixtures) MultipleColmeiaDifferentSpeciesAndStatus() []domain.Colmeia {
	return []domain.Colmeia{
		builders.NewColmeiaBuilder().
			WithID(1).
			WithSpecies(domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí")).
			WithStatus(domain.Developing).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(2).
			WithSpecies(domain.NewSpecies(2, "Plebeia Sp", "Plebeia")).
			WithStatus(domain.HoneyReady).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(3).
			WithSpecies(domain.NewSpecies(3, "Melipona Quadrifasciata", "Melipona")).
			WithStatus(domain.Unknown).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(4).
			WithSpecies(domain.NewSpecies(3, "Melipona Quadrifasciata", "Melipona")).
			WithStatus(domain.HoneyReady).
			Build(),
	}
}

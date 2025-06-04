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

// SingleColmeia returns a single Colmeia for basic tests
func (f *ColmeiaFixtures) SingleColmeia() []domain.Colmeia {
	return []domain.Colmeia{
		builders.NewColmeiaBuilder().
			WithID(1).
			WithSpecies(domain.TetragosniscaAngustula).
			WithStatus(domain.Developing).
			Build(),
	}
}

// MultipleColmeiaSameSpecies returns multiple Colmeia instances of the same species
func (f *ColmeiaFixtures) MultipleColmeiaSameSpecies() []domain.Colmeia {
	return []domain.Colmeia{
		builders.NewColmeiaBuilder().WithID(1).WithStatus(domain.HoneyReady).Build(),
		builders.NewColmeiaBuilder().WithID(2).WithStatus(domain.Developing).Build(),
		builders.NewColmeiaBuilder().WithID(3).WithStatus(domain.PetBottle).Build(),
	}
}

// MultipleColmeiaDifferentSpecies returns Colmeia instances with different species
func (f *ColmeiaFixtures) MultipleColmeiaDifferentSpecies() []domain.Colmeia {
	return []domain.Colmeia{
		builders.NewColmeiaBuilder().
			WithID(1).
			WithSpecies(domain.TetragosniscaAngustula).
			WithStatus(domain.HoneyReady).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(2).
			WithSpecies(domain.PlebeiaSp).
			WithStatus(domain.Developing).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(3).
			WithSpecies(domain.MeliponaQuadrifasciata).
			WithStatus(domain.PetBottle).
			Build(),
	}
}

// MultipleColmeiaDifferentSpeciesAndStatus returns Colmeia instances with different species and status
func (f *ColmeiaFixtures) MultipleColmeiaDifferentSpeciesAndStatus() []domain.Colmeia {
	return []domain.Colmeia{
		builders.NewColmeiaBuilder().
			WithID(1).
			WithSpecies(domain.TetragosniscaAngustula).
			WithStatus(domain.Developing).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(2).
			WithSpecies(domain.PlebeiaSp).
			WithStatus(domain.HoneyReady).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(3).
			WithSpecies(domain.MeliponaQuadrifasciata).
			WithStatus(domain.PetBottle).
			Build(),
		builders.NewColmeiaBuilder().
			WithID(4).
			WithSpecies(domain.MeliponaQuadrifasciata).
			WithStatus(domain.HoneyReady).
			Build(),
	}
}

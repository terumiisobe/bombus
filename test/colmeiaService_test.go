package test

import (
	"bombus/domain"
	"bombus/repository"
	"bombus/service"
	"bombus/test/fixtures"
	"reflect"
	"testing"
)

var (
	colmeiaService  service.ColmeiaService
	colmeiaFixtures *fixtures.ColmeiaFixtures
)

func init() {
	colmeiaFixtures = fixtures.NewColmeiaFixtures()
}

func TestColmeiaService_CountBySpecies(t *testing.T) {
	t.Run("Empty repo", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.EmptyColmeia())
		speciesRepository := repository.NewSpeciesRepositoryImplStub()
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[string]int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with one colmeia", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.SingleColmeia())
		speciesRepository := repository.NewSpeciesRepositoryImplStub()
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[domain.Species]int{
			domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí"): 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with multiple colmeia of same species", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.MultipleColmeiaSameSpecies())
		speciesRepository := repository.NewSpeciesRepositoryImplStub()
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[domain.Species]int{
			domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí"): 3,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with multiple colmeia of different species", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.MultipleColmeiaDifferentSpecies())
		speciesRepository := repository.NewSpeciesRepositoryImplStub()
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[domain.Species]int{
			domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí"):    1,
			domain.NewSpecies(2, "Plebeia Sp", "Plebeia"):               1,
			domain.NewSpecies(3, "Melipona Quadrifasciata", "Melipona"): 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
func TestColmeiaService_CountBySpeciesAndStatus(t *testing.T) {
	t.Run("Empty repo", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData([]domain.Colmeia{})
		speciesRepository := repository.NewSpeciesRepositoryImplStub()
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpeciesAndStatus()
		want := map[domain.Species]map[domain.Status]int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with one colmeia", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.SingleColmeia())
		speciesRepository := repository.NewSpeciesRepositoryImplStub()
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpeciesAndStatus()
		want := map[domain.Species]map[domain.Status]int{
			domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí"): {
				domain.Developing: 1,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with multiple colmeia of different species and status", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.MultipleColmeiaDifferentSpeciesAndStatus())
		speciesRepository := repository.NewSpeciesRepositoryImplStub()
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpeciesAndStatus()
		want := map[domain.Species]map[domain.Status]int{
			domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí"): {
				domain.Developing: 1,
			},
			domain.NewSpecies(2, "Plebeia Sp", "Plebeia"): {
				domain.HoneyReady: 1,
			},
			domain.NewSpecies(3, "Melipona Quadrifasciata", "Melipona"): {
				domain.Unknown:    1,
				domain.HoneyReady: 1,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

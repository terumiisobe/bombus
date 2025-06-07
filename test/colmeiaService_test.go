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
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[string]int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with one colmeia", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.SingleColmeia())
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[string]int{
			domain.TetragosniscaAngustula.String(): 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with multiple colmeia of same species", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.MultipleColmeiaSameSpecies())
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[string]int{
			domain.TetragosniscaAngustula.String(): 3,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with multiple colmeia of different species", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.MultipleColmeiaDifferentSpecies())
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[string]int{
			domain.TetragosniscaAngustula.String(): 1,
			domain.PlebeiaSp.String():              1,
			domain.MeliponaQuadrifasciata.String(): 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
func TestColmeiaService_CountBySpeciesAndStatus(t *testing.T) {
	t.Run("Empty repo", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData([]domain.Colmeia{})
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository)

		got, _ := colmeiaService.CountBySpeciesAndStatus()
		want := map[string]map[string]int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with one colmeia", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.SingleColmeia())
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository)

		got, _ := colmeiaService.CountBySpeciesAndStatus()
		want := map[string]map[string]int{
			domain.TetragosniscaAngustula.String(): {
				domain.Developing.String(): 1,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with multiple colmeia of different species and status", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.MultipleColmeiaDifferentSpeciesAndStatus())
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository)

		got, _ := colmeiaService.CountBySpeciesAndStatus()
		want := map[string]map[string]int{
			domain.TetragosniscaAngustula.String(): {
				domain.Developing.String(): 1,
			},
			domain.PlebeiaSp.String(): {
				domain.HoneyReady.String(): 1,
			},
			domain.MeliponaQuadrifasciata.String(): {
				domain.PetBottle.String():  1,
				domain.HoneyReady.String(): 1,
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

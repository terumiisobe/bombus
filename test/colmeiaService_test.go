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

		allSpecies, _ := speciesRepository.FindAll()
		want := make(map[domain.Species]int)
		for _, species := range allSpecies {
			want[species] = 0
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with one colmeia", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.SingleColmeia())
		speciesRepository := repository.NewSpeciesRepositoryImplStub()
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpecies()

		allSpecies, _ := speciesRepository.FindAll()
		want := make(map[domain.Species]int)
		for _, species := range allSpecies {
			want[species] = 0
			if species.GetId() == 1 {
				want[species] = 1
			}
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
		allSpecies, _ := speciesRepository.FindAll()
		want := make(map[domain.Species]int)
		for _, species := range allSpecies {
			want[species] = 0
			if species.GetId() == 1 {
				want[species] = 3
			}
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
		allSpecies, _ := speciesRepository.FindAll()
		want := make(map[domain.Species]int)
		for _, species := range allSpecies {
			want[species] = 0
			if species.GetId() == 1 || species.GetId() == 2 || species.GetId() == 3 {
				want[species] = 1
			}
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
func TestColmeiaService_CountBySpeciesAndStatus(t *testing.T) {

	speciesRepository := repository.NewSpeciesRepositoryImplStub()
	allSpecies, _ := speciesRepository.FindAll()
	emptyWant := map[domain.Species]map[domain.Status]int{}
	for _, species := range allSpecies {
		emptyWant[species] = make(map[domain.Status]int)
		for statusNum := 1; statusNum <= domain.StatusCount; statusNum++ {
			status := domain.Status(statusNum)
			emptyWant[species][status] = 0
		}
	}

	t.Run("Empty repo", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData([]domain.Colmeia{})
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpeciesAndStatus()
		want := emptyWant

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with one colmeia", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.SingleColmeia())
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpeciesAndStatus()
		want := emptyWant
		want[domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí")][domain.Developing] = 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Repo with multiple colmeia of different species and status", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData(colmeiaFixtures.MultipleColmeiaDifferentSpeciesAndStatus())
		colmeiaService = service.NewColmeiaServiceImplDefault(colmeiaRepository, speciesRepository)

		got, _ := colmeiaService.CountBySpeciesAndStatus()
		want := emptyWant
		want[domain.NewSpecies(1, "Tetragosnisca Angustula", "Jataí")][domain.Developing] = 1
		want[domain.NewSpecies(2, "Plebeia Sp.", "Mirim")][domain.HoneyReady] = 1
		want[domain.NewSpecies(3, "Melipona Quadrifasciata", "Mandaçaia")][domain.Unknown] = 1
		want[domain.NewSpecies(3, "Melipona Quadrifasciata", "Mandaçaia")][domain.HoneyReady] = 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

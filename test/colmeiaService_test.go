package test

import (
	"bombus/domain"
	"bombus/repository"
	"bombus/service"
	"reflect"
	"testing"
	"time"
)

var colmeiaService service.ColmeiaService

func TestColmeiaService_CountBySpecies(t *testing.T) {
	t.Run("Empty repo", func(t *testing.T) {
		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData([]domain.Colmeia{})
		colmeiaService = service.NewColmeiaService(colmeiaRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[string]int{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Repo with one colmeia", func(t *testing.T) {
		mockTime := time.Date(2025, time.April, 15, 10, 30, 0, 0, time.UTC)

		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData([]domain.Colmeia{
			{123, nil, nil, domain.TetragosniscaAngustula, mockTime, 1},
		})
		colmeiaService = service.NewColmeiaService(colmeiaRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[string]int{
			domain.TetragosniscaAngustula.String(): 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Repo with multiple colmeia of same species", func(t *testing.T) {
		mockTime := time.Date(2025, time.April, 15, 10, 30, 0, 0, time.UTC)

		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData([]domain.Colmeia{
			{123, nil, nil, domain.TetragosniscaAngustula, mockTime, 1},
			{456, nil, nil, domain.TetragosniscaAngustula, mockTime, 2},
			{789, nil, nil, domain.TetragosniscaAngustula, mockTime, 2},
			{234, nil, nil, domain.TetragosniscaAngustula, mockTime, 1},
		})
		colmeiaService = service.NewColmeiaService(colmeiaRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[string]int{
			domain.TetragosniscaAngustula.String(): 4,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Repo with multiple colmeia of different species", func(t *testing.T) {
		mockTime := time.Date(2025, time.April, 15, 10, 30, 0, 0, time.UTC)

		colmeiaRepository := repository.NewColmeiaRepositoryImplStubCustomData([]domain.Colmeia{
			{123, nil, nil, domain.TetragosniscaAngustula, mockTime, 1},
			{456, nil, nil, domain.TetragosniscaAngustula, mockTime, 2},
			{789, nil, nil, domain.PlebeiaSp, mockTime, 2},
			{234, nil, nil, domain.MeliponaQuadrifasciata, mockTime, 1},
		})
		colmeiaService = service.NewColmeiaService(colmeiaRepository)

		got, _ := colmeiaService.CountBySpecies()
		want := map[string]int{
			domain.TetragosniscaAngustula.String(): 2,
			domain.PlebeiaSp.String():              1,
			domain.MeliponaQuadrifasciata.String(): 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

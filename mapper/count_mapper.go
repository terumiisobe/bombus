package mapper

import (
	"bombus/domain"
)

func SpeciesCountToString(count map[domain.Species]int) map[string]int {
	dtos := make(map[string]int)
	for species, count := range count {
		dtos[species.GetCommonName()] = count
	}
	return dtos
}

func SpeciesAndStatusCountToString(count map[domain.Species]map[domain.Status]int) map[string]map[string]int {
	dtos := make(map[string]map[string]int)
	for species, statusCount := range count {
		dtos[species.GetCommonName()] = make(map[string]int)
		for status, count := range statusCount {
			dtos[species.GetCommonName()][status.String()] = count
		}
	}
	return dtos
}

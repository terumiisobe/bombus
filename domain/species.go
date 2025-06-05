package domain

import (
	"errors"
	"strconv"
)

type Species int

const (
	TetragosniscaAngustula Species = iota + 1
	PlebeiaSp
	MeliponaQuadrifasciata
	MeliponaBicolor
	ScaptotrigonaBipunctata
	ScaptotrigonaDepilis
)

const SpeciesCount = 6

func (s Species) String() string {
	switch s {
	case TetragosniscaAngustula:
		return "Tetragosnisca Angustula"
	case PlebeiaSp:
		return "Plebeia Sp"
	case MeliponaQuadrifasciata:
		return "Melipona Quadrifasciata"
	case ScaptotrigonaBipunctata:
		return "Scaptotrigona Bipunctata"
	case ScaptotrigonaDepilis:
		return "Scaptotrigona Depilis"
	case MeliponaBicolor:
		return "Melipona Bicolor"
	default:
		return "Espécie Desconhecida"
	}
}

func ValidateSpecies(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil || v <= 0 || v > SpeciesCount {
		return errors.New("not a species enum")
	}
	return nil
}

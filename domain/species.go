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

func ValidateSpecies(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil || v <= 0 || v > SpeciesCount {
		return errors.New("not a species enum")
	}
	return nil
}

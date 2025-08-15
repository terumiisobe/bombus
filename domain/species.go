package domain

import (
	"errors"
	"fmt"
	"strconv"
)

type Species int

const (
	TetragosniscaAngustula  Species = iota + 1 //jataí
	PlebeiaSp                                  //mirim genérica
	MeliponaQuadrifasciata                     //mandaçaia
	MeliponaBicolor                            //guaraipo
	ScaptotrigonaBipunctata                    //tubuna
	ScaptotrigonaDepilis                       //canudo
	PlebeiaGigantea                            //mirim-guaçu
	PlebeiaEmerina                             //mirim-emerina
	MeliponaMarginata                          //manduri
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
	case PlebeiaGigantea:
		return "Plebeia Gigantea"
	case PlebeiaEmerina:
		return "Plebeia Emerina"
	case MeliponaMarginata:
		return "Melipona Marginata"
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

func ParseSpecies(s string) (Species, error) {
	switch s {
	case TetragosniscaAngustula.String():
		return TetragosniscaAngustula, nil
	case PlebeiaSp.String():
		return PlebeiaSp, nil
	case MeliponaQuadrifasciata.String():
		return MeliponaQuadrifasciata, nil
	case MeliponaBicolor.String():
		return MeliponaBicolor, nil
	case ScaptotrigonaBipunctata.String():
		return ScaptotrigonaBipunctata, nil
	case ScaptotrigonaDepilis.String():
		return ScaptotrigonaDepilis, nil
	default:
		return Species(0), fmt.Errorf("invalid species: %s", s)
	}
}
func GetAllSpecies() []string {
	return []string{
		TetragosniscaAngustula.String(),
		PlebeiaSp.String(),
		MeliponaQuadrifasciata.String(),
		MeliponaBicolor.String(),
		ScaptotrigonaBipunctata.String(),
		ScaptotrigonaDepilis.String(),
		PlebeiaGigantea.String(),
		PlebeiaEmerina.String(),
		MeliponaMarginata.String(),
	}
}

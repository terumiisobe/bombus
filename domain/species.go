package domain

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

func IsSpeciesEnum(v int) bool {
	if v <= 0 && v > SpeciesCount {
		return false
	}
	return true
}

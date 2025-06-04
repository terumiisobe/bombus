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

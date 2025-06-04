package domain

type Status int

const (
	HoneyReady Status = iota + 1
	Induzida
	Developing
	PetBottle
	Empty
)

func (s Status) String() string {
	switch s {
	case HoneyReady:
		return "com mel"
	case Induzida:
		return "induzida"
	case Developing:
		return "em desenvolvimento"
	case PetBottle:
		return "PET"
	case Empty:
		return "vazia"
	default:
		return "desconhecido"
	}
}

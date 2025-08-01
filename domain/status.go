package domain

import (
	"errors"
	"fmt"
	"strconv"
)

type Status int

const (
	HoneyReady     Status = 1
	Induzida       Status = 2
	MelgueiraReady Status = 3
	Developing     Status = 4
	Empty          Status = 5
	Unknown        Status = 6
)

const StatusCount = 6

func (s Status) String() string {
	switch s {
	case HoneyReady:
		return "Com mel"
	case Induzida:
		return "Induzida"
	case MelgueiraReady:
		return "Pronto para melgueira"
	case Developing:
		return "Em desenvolvimento"
	case Empty:
		return "Vazia"
	default:
		return "Desconhecido"
	}
}

func ValidateStatus(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil || v <= 0 || v > StatusCount {
		return errors.New("not a status enum")
	}
	return nil
}

func ParseStatus(s string) (Status, error) {
	switch s {
	case Developing.String():
		return Developing, nil
	case HoneyReady.String():
		return HoneyReady, nil
	case MelgueiraReady.String():
		return MelgueiraReady, nil
	case Developing.String():
		return Developing, nil
	case Empty.String():
		return Empty, nil
	default:
		return Unknown, fmt.Errorf("invalid status: %s", s)
	}
}
func GetAllStatus() []string {
	return []string{
		HoneyReady.String(),
		Induzida.String(),
		MelgueiraReady.String(),
		Developing.String(),
		Empty.String(),
		Unknown.String(),
	}
}

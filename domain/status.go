package domain

import (
	"errors"
	"fmt"
	"strconv"
)

type Status int

const (
	HoneyReady Status = iota + 1
	Induzida
	Developing
	MelgueiraReady
	PetBottle
	Empty
)

const StatusCount = 5

func (s Status) String() string {
	switch s {
	case HoneyReady:
		return "com mel"
	case Induzida:
		return "induzida"
	case MelgueiraReady:
		return "pronto para melgueira"
	case Developing:
		return "em desenvolvimento"
	case Empty:
		return "vazia"
	default:
		return "desconhecido"
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
	case PetBottle.String():
		return PetBottle, nil
	default:
		return Status(0), fmt.Errorf("invalid status: %s", s)
	}
}
func GetAllStatus() []string {
	return []string{
		HoneyReady.String(),
		Induzida.String(),
		MelgueiraReady.String(),
		Developing.String(),
		Empty.String(),
	}
}

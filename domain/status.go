package domain

import (
	"errors"
	"strconv"
)

type Status int

const (
	HoneyReady Status = iota + 1
	Induzida
	Developing
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

func ValidateStatus(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil || v <= 0 || v > StatusCount {
		return errors.New("not a status enum")
	}
	return nil
}

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

func ValidateStatus(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil || v <= 0 || v > StatusCount {
		return errors.New("not a status enum")
	}
	return nil
}

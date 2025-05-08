package domain

type Status int

const (
	HoneyReady Status = iota + 1
	Induzida
	Developing
	PetBottle
	Empty
)

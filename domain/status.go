package domain

type Status int

const (
	HoneyReady Status = iota + 1
	Ready
	Induzida
	Developing
	PetBottle
	Empty
)

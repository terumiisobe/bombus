package domain

type Status int

const (
	HoneyReady Status = iota + 1
	Induzida
	Developing
	PetBottle
	Empty
)

const StatusCount = 5

func IsStatusEnum(v int) bool {
	if v <= 0 && v > StatusCount {
		return false
	}
	return true
}

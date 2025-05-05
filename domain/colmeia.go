package domain

import "time"

type Colmeia struct {
	ID           int
	ColmeiaID    int     // Additional visual ID
	QRCode       *string // Can be NULL
	Species      Species
	StartingDate time.Time
	Status       Status
}

type ColmeiaRepository interface {
	FindAll() ([]Colmeia, error)
}

package domain

import (
	"time"
)

type Colmeia struct {
	ID           uint64
	ColmeiaID    *int    // Additional visual ID
	QRCode       *string // Can be NULL
	Species      Species
	StartingDate time.Time
	Status       Status
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

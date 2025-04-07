package services

import (
	"github.com/terumiisobe/bombus/api/models"
	"github.com/terumiisobe/bombus/api/repository"
	"strconv"
	"time"
)

type Colmeia struct {
	ID                 int
	ColmeiaID          string  // Additional visual ID
	QRCode             *string // Can be NULL
	Species            models.Species
	StartingDate       time.Time
	Status             models.Status
	RequiresInspection bool
	RequiresMelgueira  bool
}

func FetchColmeias() []models.Colmeia {
	colmeias, err := repository.GetColmeias()
	if err != null {
		
	}
	return colmeias
}

func GetColmeia(id int) string {
	return "this is a colmeia " + strconv.Itoa(id)
}

func CreateColmeia(colmeia Colmeia) string {
	// Save user to DB (Placeholder)
	return "new colmeia created"
}

func DeleteColmeia(id int) string {
	return "colmeia deleted " + strconv.Itoa(id)
}

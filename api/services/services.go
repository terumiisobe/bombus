package services

import (
	"log"
	"strconv"
	"time"

	"github.com/terumiisobe/bombus/api/models"
	"github.com/terumiisobe/bombus/api/repository"
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
	if err != nil {
		log.Println("Database error: " + err.Error())
	}
	return colmeias
}

func GetColmeia(id int) models.Colmeia {
	colmeia, err := repository.GetColmeia(id)
	if err != nil {
		log.Println("Database error: " + err.Error())
	}
	return colmeia
}

func CreateColmeia(colmeia Colmeia) string {
	// Save user to DB (Placeholder)
	return "new colmeia created"
}

func DeleteColmeia(id int) string {
	return "colmeia deleted " + strconv.Itoa(id)
}

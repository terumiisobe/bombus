package services

import (
	"errors"
	"log"
	"time"

	"github.com/terumiisobe/bombus/api/models"
	"github.com/terumiisobe/bombus/api/repository"
)

type ColmeiaService struct {
	ID                 *int           `json:"id"`
	ColmeiaID          *string        `json:"colmeia_id"` // Additional visual ID
	QRCode             *string        `json:"qrcode"`
	Species            models.Species `json:"species" validate:"required"`
	StartingDate       *time.Time     `json:"starting_date"`
	Status             *models.Status `json:"status"`
	RequiresInspection *bool          `json:"requires_inspection"`
	RequiresMelgueira  *bool          `json:"requires_melgueira"`
}

func (s *ColmeiaService) toModel() *models.ColmeiaModel {
	return &models.ColmeiaModel{
		ID:                 *s.ID,
		ColmeiaID:          *s.ColmeiaID,
		QRCode:             s.QRCode,
		Species:            s.Species,
		StartingDate:       *s.StartingDate,
		Status:             *s.Status,
		RequiresInspection: *s.RequiresInspection,
		RequiresMelgueira:  *s.RequiresMelgueira,
	}
}

func FetchColmeias() []models.ColmeiaModel {
	colmeias, err := repository.GetColmeias()
	if err != nil {
		log.Println("Database error: " + err.Error())
		return nil
	}
	return colmeias
}

func GetColmeia(id int) (*models.ColmeiaModel, error) {
	colmeia, err := repository.GetColmeia(id)
	if err != nil {
		log.Println("Database error: " + err.Error())
		return nil, err
	}
	return &colmeia, nil
}

func CreateColmeia(colmeiaService ColmeiaService) error {
	err := repository.CreateColmeia(*colmeiaService.toModel())
	return err
}

func DeleteColmeia(id int) error {
	rowsAffected, err := repository.DeleteColmeia(id)
	if err != nil {
		log.Println("Database error: " + err.Error())
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Not found")
	}
	return err
}

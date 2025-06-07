package builders

import (
	"bombus/domain"
	"bombus/service/dto"
	"time"
)

// ColmeiaBuilder provides a fluent builder for creating Colmeia test instances
type ColmeiaBuilder struct {
	colmeia domain.Colmeia
}

// NewColmeiaBuilder creates a new ColmeiaBuilder with default values
func NewColmeiaBuilder() *ColmeiaBuilder {
	return &ColmeiaBuilder{
		colmeia: domain.Colmeia{
			ID:           1,
			Species:      domain.TetragosniscaAngustula,
			StartingDate: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
			Status:       domain.Developing,
		},
	}
}

// WithID sets the ID
func (b *ColmeiaBuilder) WithID(id uint64) *ColmeiaBuilder {
	b.colmeia.ID = id
	return b
}

// WithColmeiaID sets the visual ID
func (b *ColmeiaBuilder) WithColmeiaID(colmeiaID int) *ColmeiaBuilder {
	b.colmeia.ColmeiaID = &colmeiaID
	return b
}

// WithQRCode sets the QR code
func (b *ColmeiaBuilder) WithQRCode(qrCode string) *ColmeiaBuilder {
	b.colmeia.QRCode = &qrCode
	return b
}

// WithSpecies sets the species
func (b *ColmeiaBuilder) WithSpecies(species domain.Species) *ColmeiaBuilder {
	b.colmeia.Species = species
	return b
}

// WithStartingDate sets the starting date
func (b *ColmeiaBuilder) WithStartingDate(date time.Time) *ColmeiaBuilder {
	b.colmeia.StartingDate = date
	return b
}

// WithStatus sets the status
func (b *ColmeiaBuilder) WithStatus(status domain.Status) *ColmeiaBuilder {
	b.colmeia.Status = status
	return b
}

// Build returns the constructed Colmeia
func (b *ColmeiaBuilder) Build() domain.Colmeia {
	return b.colmeia
}

func (b *ColmeiaBuilder) BuildDTO() dto.Colmeia {
	species := b.colmeia.Species.String()
	status := b.colmeia.Status.String()
	startingDate := b.colmeia.StartingDate.Format(time.DateOnly)

	return dto.Colmeia{
		ID:           &b.colmeia.ID,
		ColmeiaID:    b.colmeia.ColmeiaID,
		QRCode:       b.colmeia.QRCode,
		Species:      &species,
		StartingDate: &startingDate,
		Status:       &status,
	}
}

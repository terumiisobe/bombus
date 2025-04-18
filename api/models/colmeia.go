package models

import (
	"time"

	"gorm.io/gorm"
)

type Species int

const (
	TetragosniscaAngustula Species = iota + 1
	PlebeiaSp
	MeliponaQuadrifasciata
	MeliponaBicolor
	ScaptotrigonaBipunctata
	ScaptotrigonaDepilis
)

type Status int

const (
	HoneyReady Status = iota + 1
	Ready
	Induzida
	Developing
	PetBottle
	Empty
)

type ColmeiaModel struct {
	ID                 int       `gorm:"primaryKey"`
	ColmeiaID          string    `gorm:"type:varchar(100);not null"` // Additional visual ID
	QRCode             *string   `gorm:"type:varchar(255)"`          // Can be NULL
	Species            Species   `gorm:"type:int;not null"`
	StartingDate       time.Time `gorm:"not null"`
	Status             Status    `gorm:"type:int;not null"`
	RequiresInspection bool      `gorm:"default:false"`
	RequiresMelgueira  bool      `gorm:"default:false"`

	CreatedAt time.Time      // Auto-managed by GORM
	UpdatedAt time.Time      // Auto-managed by GORM
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete support
}

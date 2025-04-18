package models

import (
	"gorm.io/gorm"
	"time"
)

type Activity int

const (
	MelgueiraPlacing Activity = iota
	HoneyHarvest
	PropolistHarvest
	Transfer
	Split
	BoxCleaning
	Feeding
)

type Inspection struct {
	ID             int          `gorm:"primaryKey;autoIncrement"`
	ColmeiaID      int          `gorm:"not null;index"`                                    // Foreign Key to Colmeia Table
	Colmeia        ColmeiaModel `gorm:"foreignKey:ColmeiaID;constraint:OnDelete:CASCADE;"` // Relationship
	InspectionDate time.Time    `gorm:"not null"`
	PreviousStatus Status       `gorm:"type:int;not null"` // Enum stored as int
	NextStatus     Status       `gorm:"type:int;not null"` // Enum stored as int
	Activity       Activity     `gorm:"type:int;not null"` // Enum stored as int
	Comment        string       `gorm:"type:text"`         // Free text (optional)

	CreatedAt time.Time      // Auto-managed by GORM
	UpdatedAt time.Time      // Auto-managed by GORM
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete support
}

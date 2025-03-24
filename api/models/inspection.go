package models

import (
	"time"
	"gorm.io/gorm"
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
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	ColmeiaID      uint           `gorm:"not null;index"` // Foreign Key to Colmeia Table
	Colmeia        Colmeia        `gorm:"foreignKey:ColmeiaID;constraint:OnDelete:CASCADE;"` // Relationship
	InspectionDate time.Time      `gorm:"not null"`
	PreviousStatus Status         `gorm:"type:int;not null"` // Enum stored as int
	NextStatus     Status         `gorm:"type:int;not null"` // Enum stored as int
	Activity       Activity       `gorm:"type:int;not null"` // Enum stored as int
	Comment        string         `gorm:"type:text"` // Free text (optional)

	CreatedAt      time.Time      // Auto-managed by GORM
	UpdatedAt      time.Time      // Auto-managed by GORM
	DeletedAt      gorm.DeletedAt `gorm:"index"` // Soft delete support
}


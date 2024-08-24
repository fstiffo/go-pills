package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// Unit represents a unit of measurement.
type Unit string

const (
	mg Unit = "mg"
	ml Unit = "ml"
	u  Unit = "U"
	ui Unit = "UI"
)

// ActiveIngredient represents an active ingredient in a medicine or in a prescription.
type ActiveIngredient struct {
	gorm.Model
	Name            string  `gorm:"unique;not null"`
	Stock           float32 `gorm:"not null;default:0"`
	Unit            Unit    `gorm:"not null;default:'mg';check:unit in ('ml', 'mg', 'U', 'UI')"`
	LastConsumedAt  sql.NullTime
	LastStockedAt   sql.NullTime
	Medicines       []Medicine
	Prescriptions   []Prescription
	StockLogs       []StockLog
	ConsumptionLogs []ConsumptionLog
}

// Medicine represents a medicine that can be purchased.
type Medicine struct {
	gorm.Model
	Name               string `gorm:"unique;not null"`
	MAH                string
	ActiveIngredientID uint `gorm:"not null"`
	ActiveIngredient   ActiveIngredient
	Dosage             float32 `gorm:"not null;check:dosage > 0"` // Dosage in active principle measure unit for each unit of medicine
	ATC                string  `gorm:"unique;not null;check:length(atc) >= 7"`
	Package            string
	Form               string
	BoxSize            int `gorm:"not null;check:box_size > 0"`
	StockingLogs       []StockLog
}

// Prescription represents a prescription for a single active ingredient.
type Prescription struct {
	gorm.Model
	ActiveIngredientID uint `gorm:"not null"`
	ActiveIngredient   ActiveIngredient
	Dosage             float32 `gorm:"not null;check:dosage > 0"`                      // Dosage in active principle measure unit
	DosageFrequency    int     `gorm:"not null;check:dosage_frequency > 0;default: 1"` // Dosage frequency in days
	StartDate          sql.NullTime
	EndDate            sql.NullTime
	ConsumptionLogs    []ConsumptionLog
}

// ConsumptionLog represents a log of a single consumption of a prescription.
type ConsumptionLog struct {
	gorm.Model
	PrescriptionID     uint `gorm:"not null"`
	Prescription       Prescription
	ActiveIngredientID uint `gorm:"not null"`
	ActiveIngredient   ActiveIngredient
	ConsumedAt         time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Units              float32   `gorm:"not null;check:units > 0"` // Units in active principle measure unit
}

// StockLog represents a log of a single stocking of an active ingredient.
type StockLog struct {
	gorm.Model
	MedicineID         uint `gorm:"not null"`
	Medicine           Medicine
	ActiveIngredientID uint `gorm:"not null"`
	ActiveIngredient   ActiveIngredient
	StockedAt          time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Boxes              int       `gorm:"not null;check:boxes > 0"` // Boxes of medicine
	Units              float32   `gorm:"not null;check:units > 0"` // Units in active principle measure unit
}

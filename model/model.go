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
	ui Unit = "UI"
)

// ActiveIngredient represents an active ingredient in a medicine or in a prescription.
type ActiveIngredient struct {
	gorm.Model
	Name               string         `gorm:"unique;not null"`
	ATC                string         `gorm:"uniqueIndex;not null;size:7;check:length(ATC) = 7"`
	StockedUnits       int64          `gorm:"not null;default:0"` // Stocked units of active principle x 1000 (e.g., 1 mg = 1000)
	Unit               Unit           `gorm:"not null;default:'mg';check:unit in ('ml', 'mg', 'UI')"`
	LastIntakeUpdate   sql.NullTime   // Last date when StockedUnits where update considering regular intake on the base of active prescriptions
	LastStockUpdate    sql.NullTime   // Last date when StockedUnits where update for restocking
	ManualStockUpdater bool           `gorm:"not null;default:false"` // If true, the last stock update was manual
	Medicines          []Medicine     `gorm:"foreignKey:RelatedATC;references:ATC"`
	Prescriptions      []Prescription `gorm:"foreignKey:RelatedATC;references:ATC"`
	StockLogs          []StockLog     `gorm:"foreignKey:RelatedATC;references:ATC"`
}

// Medicine represents a medicine that can be purchased.
type Medicine struct {
	gorm.Model
	Name       string `gorm:"unique;not null"`
	MAH        string
	RelatedATC string `gorm:"not null"`
	AIC        string `gorm:"unique;not null;size: 9;check:length(AIC) = 9"`
	Dosage     int64  `gorm:"not null;check:dosage > 0"` // Active ingredient units prescribed x 1000 (e.g., 1 mg = 1000)
	Package    string
	Form       string
	BoxSize    int `gorm:"not null;check:box_size > 0"`
	StockLogs  []StockLog
}

// Prescription represents a prescription for a single active ingredient.
type Prescription struct {
	gorm.Model
	RelatedATC      string       `gorm:"not null"`
	Dosage          int64        `gorm:"not null;check:dosage > 0"`                      // Active ingredient units prescribed x 1000 (e.g., 1 mg = 1000)
	DosingFrequency int          `gorm:"not null;check:dosing_frequency > 0;default: 1"` // Dosing frequency in days
	StartDate       sql.NullTime // Start date of validity
	EndDate         sql.NullTime // End date of validity (A prescription is not more considered after the end date of validity)
}

// PrescriptionLog represents a log of an update to a prescription.
type PrescriptionLog struct {
	gorm.Model
	PrescriptionID  uint      `gorm:"index;not null"`
	UpdatedAt       time.Time `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
	Dosage          int64     `gorm:"not null;check:dosage > 0"`                      // Active ingredient units prescribed x 1000 (e.g., 1 mg = 1000)
	DosingFrequency int       `gorm:"not null;check:dosing_frequency > 0;default: 1"` // Dosing frequency in days
}

// StockLog represents a single stocking log of an active ingredient.
type StockLog struct {
	gorm.Model
	MedicineID uint      `gorm:"index;not null"`
	RelatedATC string    `gorm:"index;not null"`
	StockedAt  time.Time `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
	Boxes      int       `gorm:"not null;check:boxes > 0"` // Boxes of medicine
	Units      int64     `gorm:"not null;check:units > 0"` // Active ingredient units stocked x 1000 (e.g., 1 mg = 1000)
}

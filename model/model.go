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
	Name string `gorm:"unique;not null"`
	ATC  string `gorm:"uniqueIndex;not null;size:7;check:length(ATC) = 7"`
	// Stock stores units of active principle x 1000 (e.g. 1 mg = 1000)
	Stock          int64 `gorm:"not null;default:0"`
	Unit           Unit  `gorm:"not null;default:'mg';check:unit in ('ml', 'mg', 'U', 'UI')"`
	LastConsumedAt sql.NullTime
	LastStockedAt  sql.NullTime
	Medicines      []Medicine     `gorm:"foreignKey:RelatedATC;references:ATC"`
	Prescriptions  []Prescription `gorm:"foreignKey:RelatedATC;references:ATC"`
	StockLogs      []StockLog     `gorm:"foreignKey:RelatedATC;references:ATC"`
	IntakeLogs     []IntakeLog    `gorm:"foreignKey:RelatedATC;references:ATC"`
}

// Medicine represents a medicine that can be purchased.
type Medicine struct {
	gorm.Model
	Name       string `gorm:"unique;not null"`
	MAH        string
	RelatedATC string `gorm:"not null"`
	AIC        string `gorm:"unique;not null;size: 9;check:length(AIC) = 9"`
	// Dosage stores units of active ingredient for each unit of medicine x 1000 (e.g. 1 mg = 1000)
	Dosage    int64 `gorm:"not null;check:dosage > 0"`
	Package   string
	Form      string
	BoxSize   int `gorm:"not null;check:box_size > 0"`
	StockLogs []StockLog
}

// Prescription represents a prescription for a single active ingredient.
type Prescription struct {
	gorm.Model
	RelatedATC string `gorm:"not null"`
	// Dosage store active ingredient units prescibed x 1000 (e.g. 1 mg = 1000)
	Dosage          int64 `gorm:"not null;check:dosage > 0"`
	DosingFrequency int   `gorm:"not null;check:dosage_frequency > 0;default: 1"` // Dosage frequency in days
	StartDate       sql.NullTime
	EndDate         sql.NullTime
	IntakeLogs      []IntakeLog
}

// IntakeLog represents a log of a single prescription active ingredient intake.
type IntakeLog struct {
	gorm.Model
	PrescriptionID uint      `gorm:"index;not null"`
	RelatedATC     string    `gorm:"index;not null"`
	ConsumedAt     time.Time `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
	// Units stores active ingredient units comsumed x 1000 (e.g. 1 mg = 1000)
	Units int64 `gorm:"not null;check:units > 0"`
}

// StockLog represents a log of a single stocking of an active ingredient.
type StockLog struct {
	gorm.Model
	MedicineID uint      `gorm:"index;not null"`
	RelatedATC string    `gorm:"index;not null"`
	StockedAt  time.Time `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
	Boxes      int       `gorm:"not null;check:boxes > 0"` // Boxes of medicine
	// Units stores active ingredient units stocked x 1000 (e.g. 1 mg = 1000)
	Units int64 `gorm:"not null;check:units > 0"`
}

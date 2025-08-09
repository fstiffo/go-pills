package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

// GetPrescriptionsSummary returns a summary of all prescriptions
func GetPrescriptionsSummary(db *gorm.DB) pterm.TableData {
	tableData := pterm.TableData{
		{"ATC", "Active Ingredient", "Dosage", "Frequency", "Valid from", "Last consumed", "Last stocked", "Stock in days"}}
	type prescription struct {
		Prescription
		Name           string
		Unit           string
		Stock          int64
		LastConsumedAt sql.NullTime
		LastStockedAt  sql.NullTime
	}
	ps := []prescription{}
	result := db.Model(&Prescription{}).
		Select("prescriptions.*, ai.name, ai.unit, ai.stock, ai.last_consumed_at, ai.last_stocked_at").
		Joins("JOIN active_ingredients ai ON ai.atc = prescriptions.related_atc").
		Scan(&ps)
	if result.Error != nil {
		log.Fatalf("failed to get prescriptions: %v", result.Error)
	}

	for _, p := range ps {
		atc := p.RelatedATC
		name := p.Name
		unit := p.Unit
		dosage := fmt.Sprintf("%.2f %s", float64(p.Dosage)/1000, unit)
		dayOrDays := " day"
		if p.DosingFrequency > 1 {
			dayOrDays = " days"
		}
		frequency := strconv.Itoa(p.DosingFrequency) + dayOrDays
		validFrom := "-"
		if p.StartDate.Valid {
			validFrom = p.StartDate.Time.Format("2006-01-02")
		}
		lastConsumed := "-"
		if p.LastConsumedAt.Valid {
			lastConsumed = p.LastConsumedAt.Time.Format("2006-01-02")
		}
		lastStocked := "-"
		if p.LastStockedAt.Valid {
			lastStocked = p.LastStockedAt.Time.Format("2006-01-02")
		}
		stockInDays := strconv.FormatInt(p.Stock/int64(p.Dosage)*int64(p.DosingFrequency), 10)
		tableData = append(tableData, []string{atc, name, dosage, frequency, validFrom, lastConsumed, lastStocked, stockInDays})
	}

	return tableData
}

// InsertPrescription inserts a new prescription and closes any previous one without an end date.
func InsertPrescription(db *gorm.DB, relatedATC string, dosage int64, dosingFrequency int, start time.Time) error {
	p := Prescription{
		RelatedATC:      relatedATC,
		Dosage:          dosage,
		DosingFrequency: dosingFrequency,
		StartDate:       sql.NullTime{Time: start, Valid: true},
	}

	return db.Transaction(func(tx *gorm.DB) error {
		// Close previous prescription without end date for the same ATC
		var prev Prescription
		err := tx.Where("related_atc = ? AND end_date IS NULL", relatedATC).First(&prev).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if err == nil {
			prev.EndDate = sql.NullTime{Time: start, Valid: true}
			if err := tx.Save(&prev).Error; err != nil {
				return err
			}
		}

		return tx.Create(&p).Error
	})
}

package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// PrescriptionSummary contains prescription data for presentation
type PrescriptionSummary struct {
	ATC              string
	Name             string
	Unit             string
	Dosage           decimal.Decimal
	DosingFrequency  int
	StartDate        sql.NullTime
	LastIntakeUpdate sql.NullTime
	LastStockUpdate  sql.NullTime
	StockInDays      decimal.Decimal
}

// GetPrescriptionsSummary returns a summary of all prescriptions
func GetPrescriptionsSummary(db *gorm.DB) []PrescriptionSummary {
	type prescription struct {
		Prescription
		Name             string
		Unit             string
		StockedUnits     decimal.Decimal
		LastIntakeUpdate sql.NullTime
		LastStockUpdate  sql.NullTime
	}
	var ps []prescription
	result := db.Model(&Prescription{}).
		Select("prescriptions.*, ai.name, ai.unit, ai.stocked_units, ai.last_intake_update, ai.last_stock_update").
		Where("prescriptions.end_date IS NULL").
		Joins("JOIN active_ingredients ai ON ai.atc = prescriptions.related_atc").
		Scan(&ps)
	if result.Error != nil {
		log.Fatalf("failed to get prescriptions: %v", result.Error)
	}

	var summaries []PrescriptionSummary
	for _, p := range ps {
		stockInDays := decimal.Zero
		if !p.Dosage.IsZero() {
			frequency := decimal.NewFromInt(int64(p.DosingFrequency))
			stockInDays = p.StockedUnits.Mul(frequency).Div(p.Dosage)
		}
		summaries = append(summaries, PrescriptionSummary{
			ATC:              p.RelatedATC,
			Name:             p.Name,
			Unit:             p.Unit,
			Dosage:           p.Dosage,
			DosingFrequency:  p.DosingFrequency,
			StartDate:        p.StartDate,
			LastIntakeUpdate: p.LastIntakeUpdate,
			LastStockUpdate:  p.LastStockUpdate,
			StockInDays:      stockInDays,
		})
	}
	summaries = sortSummaries(summaries)

	return summaries
}

func sortSummaries(summaries []PrescriptionSummary) []PrescriptionSummary {
	stockCmp := func(a, b PrescriptionSummary) int {
		return a.StockInDays.Cmp(b.StockInDays)
	}
	slices.SortFunc(summaries, stockCmp)
	return summaries
}

// UpsertPrescription inserts or updates a prescription and updates the stock.
func UpsertPrescription(db *gorm.DB, relatedATC string, dosage decimal.Decimal, dosingFrequency int, start time.Time) error {
	if start.After(time.Now()) {
		return errors.New("start date cannot be in the future")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var ai ActiveIngredient
		if err := tx.Where("atc = ?", relatedATC).First(&ai).Error; err != nil {
			return fmt.Errorf("active ingredient with ATC %s not found", relatedATC)
		}

		if ai.LastIntakeUpdate.Valid && start.Before(ai.LastIntakeUpdate.Time) {
			return fmt.Errorf("start date (%v) must be after last intake update (%v)", start, ai.LastIntakeUpdate.Time)
		}

		// End the previous prescription if it exists without triggering noisy "record not found" logs.
		var existingPrescription Prescription
		result := tx.Where("related_atc = ? AND end_date IS NULL", relatedATC).
			Limit(1).
			Find(&existingPrescription)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected > 0 {
			existingPrescription.EndDate = sql.NullTime{Time: start, Valid: true}
			if err := tx.Save(&existingPrescription).Error; err != nil {
				return err
			}
		}

		// Create the new prescription
		newPrescription := Prescription{
			RelatedATC:      relatedATC,
			Dosage:          dosage,
			DosingFrequency: dosingFrequency,
			StartDate:       sql.NullTime{Time: start, Valid: true},
		}
		if err := tx.Create(&newPrescription).Error; err != nil {
			return err
		}

		// With the correct prescription history now in the DB, update the stock.
		if err := UpdateStockedUnitsFromIntake(tx, &ai); err != nil {
			return fmt.Errorf("failed to update stock after upsert: %w", err)
		}

		return nil
	})
}

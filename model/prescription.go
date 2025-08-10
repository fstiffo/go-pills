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
		{"ATC", "Active Ingredient", "Dosage", "Frequency", "Valid from", "Last intake update", "Last stocked", "Stock in days"}}
	type prescription struct {
		Prescription
		Name             string
		Unit             string
		StockedUnits     int64
		LastIntakeUpdate sql.NullTime
		LastStockUpdate  sql.NullTime
	}
	var ps []prescription
	result := db.Model(&Prescription{}).
		Select("prescriptions.*, ai.name, ai.unit, ai.stocked_units, ai.last_intake_update, ai.last_stock_update").
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
			validFrom = p.StartDate.Time.Format(`2006-01-02`)
		}
		lastIntake := "-"
		if p.LastIntakeUpdate.Valid {
			lastIntake = p.LastIntakeUpdate.Time.Format("2006-01-02")
		}
		lastStock := "-"
		if p.LastStockUpdate.Valid {
			lastStock = p.LastStockUpdate.Time.Format("2006-01-02")
		}
		stockInDays := strconv.FormatInt(p.StockedUnits*int64(p.DosingFrequency)/p.Dosage, 10)
		tableData = append(tableData, []string{atc, name, dosage, frequency, validFrom, lastIntake, lastStock, stockInDays})
	}

	return tableData
}

// UpsertPrescription inserts or updates a prescription and updates the stock.
func UpsertPrescription(db *gorm.DB, relatedATC string, dosage int64, dosingFrequency int, start time.Time) error {
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

		// End the previous prescription if it exists
		var existingPrescription Prescription
		err := tx.Where("related_atc = ? AND end_date IS NULL", relatedATC).First(&existingPrescription).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if err == nil {
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

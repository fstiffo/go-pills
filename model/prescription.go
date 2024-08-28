package model

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

// GetPrescriptionsSummary returns a summary of all prescriptions
func GetPrescriptionsSummary(db *gorm.DB) pterm.TableData {
	tableData := pterm.TableData{{"Active Igredient", "Dosage", "Frequency", "Valid from", "Last consumed", "Stock in days"}}
	prescriptions := []Prescription{}
	result := db.Where("end_date IS NULL OR end_date <= ?", time.Now()).Find(&prescriptions)
	if result.Error != nil {
		log.Fatalf("failed to get prescriptions: %v", result.Error)
	}

	for _, p := range prescriptions {
		var activeIngredient ActiveIngredient
		result := db.Where(&ActiveIngredient{ATC: p.RelatedATC}).First(&activeIngredient)
		if result.Error != nil {
			log.Fatalf("failed to get active ingredient: %v", result.Error)
		}
		name := activeIngredient.Name
		unit := activeIngredient.Unit
		dosage := fmt.Sprintf("%.2f %s", float64(p.Dosage)/1000, unit)
		frequency := strconv.Itoa(p.DosageFrequency)
		validFrom := p.StartDate.Time.Format("2006-01-02")
		lastConsumed := activeIngredient.LastConsumedAt.Time.Format("2006-01-02")
		stockInDays := strconv.FormatInt(activeIngredient.Stock/int64(p.DosageFrequency)*int64(p.DosageFrequency), 10)
		tableData = append(tableData, []string{name, dosage, frequency, validFrom, lastConsumed, stockInDays})
	}

	return tableData
}

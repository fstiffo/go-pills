package model

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

// GetPrescriptionsSummary returns a summary of all prescriptions
func GetPrescriptionsSummary(db *gorm.DB) pterm.TableData {
	tableData := pterm.TableData{{"Active Igredient", "Dosage", "Frequency", "Valid from", "Last consumed", "Stock in days"}}
	type prescription struct {
		Prescription
		Name           string
		Unit           string
		Stock          int64
		LastConsumedAt sql.NullTime
	}
	ps := []prescription{}
	result := db.Model(&Prescription{}).
		Select("prescriptions.*, ai.name, ai.unit, ai.stock, ai.last_consumed_at").
		Joins("JOIN active_ingredients ai ON ai.atc = prescriptions.related_atc").
		Scan(&ps)
	if result.Error != nil {
		log.Fatalf("failed to get prescriptions: %v", result.Error)
	}

	for _, p := range ps {
		name := p.Name
		unit := p.Unit
		dosage := fmt.Sprintf("%.2f %s", float64(p.Dosage)/1000, unit)
		frequency := strconv.Itoa(p.DosageFrequency)
		validFrom := p.StartDate.Time.Format("2006-01-02")
		lastConsumed := p.LastConsumedAt.Time.Format("2006-01-02")
		stockInDays := strconv.FormatInt(p.Stock/int64(p.DosageFrequency)*int64(p.DosageFrequency), 10)
		tableData = append(tableData, []string{name, dosage, frequency, validFrom, lastConsumed, stockInDays})
	}

	return tableData
}

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
	tableData := pterm.TableData{
		{"Active Igredient", "Dosage", "Frequency", "Valid from", "Last consumed", "Last stocked", "Stock in days"}}
	type prescription struct {
		Prescription
		Name           string
		Unit           string
		Stock          int64
		LastConsumedAt sql.NullTime
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
		name := p.Name
		unit := p.Unit
		dosage := fmt.Sprintf("%.2f %s", float64(p.Dosage)/1000, unit)
		dayOrDays := " day "
		if p.DosingFrequency > 1 {
			dayOrDays = " days"
		}
		frequency := strconv.Itoa(p.DosingFrequency) + dayOrDays
		validFrom := p.StartDate.Time.Format("2006-01-02")
		lastConsumed := p.LastConsumedAt.Time.Format("2006-01-02")
		stockInDays := strconv.FormatInt(p.Stock/int64(p.DosingFrequency)*int64(p.DosingFrequency), 10)
		tableData = append(tableData, []string{name, dosage, frequency, validFrom, lastConsumed, stockInDays})
	}

	return tableData
}

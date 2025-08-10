package view

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/validation"
	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

func updatePrescriptionsScreen() {
	clearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("UPDATE PRESCRIPTIONS")

	// List current prescriptions
	pterm.Println("\nCurrent prescriptions:")
	summaries := model.GetPrescriptionsSummary(control.GetDB())
	tableData := prescriptionSummaryTableData(summaries)
	_ = pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithBoxed().WithData(tableData).Render()

	// Ask for ATC code to add or update
	pterm.Println("\nEnter ATC code to add/update (blank to cancel):")
	var atc string
	_, _ = fmt.Scanln(&atc)
	if atc == "" {
		return
	}
	atc = strings.ToUpper(atc)

	if err := validation.ValidateAIC(atc); err != nil {
		pterm.Error.Println(err)
		return
	}

	// Check if the active ingredient exists
	_, err := model.GetActiveIngredientByATC(control.GetDB(), atc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pterm.Warning.Printf("Active ingredient with ATC %s not found. Let's add it.\n", atc)
			name, _ := pterm.DefaultInteractiveTextInput.Show("Enter active ingredient name")
			unitStr, _ := pterm.DefaultInteractiveTextInput.Show("Enter unit (mg, ml, UI)")
			unit := model.Unit(unitStr)

			newAI := &model.ActiveIngredient{
				Name: name,
				ATC:  atc,
				Unit: unit,
			}
			if err := model.InsertActiveIngredient(control.GetDB(), newAI); err != nil {
				pterm.Error.Println(err)
				return
			}
			pterm.Success.Printf("Active ingredient %s added.\n", name)
		} else {
			pterm.Error.Println(err)
			return
		}
	}

	// Gather prescription details
	pterm.Println("Enter dosage (units x1000):")
	var dosage int64
	_, _ = fmt.Scanln(&dosage)
	if err := validation.ValidateDosage(dosage); err != nil {
		pterm.Error.Println(err)
		return
	}

	pterm.Println("Enter dosing frequency in days:")
	var freq int
	_, _ = fmt.Scanln(&freq)

	pterm.Println("Enter start date (YYYY-MM-DD, blank for today):")
	var startStr string
	_, _ = fmt.Scanln(&startStr)
	start := time.Now()
	if startStr != "" {
		if t, err := time.Parse("2006-01-02", startStr); err == nil {
			start = t
		}
	}

	if err := model.UpsertPrescription(control.GetDB(), atc, dosage, freq, start); err != nil {
		pterm.Error.Println(err)
		return
	}

	pterm.Success.Println("Prescription saved")

	// Show updated prescription and stock projections
	pterm.Println("\nUpdated prescription:")
	summaries = model.GetPrescriptionsSummary(control.GetDB())
	tableData = prescriptionSummaryTableData(summaries)
	_ = pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithBoxed().WithData(tableData).Render()
}

func prescriptionSummaryTableData(ps []model.PrescriptionSummary) pterm.TableData {
	tableData := pterm.TableData{
		{"ATC", "Active Ingredient", "Dosage", "Frequency", "Valid from", "Last intake update", "Last stocked", "Stock in days"},
	}
	for _, p := range ps {
		dosage := fmt.Sprintf("%.2f %s", float64(p.Dosage)/1000, p.Unit)
		dayOrDays := " day"
		if p.DosingFrequency > 1 {
			dayOrDays = " days"
		}
		frequency := strconv.Itoa(p.DosingFrequency) + dayOrDays
		validFrom := "-"
		if p.StartDate.Valid {
			validFrom = p.StartDate.Time.Format("2006-01-02")
		}
		lastIntake := "-"
		if p.LastIntakeUpdate.Valid {
			lastIntake = p.LastIntakeUpdate.Time.Format("2006-01-02")
		}
		lastStock := "-"
		if p.LastStockUpdate.Valid {
			lastStock = p.LastStockUpdate.Time.Format("2006-01-02")
		}
		stockInDays := strconv.FormatInt(p.StockInDays, 10)
		tableData = append(tableData, []string{p.ATC, p.Name, dosage, frequency, validFrom, lastIntake, lastStock, stockInDays})
	}
	return tableData
}

package view

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/validation"
	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

const criticalStockInDays = 15 // Alert, if stock in days is less than criticalStockInDays

func updatePrescriptionsScreen() {
	clearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("UPDATE PRESCRIPTIONS")

	// List current prescriptions
	pterm.Println("\nCurrent prescriptions:")
	summaries := model.GetPrescriptionsSummary(control.GetDB())
	tableData := prescriptionSummaryTableData(summaries)
	_ = pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithBoxed().WithData(tableData).Render()

	// Ask for ATC code to add or update a prescription
	atcStr, _ := pterm.DefaultInteractiveTextInput.Show("Enter ATC code to add/update a prescription (blank to cancel)")
	if atcStr == "" {
		return
	}
	atc := strings.ToUpper(atcStr)

	if err := validation.ValidateATC(atc); err != nil {
		pterm.Error.Println(err)
		return
	}

	// Check if the active ingredient exists
	ai, err := model.GetActiveIngredientByATC(control.GetDB(), atc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// The active ingredient does not exist
			pterm.Warning.Printf("Active ingredient with ATC %s not found. Let's add it.\n", atc)
			if newAI, err := insertActiveIngredient(atc, ai); err != nil {
				pterm.Error.Println(err)
				return
			} else {
				ai = newAI
			}
		} else {
			pterm.Error.Println(err)
			return
		}
	}

	// Gather prescription details
	dosageStr, _ := pterm.DefaultInteractiveTextInput.Show("Enter dosage (in " + string(ai.Unit) + ")")
	dosageFloat, err := strconv.ParseFloat(dosageStr, 64)
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	dosage := int64(math.Round(dosageFloat * 1000))
	if err := validation.ValidateDosage(dosage); err != nil {
		pterm.Error.Println(err)
		return
	}

	freqStr, _ := pterm.DefaultInteractiveTextInput.Show("Enter dosing frequency (in days)")
	freq, err := strconv.Atoi(freqStr)
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	if err := validation.ValidateFrequency(freq); err != nil {
		pterm.Error.Println(err)
		return
	}

	startStr, _ := pterm.DefaultInteractiveTextInput.Show("Enter start date (YYYY-MM-DD, blank for today)")
	start := time.Now()
	if startStr != "" {
		if t, err := time.Parse("2006-01-02", startStr); err != nil {
			pterm.Error.Println(err)
			return
		} else {
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
		if p.StockInDays < criticalStockInDays {
			stockInDays += "<--" // Alert
		}
		tableData = append(tableData, []string{p.ATC, p.Name, dosage, frequency, validFrom, lastIntake, lastStock, stockInDays})
	}
	return tableData
}

func insertActiveIngredient(atc string, newAI *model.ActiveIngredient) (*model.ActiveIngredient, error) {
	name, _ := pterm.DefaultInteractiveTextInput.Show("Enter active ingredient name")
	unitStr, _ := pterm.DefaultInteractiveTextInput.Show("Enter unit (mg, ml, UI)")
	if err := validation.ValidateUnit(unitStr); err != nil {
		return nil, err
	}
	unit := model.Unit(unitStr)

	newAI = &model.ActiveIngredient{
		Name: name,
		ATC:  atc,
		Unit: unit,
	}
	if err := model.InsertActiveIngredient(control.GetDB(), newAI); err != nil {
		return nil, err
	}
	pterm.Success.Printf("Active ingredient %s added.\n", name)
	return newAI, nil
}

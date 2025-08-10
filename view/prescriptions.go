package view

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

func updatePrescriptionsScreen() {
	clearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("UPDATE PRESCRIPTIONS")

	// List current prescriptions
	pterm.Println("\nCurrent prescriptions:")
	tableData := model.GetPrescriptionsSummary(control.GetDB())
	_ = pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithBoxed().WithData(tableData).Render()

	// Ask for ATC code to add or update
	pterm.Println("\nEnter ATC code to add/update (blank to cancel):")
	var atc string
	_, _ = fmt.Scanln(&atc)
	if atc == "" {
		return
	}
	atc = strings.ToUpper(atc)

	if len(atc) != 7 {
		pterm.Error.Println("Invalid ATC code: must be 7 characters long")
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
	tableData = model.GetPrescriptionsSummary(control.GetDB())
	_ = pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithBoxed().WithData(tableData).Render()
}

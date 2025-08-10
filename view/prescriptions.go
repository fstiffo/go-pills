package view

import (
	"fmt"
	"strings"
	"time"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/pterm/pterm"
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

package view

import (
	"strings"
	"time"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/validation"
	"github.com/pterm/pterm"
)

func updatePrescriptionsScreen() {
	clearScreen()

	// List current prescriptions
	pterm.Println("\nCurrent prescriptions:")
	ShowPrescriptionsSummaryTable()

	// Ask for ATC code to add or update a prescription
	atcStr, _ := promptAndValidate("Enter ATC code to add/update a prescription (leave blank any prompt to leave)", validation.ValidateATC, true)
	if atcStr == "" {
		return
	}
	atc := strings.ToUpper(atcStr)

	ai, err := getOrPromptActiveIngredient(atc)
	if err != nil {
		pterm.Error.Printf("failed to get or create active ingredient: %v\n", err)
		return
	}

	// Gather prescription details
	dosage, _ := promptAndValidate("Enter dosage (in "+string(ai.Unit)+")", validation.ValidateDosage, true)
	if dosage == 0 {
		return
	}

	freq, _ := promptAndValidate("Enter dosing frequency (in days)", validation.ValidateFrequency, true)
	if freq == 0 {
		return
	}

	startStr, _ := pterm.DefaultInteractiveTextInput.Show("Enter start date (YYYY-MM-DD, blank for today)")
	start := time.Now()
	if startStr != "" {
		t, err := time.Parse("2006-01-02", startStr)
		if err != nil {
			pterm.Error.Println(err)
			return
		}
		start = t
	}

	if err := model.UpsertPrescription(control.GetDB(), atc, dosage, freq, start); err != nil {
		pterm.Error.Println(err)
		return
	}

	pterm.Success.Println("Prescription saved")

	// Show updated prescription and stock projections
	pterm.Println("\nUpdated prescription:")
	ShowPrescriptionsSummaryTable()
}

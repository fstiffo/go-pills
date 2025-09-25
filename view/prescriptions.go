package view

import (
	"strings"
	"time"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/validation"
	"github.com/pterm/pterm"
	"github.com/shopspring/decimal"
)

// updatePrescriptionsScreen displays the prescriptions management screen, allowing the user to add or update a prescription.
// It lists current prescriptions, prompts for an ATC code, retrieves or creates the corresponding active ingredient,
// and gathers prescription details such as dosage, frequency, and start date. The function validates user input,
// saves the prescription to the database, and displays the updated prescription summary. If any step fails or is cancelled,
// the function exits gracefully with appropriate feedback.
func updatePrescriptionsScreen() {
	clearScreen()

	// List current prescriptions
	ShowPrescriptionsSummaryTable()

	// Ask for ATC code to add or update a prescription
	atcStr, _ := promptAndValidate("Enter ATC code to add/update a prescription (leave any prompt blank to exit)", validation.ValidateATC, true)
	if atcStr == "" {
		return
	}
	atc := strings.ToUpper(atcStr)

	ai, err := getOrPromptActiveIngredient(atc)
	if err != nil {
		ShowErrorWithConfirm("failed to get or create active ingredient: %v\n", err)
		return
	}

	// Gather prescription details
	dosage, _ := promptAndValidate("Enter dosage (in "+string(ai.Unit)+")", validation.ValidateDosage, true)
	if dosage.Equal(decimal.Zero) {
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
			ShowErrorWithConfirm("Invalid date format: %v\n", err)
			return
		}
		start = t
	}

	if err := model.UpsertPrescription(control.GetDB(), atc, dosage, freq, start); err != nil {
		ShowErrorWithConfirm("Failed to save prescription: %v\n", err)
		return
	}

	pterm.Success.Println("Prescription saved")

	// Show updated prescription and stock projections
	pterm.Println("\nUpdated prescription:")
	ShowPrescriptionsSummaryTable()
}

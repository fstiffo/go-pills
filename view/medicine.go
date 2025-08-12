package view

import (
	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/validation"
	"github.com/pterm/pterm"
)

func addMedicineScreen() {
	clearScreen()

	// List current prescriptions
	pterm.Println("\nCurrent medicines:")
	ShowMedicinesSummaryTable()

	name, _ := promptAndValidate("Enter medicine name to add a new one (leave blank any prompt to leave)", validation.ValidateName, true)
	if name == "" {
		return
	}

	mah, _ := pterm.DefaultInteractiveTextInput.Show("Marketing authorisation holder")

	aic, _ := promptAndValidate("AIC code", validation.ValidateAIC, true)
	if aic == "" {
		return
	}

	atc, _ := promptAndValidate("ATC code", validation.ValidateATC, true)
	if atc == "" {
		return
	}

	ai, err := getOrPromptActiveIngredient(atc)
	if err != nil {
		pterm.Error.Printf("failed to get or create active ingredient: %v\n", err)
		return
	}

	dosage, _ := promptAndValidate("Enter dosage (in "+string(ai.Unit)+")", validation.ValidateDosage, true)
	if dosage == 0 {
		return
	}

	packageStr, _ := promptAndValidate("Package (e.g. blister)", validation.ValidateName, true)
	if packageStr == "" {
		return
	}

	form, _ := promptAndValidate("Form (e.g. tablet)", validation.ValidateName, true)
	if form == "" {
		return
	}

	boxSize, _ := promptAndValidate("Box size", validation.ValidateBoxSize, true)
	if boxSize == 0 {
		return
	}

	med := model.Medicine{
		Name:       name,
		MAH:        mah,
		RelatedATC: atc,
		AIC:        aic,
		Dosage:     dosage,
		Package:    packageStr,
		Form:       form,
		BoxSize:    boxSize,
	}

	if err := model.InsertMedicine(control.GetDB(), &med); err != nil {
		pterm.Error.Printf("failed to save medicine: %v\n", err)
		return
	}

	pterm.Success.Println(name + " medicine added")
	ShowMedicinesSummaryTable()
}

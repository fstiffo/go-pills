package view

import (
	"errors"
	"strconv"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

func addMedicineScreen() {
	clearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("ADD MEDICINE")

	db := control.GetDB()

	name, _ := pterm.DefaultInteractiveTextInput.Show("Medicine name")
	if name == "" {
		pterm.Warning.Println("Name cannot be empty")
		return
	}

	mah, _ := pterm.DefaultInteractiveTextInput.Show("Marketing authorisation holder")

	aic, _ := pterm.DefaultInteractiveTextInput.Show("AIC code")
	if aic == "" {
		pterm.Warning.Println("AIC code cannot be empty")
		return
	}

	atc, _ := pterm.DefaultInteractiveTextInput.Show("ATC code")
	if atc == "" {
		pterm.Warning.Println("ATC code cannot be empty")
		return
	}

	var ai model.ActiveIngredient
	if err := db.Where("atc = ?", atc).First(&ai).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			create, _ := pterm.DefaultInteractiveConfirm.WithDefaultValue(false).Show("ATC not found. Create new active ingredient?")
			if !create {
				pterm.Warning.Println("Cannot add medicine without active ingredient")
				return
			}
			aiName, _ := pterm.DefaultInteractiveTextInput.Show("Active ingredient name")
			unitStr, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("mg").Show("Unit (mg/ml/UI)")
			ai = model.ActiveIngredient{Name: aiName, ATC: atc, Unit: model.Unit(unitStr)}
			if err := model.InsertActiveIngredient(db, &ai); err != nil {
				pterm.Error.Printf("failed to create active ingredient: %v\n", err)
				return
			}
		} else {
			pterm.Error.Printf("failed to retrieve active ingredient: %v\n", err)
			return
		}
	}

	dosageStr, _ := pterm.DefaultInteractiveTextInput.Show("Dosage (" + string(ai.Unit) + ") per unit")
	dosageF, err := strconv.ParseFloat(dosageStr, 64)
	if err != nil || dosageF <= 0 {
		pterm.Warning.Println("Invalid dosage")
		return
	}
	dosage := int64(dosageF * 1000)

	packageStr, _ := pterm.DefaultInteractiveTextInput.Show("Package")
	form, _ := pterm.DefaultInteractiveTextInput.Show("Form")
	boxSizeStr, _ := pterm.DefaultInteractiveTextInput.Show("Box size")
	boxSize, err := strconv.Atoi(boxSizeStr)
	if err != nil || boxSize <= 0 {
		pterm.Warning.Println("Invalid box size")
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

	if err := model.InsertMedicine(db, &med); err != nil {
		pterm.Error.Printf("failed to save medicine: %v\n", err)
		return
	}

	pterm.Success.Println("Medicine added")
}

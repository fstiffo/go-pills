package view

import (
	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/validation"
	"github.com/pterm/pterm"
)

const backOption = "__BACK__"

// updatePharmacyScreen displays an interactive screen for updating the pharmacy's medicine stock.
// It retrieves the list of medicines from the database, allows the user to select a medicine,
// prompts for the number of boxes to add, and updates the stock accordingly. The function also
// handles cases where no medicines are available, validates user input, and provides options to
// continue adding stock or exit. Errors encountered during database operations are shown to the user.
func updatePharmacyScreen() {
	clearScreen()

	db := control.GetDB()

	var medicines []model.Medicine
	if err := db.Order("name").Find(&medicines).Error; err != nil {
		ShowErrorWithConfirm("failed to retrieve medicines: %v\n", err)
		return
	}
	if len(medicines) == 0 {
		pterm.Warning.Println("No medicines available")
		return
	}

	var options []string
	options = append(options, backOption)
	for _, m := range medicines {
		options = append(options, m.Name)
	}

	for {
		selected, _ := pterm.
			DefaultInteractiveSelect.
			WithOptions(options).
			WithMaxHeight(len(options)).
			Show("Choose a medicine to add new boxes to pharmacy stock (" + backOption + " to leave)")
		if selected == backOption {
			break
		}

		var med model.Medicine
		for _, m := range medicines {
			if m.Name == selected {
				med = m
				break
			}
		}
		if med.Name == "" {
			pterm.Error.Println("failed to retrieve medicine")
			return
		}

		ai, err := model.GetActiveIngredientByATC(db, med.RelatedATC)
		if err != nil {
			ShowErrorWithConfirm("failed to load active ingredient: %v\n", err)
			continue
		}

		boxes, _ := promptAndValidate("Boxes to add (leave any prompt blank to exit)", validation.ValidateBoxSize, true)
		if boxes == 0 {
			continue
		}

		reset, _ := pterm.DefaultInteractiveConfirm.
			WithDefaultValue(false).
			Show("Is this the first time you are adding stock for this active ingredient?")

		units, err := model.CreateStockLog(db, med, boxes)
		if err != nil {
			ShowErrorWithConfirm("failed to create stock log: %v\n", err)
			continue
		}
		if err := model.IncrementActiveIngredientStock(db, med.RelatedATC, units, reset); err != nil {
			ShowErrorWithConfirm("failed to update stock log: %v\n", err)
			continue
		}
		formattedUnits := formatCompactDosage(units, string(ai.Unit))
		pterm.Success.Printf("Added %d boxes of %s (%s)", boxes, med.Name, formattedUnits)

		cont, _ := pterm.DefaultInteractiveConfirm.WithDefaultValue(true).Show("\nAdd more?")
		if !cont {
			break
		}
	}

	overviewScreen()
}

package view

import (
	"strconv"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/pterm/pterm"
)

const backOption = "__BACK__"

func updatePharmacyScreen() {
	clearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("UPDATE PHARMACY")

	db := control.GetDB()

	var medicines []model.Medicine
	if err := db.Find(&medicines).Error; err != nil {
		pterm.Error.Printf("failed to retrive medicines: %v\n", err)
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
		selected, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
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

		boxesStr, _ := pterm.DefaultInteractiveTextInput.Show("Boxes to add")
		boxes, err := strconv.Atoi(boxesStr)
		if err != nil || boxes < 1 {
			pterm.Warning.Println("Invalid number of boxes")
			continue
		}

		units, err := model.CreateStockLog(db, med, boxes)
		if err != nil {
			pterm.Error.Printf("failed to create stock log: %v\n", err)
			continue
		}
		if err := model.IncrementActiveIngredientStock(db, med.RelatedATC, units); err != nil {
			pterm.Error.Printf("failed to update stock log: %v\n", err)
			continue
		}
		pterm.Success.Printf("Added %d boxes of %s", boxes, med.Name)

		cont, _ := pterm.DefaultInteractiveConfirm.WithDefaultValue(true).Show("Add more?")
		if !cont {
			break
		}
	}

	summaryScreen()
}

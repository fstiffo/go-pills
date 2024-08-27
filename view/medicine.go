package view

import "github.com/pterm/pterm"

func addMedicineScreen() {
	clear()
	pterm.DefaultHeader.WithFullWidth().Println("Add Medicine")
}

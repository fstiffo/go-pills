package view

import "github.com/pterm/pterm"

func updatePrescriptionsScreen() {
	clear()
	pterm.DefaultHeader.WithFullWidth().Println("UPDATE PRESCRIPTIONS")
}

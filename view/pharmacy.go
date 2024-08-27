package view

import "github.com/pterm/pterm"

func updatePharmacyScreen() {
	clear()
	pterm.DefaultHeader.WithFullWidth().Println("Update Pharmacy")
}

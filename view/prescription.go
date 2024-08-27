package view

import "github.com/pterm/pterm"

func updatePrescriptionScreen() {
	clear()
	pterm.DefaultHeader.WithFullWidth().Println("Update Prescription")
}

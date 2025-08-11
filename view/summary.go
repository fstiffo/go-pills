package view

import (
	"github.com/pterm/pterm"
)

func summaryScreen() {
	clearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("SUMMARY")
	pterm.Println("\nPrescriptions:")

	ShowPrescriptionsSummaryTable()
	// pterm.Println("\nLast logs refresh: ", model.LastRefresh(control.GetDB()))
}

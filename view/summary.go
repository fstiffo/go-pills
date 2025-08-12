package view

import (
	"github.com/pterm/pterm"
)

func summaryScreen() {
	clearScreen()

	pterm.Println("\nSummary:")

	ShowPrescriptionsSummaryTable()
	// pterm.Println("\nLast logs refresh: ", model.LastRefresh(control.GetDB()))
}

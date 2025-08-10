package view

import (
	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/pterm/pterm"
)

func summaryScreen() {
	clearScreen()
	pterm.DefaultHeader.WithFullWidth().Println("SUMMARY")
	pterm.Println("\nPrescriptions:")

	summaries := model.GetPrescriptionsSummary(control.GetDB())
	tableData := prescriptionSummaryTableData(summaries)
	_ = pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithBoxed().WithData(tableData).Render()
	// pterm.Println("\nLast logs refresh: ", model.LastRefresh(control.GetDB()))
}

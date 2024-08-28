package view

import (
	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/pterm/pterm"
)

func summaryScreen() {
	clear()
	pterm.DefaultHeader.WithFullWidth().Println("SUMMARY")

	tableData := model.GetPrescriptionsSummary(control.GetDB())
	pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithBoxed().WithData(tableData).Render()
}

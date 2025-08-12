package view

func overviewScreen() {
	clearScreen()
	ShowOverviewTable()
	// pterm.Println("\nLast logs refresh: ", model.LastRefresh(control.GetDB()))
}

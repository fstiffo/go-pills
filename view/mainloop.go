package view

import (
	"fmt"

	"github.com/fstiffo/go-pills/control"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

// MainLoop is the main loop view function
func MainLoop() {
	introScreen()
	for {
		switch control.GetScreen() {
		case control.SummaryScreen:
			summaryScreen()
		case control.UpdatePharmacyScreen:
			updatePharmacyScreen()
		case control.UpdatePrescriptionScreen:
			updatePrescriptionsScreen()
		case control.AddMedicineScreen:
			addMedicineScreen()
		}
		menu()
		if err := control.WaitForCommand(); err != nil {
			break
		}
	}
}

func introScreen() {
	clear()
	pillsLogo, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("PILLS", pterm.NewStyle(pterm.FgBlue))).Srender()

	pterm.DefaultCenter.Print(pillsLogo)
	pterm.DefaultCenter.Println("Quickly take control of your supply of pills")
	pterm.DefaultCenter.Println("(Press ENTER to continue)")
	fmt.Scanln()
}

func menu() {
	menuHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgLightYellow, pterm.BgBlue),
		BackgroundStyle: pterm.NewStyle(pterm.BgBlue),
		Margin:          0,
	}
	menuHeader.Println(" F1: Summary   F2: Update Pharmacy   F3: Update Prescriptions   F4: Add Medicine Boxes   F5: Refresh   ESC: Exit ")
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

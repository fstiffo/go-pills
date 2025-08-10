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
		default:
			panic("unhandled default case")
		}
		menu()
		if err := control.WaitForCommand(); err != nil {
			break
		}
	}
}

func introScreen() {
	clearScreen()
	pillsLogo, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("PILLS", pterm.NewStyle(pterm.FgBlue))).Srender()

	pterm.DefaultCenter.Print(pillsLogo)
	pterm.DefaultCenter.Println("Quickly take control of your supply of pills")
	pterm.DefaultCenter.Println("(Press ENTER to continue)")
	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}

func menu() {
	pterm.
		NewStyle(pterm.FgLightYellow, pterm.BgBlue).
		Println("[(S)ummary] [Update (P)harmacy] [Update P(r)escriptions] [(A)dd Medicine Boxes] [Re(f)resh] [(Q)uit]")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

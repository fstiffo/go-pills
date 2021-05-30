package ui

import (
	"github.com/pterm/pterm"
)

func IntroScreen() {
	pillsLog, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("ills", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()
	pterm.DefaultCenter.Print(pillsLog)

	pterm.DefaultCenter.Print("Quickly take control of your supply of pills")

}

func Clear() {
	print("\033[H\033[2J")
}

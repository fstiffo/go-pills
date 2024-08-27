package view

import (
	"time"

	"github.com/pterm/pterm"
)

func summaryScreen() {
	clear()
	pterm.DefaultHeader.WithFullWidth().Println("Summary")

	pterm.Info.Println("Lorem ipsum dolor sit amet, consectetur adipiscing elit." +
		"\nSed do eiusmod tempor incididunt ut labore et dolore magna aliqua." +
		"\nUt enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat." +
		"\nDuis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur." +
		"\nExcepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." +
		"\n" +
		"Summary was updated at: " + pterm.Green(time.Now().Format("2006-01-02 15:04")))
}

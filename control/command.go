package control

import (
	"errors"
	"strings"

	"github.com/pterm/pterm"
)

// WaitForCommand waits for a command from the user
func WaitForCommand() error {
	for {
		input, err := pterm.DefaultInteractiveTextInput.Show("Choose a command")

		if err != nil {
			return err
		}

		command := strings.ToLower(strings.TrimSpace(input))
		if len(command) == 0 {
			continue
		}

		switch command[0] {
		case 's':
			return handleCommand(Overview)
		case 'p':
			return handleCommand(UpdatePharmacy)
		case 'r':
			return handleCommand(UpdatePrescription)
		case 'a':
			return handleCommand(AddMedicine)
		case 'f':
			return handleCommand(Refresh)
		case 'q':
			return handleCommand(Exit)
		default:
			pterm.Warning.Println("Invalid command")
		}
	}
}

func handleCommand(c Command) error {
	appState.lastCommand = c
	switch c {
	case Overview:
		appState.screen = OverviewScreen
	case UpdatePharmacy:
		appState.screen = UpdatePharmacyScreen
	case UpdatePrescription:
		appState.screen = UpdatePrescriptionScreen
	case AddMedicine:
		appState.screen = AddMedicineScreen
	case Refresh:
		if err := RefreshData(); err != nil {
			pterm.Error.Printf("Failed to refresh. Error:%v\n", err)
		} else {
			pterm.Success.Println("Stocked units updated")
			options := []string{"Continue"}
			_, _ = pterm.DefaultInteractiveContinue.WithOptions(options).WithDefaultText("").Show()
		}
		appState.screen = OverviewScreen
	case Exit:
		return errors.New("exit")
	}
	return nil
}

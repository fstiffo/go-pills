package control

import (
	"errors"
	"strings"

	"github.com/pterm/pterm"
)

// WaitForCommand prompts the user to enter a command interactively and executes the corresponding action.
// It continues to prompt until a valid command is entered or an exit condition is met.
// Supported commands:
//   - 's': Show overview
//   - 'p': Update pharmacy
//   - 'r': Update prescription
//   - 'a': Add medicine
//   - 'f': Refresh
//   - 'q': Exit
//
// Returns an error if input fails or if the command handler returns an error.
func WaitForCommand() error {
	for {
		input, err := pterm.DefaultInteractiveTextInput.Show("Choose a command")

		if err != nil {
			return err
		}

		command := strings.ToLower(strings.TrimSpace(input))
		if len(command) == 0 {
			return handleCommand(Exit)
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

// handleCommand processes the given Command and updates the application state accordingly.
// It sets the current screen based on the command, handles data refresh, and manages exit logic.
// Returns an error if the Exit command is received or if a refresh operation fails.
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

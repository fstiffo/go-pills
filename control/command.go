package control

import (
	"errors"
	"strings"

	"github.com/pterm/pterm"
)

// WaitForCommand waits for a command from the user
func WaitForCommand() error {
	input, err := pterm.DefaultInteractiveTextInput.Show("Choose a command:")

	if err != nil {
		return err
	}

	command := strings.ToLower(strings.TrimSpace(input))
	if len(command) == 0 {
		return nil
	}

	switch command[0] {
	case 's':
		return handleCommand(Summary)
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
	}

	return nil
}

func handleCommand(c Command) error {
	appState.lastCommand = c
	switch c {
	case Summary:
		appState.screen = SummaryScreen
	case UpdatePharmacy:
		appState.screen = UpdatePharmacyScreen
	case UpdatePrescription:
		appState.screen = UpdatePrescriptionScreen
	case AddMedicine:
		appState.screen = AddMedicineScreen
	case Refresh:
		appState.screen = SummaryScreen
	case Exit:
		return errors.New("exit")
	}
	return nil
}

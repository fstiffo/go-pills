package control

import (
	"errors"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

// WaitForCommand waits for a command from the user
func WaitForCommand() error {
	f := func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.CtrlC, keys.Escape:
			return true, handleCommand(Exit) // Return true to stop listener
		case keys.RuneKey:
		case keys.F1:
			return true, handleCommand(Summary)
		case keys.F2:
			return true, handleCommand(UpdatePharmacy)
		case keys.F3:
			return true, handleCommand(UpdatePrescription)
		case keys.F4:
			return true, handleCommand(AddMedicine)
		case keys.F5:
			return true, handleCommand(Refresh)
		default:
		}

		return false, nil // Return false to continue listening
	}

	return keyboard.Listen(f)
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

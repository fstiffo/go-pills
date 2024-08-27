package control

// Command is a type for the different commands that the user can execute
type Command int

const (
	// Summary command
	Summary Command = iota
	// UpdatePharmacy command
	UpdatePharmacy
	// UpdatePrescription command
	UpdatePrescription
	// AddMedicine command
	AddMedicine
	// Refresh command
	Refresh
	// Exit command
	Exit
)

// Screen is a type for the different screens that the user can see
type Screen int

const (
	// IntroScreen is the first screen
	IntroScreen Screen = iota
	// SummaryScreen is the second screen
	SummaryScreen
	// UpdatePharmacyScreen is the third screen
	UpdatePharmacyScreen
	// UpdatePrescriptionScreen is the fourth screen
	UpdatePrescriptionScreen
	// AddMedicineScreen is the fifth screen
	AddMedicineScreen
)

type applicationState struct {
	lastCommand Command
	screen      Screen
}

// AppState is the application state
var appState = applicationState{screen: SummaryScreen, lastCommand: Summary}

// GetScreen get actual screen
func GetScreen() Screen {
	return appState.screen
}

package control

import (
	"fmt"
	"github.com/fstiffo/go-pills/model"
	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

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
	db          *gorm.DB
	lastCommand Command
	screen      Screen
}

// AppState is the application state
var appState = applicationState{screen: SummaryScreen, lastCommand: Summary}

// SetDB creates a new application state
func SetDB(db *gorm.DB) {
	appState.db = db
}

// GetDB get actual database
func GetDB() *gorm.DB {
	return appState.db
}

// GetScreen get actual screen
func GetScreen() Screen {
	return appState.screen
}

// RefreshData updates the stock levels of all active ingredients.
func RefreshData() error {
	var activeIngredients []model.ActiveIngredient
	if err := appState.db.Find(&activeIngredients).Error; err != nil {
		return err
	}

	var errs []error
	for i := range activeIngredients {
		if err := model.UpdateStockedUnitsFromIntake(appState.db, &activeIngredients[i]); err != nil {
			// It's better to log the error and continue with the next active ingredient.
			pterm.Error.Printf("Failed to update stock for %s: %v\n", activeIngredients[i].Name, err)
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to update stock for %d ingredients", len(errs))
	}

	return nil
}

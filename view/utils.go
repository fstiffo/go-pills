package view

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/validation"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/pterm/pterm"
	"gorm.io/gorm"
)

// criticalStockInDays is a constant used to alert if stock in days is less than this value.
const criticalStockInDays = 15

// ShowErrorWithConfirm displays an error message and waits for user confirmation before continuing.
func ShowErrorWithConfirm(format string, args ...interface{}) {
	pterm.Error.Printf(format, args...)
	pterm.Info.Println("Press ENTER to continue...")
	fmt.Scanln()
}

// formatCompactDate formats a date as MM/DD/YY for compact display
func formatCompactDate(date time.Time) string {
	return date.Format("01/02/06")
}

// formatCompactDosage formats dosage without unnecessary decimals
func formatCompactDosage(dosage int64, unit string) string {
	value := float64(dosage) / 1000
	if value == float64(int64(value)) {
		return fmt.Sprintf("%.0f %s", value, unit)
	}
	return fmt.Sprintf("%.2f %s", value, unit)
}

// formatCompactFrequency formats frequency in compact d format
func formatCompactFrequency(freq int) string {
	if freq == 1 {
		return "1d"
	}
	return fmt.Sprintf("%dd", freq)
}

// formatCompactText truncates text to fit in table with specified max length
func formatCompactText(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}
	// Reserve 3 characters for "..."
	if maxLength <= 3 {
		return "..."[:maxLength]
	}
	return text[:maxLength-3] + "..."
}

// ShowPrescriptionsSummaryTable retrieves and displays a comprehensive summary of all prescriptions in a formatted table.
func ShowPrescriptionsSummaryTable() {
	summaries := model.GetPrescriptionsSummary(control.GetDB())

	t := table.NewWriter()
	t.SetTitle("Prescriptions Summary")
	t.SetStyle(table.StyleColoredDark)

	// Set column alignment and colors for comprehensive prescriptions table
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter}, // ATC
		{Number: 2, Align: text.AlignLeft},   // Active Ingredient
		{Number: 3, Align: text.AlignRight},  // Dosage
		{Number: 4, Align: text.AlignCenter}, // Frequency
		{Number: 5, Align: text.AlignCenter}, // Valid from
		{Number: 6, Align: text.AlignCenter}, // Last intake update
		{Number: 7, Align: text.AlignCenter}, // Last stocked
		{Number: 8, Align: text.AlignRight},  // Stock in days
	})

	// Add header
	t.AppendHeader(table.Row{"ATC", "Ingredient", "Dosage", "Freq", "Valid From", "Last Intake", "Stocked", "Days Left"})

	// Add data rows
	for _, p := range summaries {
		dosage := formatCompactDosage(p.Dosage, p.Unit)
		frequency := formatCompactFrequency(p.DosingFrequency)
		validFrom := "-"
		if p.StartDate.Valid {
			validFrom = formatCompactDate(p.StartDate.Time)
		}
		lastIntake := "-"
		if p.LastIntakeUpdate.Valid {
			lastIntake = formatCompactDate(p.LastIntakeUpdate.Time)
		}
		lastStock := "-"
		if p.LastStockUpdate.Valid {
			lastStock = formatCompactDate(p.LastStockUpdate.Time)
		}
		stockInDays := strconv.FormatInt(p.StockInDays, 10)
		if p.StockInDays < criticalStockInDays {
			stockInDays = text.Colors{text.FgRed, text.Bold}.Sprint(stockInDays + "⚠️")
		}

		t.AppendRow(table.Row{p.ATC, p.Name, dosage, frequency, validFrom, lastIntake, lastStock, stockInDays})
	}

	fmt.Println(t.Render())
}

// ShowOverviewTable retrieves and displays a compact overview of all prescriptions with mixed column alignment.
func ShowOverviewTable() {
	summaries := model.GetPrescriptionsSummary(control.GetDB())

	t := table.NewWriter()
	today := time.Now().Format("2006-01-02")
	t.SetTitle(fmt.Sprintf("Stock Overview%58s", today))
	t.SetStyle(table.StyleColoredDark)

	// Set column alignment: Name left, numbers right
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft},   // Active Ingredient
		{Number: 2, Align: text.AlignRight},  // Dosage
		{Number: 3, Align: text.AlignCenter}, // Frequency
		{Number: 4, Align: text.AlignCenter}, // Last intake update
		{Number: 5, Align: text.AlignCenter}, // Last stocked
		{Number: 6, Align: text.AlignRight},  // Stock in days
	})

	// Add header
	t.AppendHeader(table.Row{"Ingredient", "Dosage", "Freq", "Last Intake", "Stocked", "Days Left"})

	// Add data rows
	for _, p := range summaries {
		dosage := formatCompactDosage(p.Dosage, p.Unit)
		frequency := formatCompactFrequency(p.DosingFrequency)
		lastIntake := "-"
		if p.LastIntakeUpdate.Valid {
			lastIntake = formatCompactDate(p.LastIntakeUpdate.Time)
		}
		lastStock := "-"
		if p.LastStockUpdate.Valid {
			lastStock = formatCompactDate(p.LastStockUpdate.Time)
		}
		stockInDays := strconv.FormatInt(p.StockInDays, 10)
		if p.StockInDays < criticalStockInDays {
			stockInDays = text.Colors{text.FgRed, text.Bold}.Sprint(stockInDays + "⚠️")
		}

		t.AppendRow(table.Row{p.Name, dosage, frequency, lastIntake, lastStock, stockInDays})
	}

	fmt.Println(t.Render())
}

// ShowMedicinesSummaryTable retrieves and displays a summary of all medicines in a formatted table with mixed column alignment.
func ShowMedicinesSummaryTable() {
	summaries := model.GetMedicinesSummary(control.GetDB())

	t := table.NewWriter()
	t.SetTitle("Medicines Summary")
	t.SetStyle(table.StyleColoredDark)

	// Set column alignment
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft},   // Name
		{Number: 2, Align: text.AlignLeft},   // MAH
		{Number: 3, Align: text.AlignCenter}, // ATC
		{Number: 4, Align: text.AlignCenter}, // AIC
		{Number: 5, Align: text.AlignRight},  // Dosage
		{Number: 6, Align: text.AlignLeft},   // Package
		{Number: 7, Align: text.AlignLeft},   // Form
		{Number: 8, Align: text.AlignRight},  // Box Size
	})

	// Add header
	t.AppendHeader(table.Row{"Name", "MAH", "ATC", "AIC", "Dosage", "Pkg", "Form", "Box"})

	// Add data rows
	for _, med := range summaries {
		dosage := formatCompactDosage(med.Dosage, string(med.Unit))
		mah := formatCompactText(med.MAH, 15)
		form := formatCompactText(med.Form, 20)
		pkg := formatCompactText(med.Package, 8)
		t.AppendRow(table.Row{
			med.Name,
			mah,
			med.RelatedATC,
			med.AIC,
			dosage,
			pkg,
			form,
			strconv.Itoa(med.BoxSize),
		})
	}

	fmt.Println(t.Render())
}

func getOrPromptActiveIngredient(atc string) (*model.ActiveIngredient, error) {
	ai, err := model.GetActiveIngredientByATC(control.GetDB(), atc)
	if err == nil {
		return ai, nil // Happy path: found
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// It's not a "not found" error, so it's an unexpected error.
		return nil, err
	}

	// It is a "not found" error, so we proceed to create it.
	pterm.Warning.Printf("Active ingredient with ATC %s not found. Let's add it.\n", atc)
	return promptAndInsertActiveIngredient(atc)
}

func promptAndInsertActiveIngredient(atc string) (*model.ActiveIngredient, error) {
	name, _ := promptAndValidate("Enter active ingredient name", validation.ValidateName, true)
	unit, _ := promptAndValidate("Enter unit ("+validation.AllowedUnits()+")", validation.ValidateUnit, true)

	cont, _ := pterm.DefaultInteractiveConfirm.WithDefaultValue(true).Show("\nInsert new active ingredient?")
	if !cont {
		return nil, fmt.Errorf("Insert aborted")
	}

	newAI := &model.ActiveIngredient{
		Name: name,
		ATC:  atc,
		Unit: unit,
	}
	if err := model.InsertActiveIngredient(control.GetDB(), newAI); err != nil {
		return nil, err
	}
	pterm.Success.Printf("Active ingredient %s added.\n", name)
	return newAI, nil
}

func promptAndValidate[T any](prompt string, validateFunc func(string) (T, error), emptyIsValid bool) (T, error) {
	for {
		input, _ := pterm.DefaultInteractiveTextInput.Show(prompt)
		if input == "" && emptyIsValid {
			var zero T
			return zero, nil
		}
		value, err := validateFunc(input)
		if err != nil {
			pterm.Error.Println(err)
			continue
		}
		return value, nil
	}
}

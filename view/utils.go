package view

import (
	"errors"
	"fmt"
	"strconv"

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

// ShowPrescriptionsSummaryTable retrieves and displays a summary of all prescriptions in a formatted table.
func ShowPrescriptionsSummaryTable() {
	summaries := model.GetPrescriptionsSummary(control.GetDB())
	tableData := PrescriptionSummaryTableData(summaries)
	_ = pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithBoxed().WithData(tableData).Render()
}

// ShowSummaryTable retrieves and displays a compact summary of all prescriptions with mixed column alignment.
func ShowSummaryTable() {
	summaries := model.GetPrescriptionsSummary(control.GetDB())

	t := table.NewWriter()
	t.SetTitle("Prescriptions Summary")

	// Set column alignment: Name left, numbers right
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft},   // Active Ingredient
		{Number: 2, Align: text.AlignRight},  // Dosage
		{Number: 3, Align: text.AlignCenter}, // Frequency
		{Number: 4, Align: text.AlignCenter}, // Last stocked
		{Number: 5, Align: text.AlignRight},  // Stock in days
	})

	// Add header
	t.AppendHeader(table.Row{"Active Ingredient", "Dosage", "Frequency", "Last stocked", "Stock in days"})

	// Add data rows
	for _, p := range summaries {
		dosage := fmt.Sprintf("%.2f %s", float64(p.Dosage)/1000, p.Unit)
		dayOrDays := " day"
		if p.DosingFrequency > 1 {
			dayOrDays = " days"
		}
		frequency := strconv.Itoa(p.DosingFrequency) + dayOrDays
		lastStock := "-"
		if p.LastStockUpdate.Valid {
			lastStock = p.LastStockUpdate.Time.Format("2006-01-02")
		}
		stockInDays := strconv.FormatInt(p.StockInDays, 10)
		if p.StockInDays < criticalStockInDays {
			stockInDays += " ⚠️"
		}

		t.AppendRow(table.Row{p.Name, dosage, frequency, lastStock, stockInDays})
	}

	fmt.Println(t.Render())
}

// PrescriptionSummaryTableData prepares the data for displaying prescription summaries in a pterm table.
func PrescriptionSummaryTableData(ps []model.PrescriptionSummary) pterm.TableData {
	tableData := pterm.TableData{
		{"ATC", "Active Ingredient", "Dosage", "Frequency", "Valid from", "Last intake update", "Last stocked", "Stock in days"},
	}
	for _, p := range ps {
		dosage := fmt.Sprintf("%.2f %s", float64(p.Dosage)/1000, p.Unit)
		dayOrDays := " day"
		if p.DosingFrequency > 1 {
			dayOrDays = " days"
		}
		frequency := strconv.Itoa(p.DosingFrequency) + dayOrDays
		validFrom := "-"
		if p.StartDate.Valid {
			validFrom = p.StartDate.Time.Format("2006-01-02")
		}
		lastIntake := "-"
		if p.LastIntakeUpdate.Valid {
			lastIntake = p.LastIntakeUpdate.Time.Format("2006-01-02")
		}
		lastStock := "-"
		if p.LastStockUpdate.Valid {
			lastStock = p.LastStockUpdate.Time.Format("2006-01-02")
		}
		stockInDays := strconv.FormatInt(p.StockInDays, 10)
		if p.StockInDays < criticalStockInDays {
			stockInDays += "<--" // Alert
		}
		tableData = append(tableData, []string{p.ATC, p.Name, dosage, frequency, validFrom, lastIntake, lastStock, stockInDays})
	}
	return tableData
}

// ShowMedicinesSummaryTable retrieves and displays a summary of all medicines in a formatted table with mixed column alignment.
func ShowMedicinesSummaryTable() {
	summaries := model.GetMedicinesSummary(control.GetDB())
	
	t := table.NewWriter()
	t.SetTitle("Medicines Summary")
	
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
	t.AppendHeader(table.Row{"Name", "MAH", "ATC", "AIC", "Dosage", "Package", "Form", "Box Size"})
	
	// Add data rows
	for _, med := range summaries {
		dosage := fmt.Sprintf("%.2f %s", float64(med.Dosage)/1000, med.Unit)
		t.AppendRow(table.Row{
			med.Name,
			med.MAH,
			med.RelatedATC,
			med.AIC,
			dosage,
			med.Package,
			med.Form,
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

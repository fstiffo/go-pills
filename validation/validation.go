package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fstiffo/go-pills/model"
	"github.com/shopspring/decimal"
)

var atcCodeRegex = regexp.MustCompile(`^[A-Z][0-9]{2}[A-Z][A-Z][0-9]{2}$`)
var aicCodeRegex = regexp.MustCompile(`^[0-9]{9}$`)

// ValidateName checks if the input string is a valid name.
// A valid name is not empty or whitespace.
func ValidateName(name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	return name, nil
}

// ValidateATC checks if the input string is a valid ATC code.
func ValidateATC(atc string) (string, error) {
	if !atcCodeRegex.MatchString(atc) {
		return "", fmt.Errorf("invalid ATC code format: %s", atc)
	}
	return atc, nil
}

// ValidateAIC checks if the input string is a valid AIC code.
func ValidateAIC(aic string) (string, error) {
	if !aicCodeRegex.MatchString(aic) {
		return "", fmt.Errorf("invalid AIC code format: %s", aic)
	}
	return aic, nil
}

// ValidateDosage checks if the input string is a valid dosage.
// A valid dosage is a positive number returned as a decimal value.
func ValidateDosage(input string) (decimal.Decimal, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return decimal.Zero, fmt.Errorf("dosage cannot be empty")
	}
	value, err := decimal.NewFromString(trimmed)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid number")
	}
	if !value.GreaterThan(decimal.Zero) {
		return decimal.Zero, fmt.Errorf("dosage must be positive, got %s", value.String())
	}
	return value, nil
}

// ValidateFrequency checks if the input string is a valid frequency.
// A valid frequency is a positive integer.
func ValidateFrequency(input string) (int, error) {
	freq, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid number")
	}
	if freq <= 0 {
		return 0, fmt.Errorf("frequency must be > 0, got %d", freq)
	}
	return freq, nil
}

// ValidateBoxSize checks if the input string is a valid box size.
// A valid box size is a positive integer.
func ValidateBoxSize(input string) (int, error) {
	freq, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid number")
	}
	if freq <= 0 {
		return 0, fmt.Errorf("box size must be > 0, got %d", freq)
	}
	return freq, nil
}

// ValidateUnit checks if the input string is a valid unit.
func ValidateUnit(u string) (model.Unit, error) {
	unit := strings.ToUpper(u)
	for _, v := range model.Units() {
		if strings.ToUpper(v) == unit {
			return model.Unit(unit), nil
		}
	}

	var zero model.Unit
	return zero, fmt.Errorf("invalid unit: %s (allowed: %s)", u, AllowedUnits())
}

// AllowedUnits returns a string with the allowed units.
func AllowedUnits() string {
	return strings.Join(model.Units(), ", ")
}

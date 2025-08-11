package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fstiffo/go-pills/model"
)

var atcCodeRegex = regexp.MustCompile(`^[A-Z][0-9]{2}[A-Z][A-Z][0-9]{2}$`)
var aicCodeRegex = regexp.MustCompile(`^[0-9]{9}$`)

func ValidateATC(atc string) error {
	if !atcCodeRegex.MatchString(atc) {
		return fmt.Errorf("invalid ATC code format: %s", atc)
	}
	return nil
}

func ValidateAIC(aic string) error {
	if !aicCodeRegex.MatchString(aic) {
		return fmt.Errorf("invalid AIC code format: %s", aic)
	}
	return nil
}

func ValidateDosage(dosage int64) error {
	if dosage <= 0 {
		return fmt.Errorf("dosage must be positive, got %d", dosage/1000)
	}
	return nil
}

func ValidateFrequency(freq int) error {
	if freq <= 0 {
		return fmt.Errorf("frequency must be > 0, got %d", freq)
	}
	return nil
}

func ValidateUnit(u string) error {
	unit := strings.ToUpper(u)
	for _, v := range model.Units() {
		if strings.ToUpper(v) == unit {
			return nil
		}
	}

	allowed := strings.Join(model.Units(), ", ")
	return fmt.Errorf("invalid unit: %s (allowed: %s)", u, allowed)
}

package validation

import (
	"fmt"
	"regexp"
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
		return fmt.Errorf("dosage must be positive, got %d", dosage)
	}
	return nil
}

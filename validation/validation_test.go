package validation_test

import (
	"testing"

	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/validation"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestValidateName(t *testing.T) {
	name, err := validation.ValidateName("Acetylsalicylic Acid")
	assert.NoError(t, err)
	assert.Equal(t, "Acetylsalicylic Acid", name)

	_, err = validation.ValidateName("   \t\n  ")
	assert.Error(t, err)
}

func TestValidateATC(t *testing.T) {
	atc, err := validation.ValidateATC("A10BA02")
	assert.NoError(t, err)
	assert.Equal(t, "A10BA02", atc)

	_, err = validation.ValidateATC("INVALID")
	assert.Error(t, err)
}

func TestValidateAIC(t *testing.T) {
	aic, err := validation.ValidateAIC("012345678")
	assert.NoError(t, err)
	assert.Equal(t, "012345678", aic)

	_, err = validation.ValidateAIC("12345")
	assert.Error(t, err)
}

func TestValidateDosage(t *testing.T) {
	d, err := validation.ValidateDosage("7.5")
	assert.NoError(t, err)
	assert.True(t, d.Equal(decimal.RequireFromString("7.5")))

	_, err = validation.ValidateDosage("0")
	assert.Error(t, err)

	_, err = validation.ValidateDosage("not-a-number")
	assert.Error(t, err)
}

func TestValidateFrequency(t *testing.T) {
	freq, err := validation.ValidateFrequency("3")
	assert.NoError(t, err)
	assert.Equal(t, 3, freq)

	_, err = validation.ValidateFrequency("0")
	assert.Error(t, err)

	_, err = validation.ValidateFrequency("abc")
	assert.Error(t, err)
}

func TestValidateBoxSize(t *testing.T) {
	size, err := validation.ValidateBoxSize("24")
	assert.NoError(t, err)
	assert.Equal(t, 24, size)

	_, err = validation.ValidateBoxSize("-1")
	assert.Error(t, err)
}

func TestValidateUnit(t *testing.T) {
	unit, err := validation.ValidateUnit("mg")
	assert.NoError(t, err)
	assert.Equal(t, model.Unit("MG"), unit)

	unit, err = validation.ValidateUnit("ml")
	assert.NoError(t, err)
	assert.Equal(t, model.Unit("ML"), unit)

	_, err = validation.ValidateUnit("kg")
	assert.Error(t, err)
}

func TestAllowedUnits(t *testing.T) {
	allowed := validation.AllowedUnits()
	assert.Contains(t, allowed, "mg")
	assert.Contains(t, allowed, "ml")
	assert.Contains(t, allowed, "UI")
}

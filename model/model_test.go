package model_test

import (
	"testing"

	"github.com/fstiffo/go-pills/model"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestUnitsContainsExpectedValues(t *testing.T) {
	units := model.Units()
	assert.ElementsMatch(t, []string{"mg", "ml", "UI"}, units)
}

func TestDecimalFieldsHoldPrecision(t *testing.T) {
	ai := model.ActiveIngredient{StockedUnits: decimal.RequireFromString("12.345"), Unit: model.Unit("mg")}
	assert.True(t, ai.StockedUnits.Equal(decimal.RequireFromString("12.345")))

	med := model.Medicine{Dosage: decimal.RequireFromString("37.75")}
	assert.True(t, med.Dosage.Equal(decimal.RequireFromString("37.75")))

	rx := model.Prescription{Dosage: decimal.RequireFromString("2.25")}
	assert.True(t, rx.Dosage.Equal(decimal.RequireFromString("2.25")))

	logEntry := model.StockLog{Units: decimal.RequireFromString("48.125")}
	assert.True(t, logEntry.Units.Equal(decimal.RequireFromString("48.125")))
}

func TestZeroValuesDefaultToZeroDecimals(t *testing.T) {
	var ai model.ActiveIngredient
	assert.True(t, ai.StockedUnits.Equal(decimal.Zero))

	var med model.Medicine
	assert.True(t, med.Dosage.Equal(decimal.Zero))

	var rx model.Prescription
	assert.True(t, rx.Dosage.Equal(decimal.Zero))

	var sl model.StockLog
	assert.True(t, sl.Units.Equal(decimal.Zero))
}

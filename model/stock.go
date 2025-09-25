package model

import (
	"time"

	"github.com/fstiffo/go-pills/utils"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// CreateStockLog creates a StockLog entry for a medicine stocking and returns the units stocked.
func CreateStockLog(db *gorm.DB, med Medicine, boxes int) (decimal.Decimal, error) {
	boxCount := decimal.NewFromInt(int64(med.BoxSize * boxes))
	units := med.Dosage.Mul(boxCount)
	log := StockLog{
		MedicineID: med.ID,
		RelatedATC: med.RelatedATC,
		Boxes:      boxes,
		Units:      units,
		StockedAt:  time.Now(),
	}
	if err := db.Create(&log).Error; err != nil {
		return decimal.Zero, err
	}
	return units, nil
}

// IncrementActiveIngredientStock increments stocked units and updates the last stock update time for an active ingredient.
func IncrementActiveIngredientStock(db *gorm.DB, atc string, units decimal.Decimal, reset bool) error {
	var err error
	if reset {
		err = db.Model(&ActiveIngredient{}).
			Where("atc = ?", atc).
			Updates(map[string]any{
				"stocked_units":        units,
				"last_intake_update":   time.Now(),
				"last_stock_update":    time.Now(),
				"manual_stock_updater": true,
			}).Error
	} else {
		err = db.Model(&ActiveIngredient{}).
			Where("atc = ?", atc).
			Updates(map[string]any{
				"stocked_units":     gorm.Expr("stocked_units + ?", units),
				"last_stock_update": time.Now(),
			}).Error
	}
	return err
}

// UpdateStockedUnitsFromIntake updates the stocked units of an active ingredient based on prescription intake.
func UpdateStockedUnitsFromIntake(db *gorm.DB, ai *ActiveIngredient) error {
	now := time.Now()
	if ai.LastIntakeUpdate.Valid {
		if utils.IsDateAfterOrEqual(ai.LastIntakeUpdate.Time, now) {
			return nil
		}
	}

	if !ai.LastIntakeUpdate.Valid {
		return db.Model(ai).Update("last_intake_update", now).Error
	}

	var calculationStart time.Time
	if ai.LastStockUpdate.Valid {
		calculationStart = utils.ToDateOnly(ai.LastStockUpdate.Time)
	} else {
		calculationStart = utils.ToDateOnly(ai.CreatedAt)
	}
	today := utils.ToDateOnly(now)

	var prescriptions []Prescription
	if err := db.Where("related_atc = ?", ai.ATC).
		Where("(start_date IS NULL OR start_date <= ?) AND (end_date IS NULL OR end_date > ?)", now, now).
		Order("start_date asc").
		Find(&prescriptions).Error; err != nil {
		return err
	}

	if len(prescriptions) == 0 {
		return db.Model(ai).Update("last_intake_update", now).Error
	}

	totalConsumption := decimal.Zero

	// Iterate from calculationStart up to (but not including) today
	for d := calculationStart; d.Before(today); d = d.Add(24 * time.Hour) {
		var dailyConsumption decimal.Decimal
		// Find the prescription active on day `d`
		for _, p := range prescriptions {
			var startDate time.Time
			if p.StartDate.Valid {
				startDate = utils.ToDateOnly(p.StartDate.Time)
			} else {
				// If no start date, treat as active from the beginning of time
				startDate = time.Time{}
			}

			endDate := today.Add(24 * time.Hour) // Default to tomorrow if no end date
			if p.EndDate.Valid {
				endDate = utils.ToDateOnly(p.EndDate.Time)
			}

			// Check if day `d` is within the prescription's active period
			if (d.Equal(startDate) || d.After(startDate)) && d.Before(endDate) {
				// Check if it's a dosing day
				daysSinceStart := int(d.Sub(startDate).Hours() / 24)
				if p.DosingFrequency > 0 && daysSinceStart%p.DosingFrequency == 0 {
					dailyConsumption = dailyConsumption.Add(p.Dosage)
				}
			}
		}
		totalConsumption = totalConsumption.Add(dailyConsumption)
	}

	if totalConsumption.GreaterThan(decimal.Zero) {
		return db.Model(ai).Updates(map[string]interface{}{
			"stocked_units":      gorm.Expr("stocked_units - ?", totalConsumption),
			"last_intake_update": now,
		}).Error
	}

	// Always update the timestamp to avoid recalculating the same period
	return db.Model(ai).Update("last_intake_update", now).Error
}

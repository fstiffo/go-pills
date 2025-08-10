package model

import (
	"time"

	"gorm.io/gorm"
)

// CreateStockLog creates a StockLog entry for a medicine stocking and returns the units stocked.
func CreateStockLog(db *gorm.DB, med Medicine, boxes int) (int64, error) {
	units := med.Dosage * int64(med.BoxSize*boxes)
	log := StockLog{
		MedicineID: med.ID,
		RelatedATC: med.RelatedATC,
		Boxes:      boxes,
		Units:      units,
		StockedAt:  time.Now(),
	}
	if err := db.Create(&log).Error; err != nil {
		return 0, err
	}
	return units, nil
}

// IncrementActiveIngredientStock increments stocked units and updates the last stock update time for an active ingredient.
func IncrementActiveIngredientStock(db *gorm.DB, atc string, units int64, reset bool) error {
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
	var prescriptions []Prescription

	// Find all prescriptions for the given active ingredient that could have been active
	if err := db.Where("related_atc = ?", ai.ATC).
		Where("start_date <= ?", now).
		Order("start_date asc").
		Find(&prescriptions).Error; err != nil {
		return err
	}

	if len(prescriptions) == 0 {
		// If there are no prescriptions, we might still need to update the time to now
		// to prevent re-running on an empty set.
		return db.Model(ai).Update("last_intake_update", now).Error
	}

	calculationStart := ai.LastIntakeUpdate.Time
	if !ai.LastIntakeUpdate.Valid {
		// Per user feedback, if LastIntakeUpdate is null, it has to be considered as Time Zero.
		calculationStart = time.Time{}
	}

	// Truncate times to the beginning of the day for accurate day-by-day calculation
	calculationStart = time.Date(calculationStart.Year(), calculationStart.Month(), calculationStart.Day(), 0, 0, 0, 0, calculationStart.Location())
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	var totalConsumption int64

	// Iterate from calculationStart up to (but not including) today
	for d := calculationStart; d.Before(today); d = d.Add(24 * time.Hour) {
		var dailyConsumption int64
		// Find the prescription active on day `d`
		for _, p := range prescriptions {
			if !p.StartDate.Valid {
				continue
			}

			startDate := time.Date(p.StartDate.Time.Year(), p.StartDate.Time.Month(), p.StartDate.Time.Day(), 0, 0, 0, 0, p.StartDate.Time.Location())

			endDate := today.Add(24 * time.Hour) // Default to tomorrow if no end date
			if p.EndDate.Valid {
				endDate = time.Date(p.EndDate.Time.Year(), p.EndDate.Time.Month(), p.EndDate.Time.Day(), 0, 0, 0, 0, p.EndDate.Time.Location())
			}

			// Check if day `d` is within the prescription's active period
			if (d.Equal(startDate) || d.After(startDate)) && d.Before(endDate) {
				// Check if it's a dosing day
				daysSinceStart := int(d.Sub(startDate).Hours() / 24)
				if p.DosingFrequency > 0 && daysSinceStart%p.DosingFrequency == 0 {
					dailyConsumption += p.Dosage
				}
			}
		}
		totalConsumption += dailyConsumption
	}

	if totalConsumption > 0 {
		return db.Model(ai).Updates(map[string]interface{}{
			"stocked_units":      gorm.Expr("stocked_units - ?", totalConsumption),
			"last_intake_update": now,
		}).Error
	}

	// Always update the timestamp to avoid recalculating the same period
	return db.Model(ai).Update("last_intake_update", now).Error
}

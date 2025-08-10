package model

import (
	"errors"
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
func IncrementActiveIngredientStock(db *gorm.DB, atc string, units int64) error {
	return db.Model(&ActiveIngredient{}).
		Where("atc = ?", atc).
		Updates(map[string]any{
			"stocked_units":     gorm.Expr("stocked_units + ?", units),
			"last_stock_update": time.Now(),
		}).Error
}

// UpdateStockedUnitsFromIntake updates the stocked units of an active ingredient based on prescription intake.
func UpdateStockedUnitsFromIntake(db *gorm.DB, ai *ActiveIngredient) error {
	var activePrescriptions []Prescription
	now := time.Now()

	// Find active prescriptions for the given active ingredient
	if err := db.Where("related_atc = ?", ai.ATC).
		Where("start_date <= ?", now).
		Where("end_date IS NULL OR end_date > ?", now).
		Find(&activePrescriptions).Error; err != nil {
		return err
	}

	// If there are no active prescriptions, there's nothing to do.
	if len(activePrescriptions) == 0 {
		return nil
	}

	// Determine the start date for consumption calculation.
	var startDate time.Time
	if ai.LastIntakeUpdate.Valid {
		startDate = ai.LastIntakeUpdate.Time
	} else {
		// If LastIntakeUpdate is not set, use the earliest start date from active prescriptions.
		if len(activePrescriptions) > 0 {
			startDate = activePrescriptions[0].StartDate.Time
			for _, p := range activePrescriptions {
				if p.StartDate.Time.Before(startDate) {
					startDate = p.StartDate.Time
				}
			}
		} else {
			// Should not happen due to the check above, but as a safeguard.
			return nil
		}
	}

	// Get the IDs of the active prescriptions
	var prescriptionIDs []uint
	for _, p := range activePrescriptions {
		prescriptionIDs = append(prescriptionIDs, p.ID)
	}

	// Find all prescription logs for the active prescriptions between the start date and now.
	var logs []PrescriptionLog
	if err := db.Where("prescription_id IN ?", prescriptionIDs).
		Where("updated_at BETWEEN ? AND ?", startDate, now).
		Order("updated_at ASC").
		Find(&logs).Error; err != nil {
		return err
	}

	// prescriptionState stores the dosage and frequency for each prescription ID. It's initialized
	// with the state at startDate and then updated as we process logs.
	prescriptionState := make(map[uint]Prescription)
	for _, p := range activePrescriptions {
		pCopy := p // Start with current values as a fallback.

		var lastLogBeforeStart PrescriptionLog
		err := db.Where("prescription_id = ? AND updated_at <= ?", p.ID, startDate).
			Order("updated_at DESC").
			First(&lastLogBeforeStart).Error

		if err == nil {
			// Found a log entry that defines the state at startDate
			pCopy.Dosage = lastLogBeforeStart.Dosage
			pCopy.DosingFrequency = lastLogBeforeStart.DosingFrequency
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err // Real DB error
		}
		// If no log is found before startDate, we assume the prescription's
		// current values have been valid since its creation.

		prescriptionState[p.ID] = pCopy
	}

	var totalConsumption int64
	lastEventTime := startDate

	// Process each log entry as a point in time where dosage might have changed.
	for _, log := range logs {
		// Calculate consumption from lastEventTime to the time of the log.
		segmentDuration := log.UpdatedAt.Sub(lastEventTime)
		days := int(segmentDuration.Hours() / 24)

		if days > 0 {
			// Find the total daily dosage for this period.
			var segmentConsumption int64
			for _, p := range activePrescriptions {
				// Consider only prescriptions active in this segment
				if p.StartDate.Time.Before(log.UpdatedAt) && (p.EndDate.Time.IsZero() || p.EndDate.Time.After(lastEventTime)) {
					state := prescriptionState[p.ID]
					if state.DosingFrequency > 0 {
						dosesInSegment := int64(days / state.DosingFrequency)
						segmentConsumption += dosesInSegment * state.Dosage
					}
				}
			}
			totalConsumption += segmentConsumption
		}

		// Update the state of the prescription based on the log.
		if p, ok := prescriptionState[log.PrescriptionID]; ok {
			p.Dosage = log.Dosage
			p.DosingFrequency = log.DosingFrequency
			prescriptionState[log.PrescriptionID] = p
		}

		lastEventTime = log.UpdatedAt
	}

	// Calculate consumption from the last log entry to now.
	segmentDuration := now.Sub(lastEventTime)
	days := int(segmentDuration.Hours() / 24)

	if days > 0 {
		var segmentConsumption int64
		for _, p := range activePrescriptions {
			state := prescriptionState[p.ID]
			if state.DosingFrequency > 0 {
				dosesInSegment := int64(days / state.DosingFrequency)
				segmentConsumption += dosesInSegment * state.Dosage
			}
		}
		totalConsumption += segmentConsumption
	}

	// Update the active ingredient's stock and last intake update time.
	if totalConsumption > 0 {
		return db.Model(ai).
			Updates(map[string]interface{}{
				"stocked_units":      gorm.Expr("stocked_units - ?", totalConsumption),
				"last_intake_update": now,
			}).Error
	}

	return nil
}

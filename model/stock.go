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
func IncrementActiveIngredientStock(db *gorm.DB, atc string, units int64) error {
	return db.Model(&ActiveIngredient{}).
		Where("atc = ?", atc).
		Updates(map[string]any{
			"stocked_units":     gorm.Expr("stocked_units + ?", units),
			"last_stock_update": time.Now(),
		}).Error
}

package model

import (
	"time"

	"gorm.io/gorm"
)

// LastRefresh returns the last time the logs were refreshed.
func LastRefresh(db *gorm.DB) string {
	var il IntakeLog
	var sl StockLog
	intakeErr := db.Model(&IntakeLog{}).
		Order("consumed_at desc").
		Select("consumed_at").
		First(&il).Error
	stockErr := db.Model(&StockLog{}).
		Order("stocked_at desc").
		Select("stocked_at").
		First(&sl).Error

	var zero time.Time

	if intakeErr != nil {
		il.ConsumedAt = zero
	}
	if stockErr != nil {
		sl.StockedAt = zero
	}

	if il.ConsumedAt.IsZero() && sl.StockedAt.IsZero() {
		return "Never"
	}
	if il.ConsumedAt.After(sl.StockedAt) {
		return il.ConsumedAt.Format("2006-01-02 15:04:05")
	}
	return sl.StockedAt.Format("2006-01-02 15:04:05")
}

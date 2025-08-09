package model

import (
	"gorm.io/gorm"
)

// LastRefresh returns the last time the logs were refreshed.
func LastRefresh(db *gorm.DB) string {
	var il IntakeLog
	db.Model(&IntakeLog{}).
		Order("consumed_at desc").
		Select("consumed_at").
		Limit(1).
		Find(&il)

	var sl StockLog
	db.Model(&StockLog{}).
		Order("stocked_at desc").
		Select("stocked_at").
		Limit(1).
		Find(&sl)

	t1 := il.ConsumedAt
	t2 := sl.StockedAt

	if t1.IsZero() && t2.IsZero() {
		return "Never"
	}
	if t1.After(t2) {
		return t1.Format("2006-01-02 15:04:05")
	}
	return t2.Format("2006-01-02 15:04:05")
}

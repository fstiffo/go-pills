package model

import (
	"time"

	"gorm.io/gorm"
)

// LastRefresh returns the last time the logs were refreshed.
func LastRefresh(db *gorm.DB) string {
	type ConsumptionLog struct {
		ConsumedAt time.Time
	}
	var cl ConsumptionLog
	db.First(&cl)
	type StockLog struct {
		StockedAt time.Time
	}
	var sl StockLog
	db.First(&sl)
	if cl.ConsumedAt.After(sl.StockedAt) {
		return cl.ConsumedAt.Format("2006-01-02 15:04:05")
	}
	return sl.StockedAt.Format("2006-01-02 15:04:05")

}

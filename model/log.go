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
	db.Order("consumed_at desc").Select("consumed_at").First(&cl)
	type StockLog struct {
		StockedAt time.Time
	}
	var sl StockLog
	db.Order("stocked_at desc").Select("stocked_at").First(&sl)
	if cl.ConsumedAt.After(sl.StockedAt) {
		return cl.ConsumedAt.Format("2006-01-02 15:04:05")
	}
	return sl.StockedAt.Format("2006-01-02 15:04:05")
}

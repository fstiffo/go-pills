package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/fstiffo/go-pills/model"
)

func main() {
	db, err := gorm.Open(sqlite.Open("pills.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.ActiveIngredient{}, &model.Medicine{}, &model.Prescription{}, &model.ConsumptionLog{}, &model.StockLog{})
}

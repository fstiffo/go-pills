package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/view"
)

func main() {
	view.MainLoop()
	return

	db, err := gorm.Open(sqlite.Open("pills.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema and populate the database
	model.Populate(db)
}

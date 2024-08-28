package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/view"
)

func main() {
	db, err := gorm.Open(sqlite.Open("pills.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	control.SetDB(db)
	view.MainLoop()

	// Migrate the schema and populate the database
	// model.Populate(db)
}

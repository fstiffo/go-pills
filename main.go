package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/fstiffo/go-pills/model"
)

func main() {
	db, err := gorm.Open(sqlite.Open("pills.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema and populate the database
	model.Populate(db)
}

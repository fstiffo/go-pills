package main

import (
	"flag"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/fstiffo/go-pills/control"
	"github.com/fstiffo/go-pills/model"
	"github.com/fstiffo/go-pills/view"
)

func main() {
	reset := flag.Bool("reset", false, "reset database and seed data")
	flag.Parse()

	dbPath := "pills.db"
	_, err := os.Stat(dbPath)
	newDB := os.IsNotExist(err)

	db, err := gorm.Open(sqlite.Open("pills.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	control.SetDB(db)

	if newDB || *reset {
		model.Populate(control.GetDB(), *reset)
	} else {
		if err := model.Migrate(control.GetDB()); err != nil {
			log.Fatalf("failed to migrate schema: %v", err)
		}
	}
	view.MainLoop()
}

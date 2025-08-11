package main

import (
	"flag"
	"fmt"
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

	dbPath := os.Getenv("PILLS_DB_PATH")
	if dbPath == "" {
		dbPath = "pills.db"
	}

	_, err := os.Stat(dbPath)
	newDB := os.IsNotExist(err)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect database: %v\n", err)
		os.Exit(1)
	}
	control.SetDB(db)

	if newDB || *reset {
		if err := model.Populate(control.GetDB(), *reset); err != nil {
			fmt.Printf("failed to populate database: %v\n", err)
			os.Exit(1)
		}
	} else {
		if err := model.Migrate(control.GetDB()); err != nil {
			fmt.Printf("failed to migrate schema: %v\n", err)
			os.Exit(1)
		}
	}
	view.MainLoop()
}

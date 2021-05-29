package main

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Medication struct {
	gorm.Model
	BrandName    string
	Manufacturer string
	DrugName     string
	Strength     uint
	Unity        string
	Quantity     uint
	Purchases    []Purchase
}

type Prescription struct {
	gorm.Model
	Date     time.Time
	DrugName string
	Dose     uint
	Unity    string
	Days     uint
}

type Purchase struct {
	gorm.Model
	Date         time.Time
	MedicationID uint
	Medication   Medication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity     uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("pills.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.Exec("DROP TABLE IF EXISTS medications;DROP TABLE IF EXISTS prescriptions;DROP TABLE IF EXISTS purchases")
	db.AutoMigrate(&Medication{})
	db.AutoMigrate(&Prescription{})
	db.AutoMigrate(&Purchase{})

	// Create
	m1 := Medication{
		Model:        gorm.Model{},
		BrandName:    "Tachiprina",
		Manufacturer: "Angelini",
		DrugName:     "paracetamolo",
		Strength:     1000,
		Unity:        "mg",
		Quantity:     20,
	}
	db.Create(&m1)

	m2 := Medication{
		Model:        gorm.Model{},
		BrandName:    "Cardisoaspirin",
		Manufacturer: "Bayer",
		DrugName:     "acido acetilsalicilico",
		Strength:     100,
		Unity:        "mg",
		Quantity:     30,
	}
	db.Create(&m2)

	db.Create(&Prescription{
		Model:    gorm.Model{},
		Date:     time.Now(),
		DrugName: "acido acetilsalicilico",
		Dose:     100,
		Unity:    "mg",
		Days:     1,
	})
	db.Create(&Prescription{
		Model:    gorm.Model{},
		Date:     time.Now(),
		DrugName: "paracetamolo",
		Dose:     500,
		Unity:    "mg",
		Days:     7,
	})

	db.Create(&Purchase{
		Model:      gorm.Model{},
		Date:       time.Now(),
		Medication: m1,
		Quantity:   1,
	})
	db.Create(&Purchase{
		Model:      gorm.Model{},
		Date:       time.Now(),
		Medication: m2,
		Quantity:   3,
	})
	// Read

	// Update - update product's price to 200
	// Update - update multiple fields

	// Delete - delete product
	db.Select("Purchases").Delete(&m1)
}

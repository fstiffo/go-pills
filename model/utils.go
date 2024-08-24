package model

import (
	"log"

	"gorm.io/gorm"
)

// Populate creates the tables in the database and pupulates them with the necessary data
func Populate(db *gorm.DB) {
	// Clear the database
	db.Migrator().DropTable(&ActiveIngredient{}, &Medicine{}, &Prescription{}, &ConsumptionLog{}, &StockLog{})

	// Migrate the schema
	err := db.AutoMigrate(&ActiveIngredient{}, &Medicine{}, &Prescription{}, &ConsumptionLog{}, &StockLog{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}
	log.Println("Schema migrated")

	// Populate the database with the necessary data
	populateActiveIngredients(db)
	populatePrescriptions(db)
	log.Println("Database populated")
}

func populateActiveIngredients(db *gorm.DB) {
	// Create the active ingredients
	activeIngredients := []ActiveIngredient{
		{Name: "acido acetilsalicilico"},
		{Name: "allopurinolo"},
		{Name: "amlodipina"},
		{Name: "colicalciferolo", Unit: ui},
		{Name: "doxazosina"},
		{Name: "insulina glargine", Unit: u},
		{Name: "metoprololo"},
		{Name: "micofenolato mofetile"},
		{Name: "prednisone"},
		{Name: "zofenopril calcio"},
	}

	result := db.Create(&activeIngredients)

	if result.Error != nil {
		log.Fatalf("failed to populate active ingredients: %v", result.Error)
	}
	log.Printf("Active ingredients populated, %d records inserted", result.RowsAffected)
}

func populatePrescriptions(db *gorm.DB) {
	// Recover the active ingredients
	var activeIngredients []ActiveIngredient
	result := db.Select("ID", "name").Find(&activeIngredients)
	if result.Error != nil {
		log.Fatalf("failed to recover active ingredients: %v", result.Error)
	}

	ingredientMap := make(map[string]uint)
	for _, ingredient := range activeIngredients {
		ingredientMap[ingredient.Name] = ingredient.ID
	}
	prescriptions := []Prescription{
		{ActiveIngredientID: ingredientMap["acido acetilsalicilico"], Dosage: 100, DosageFrequency: 1},
		{ActiveIngredientID: ingredientMap["allopurinolo"], Dosage: 150, DosageFrequency: 1},
		{ActiveIngredientID: ingredientMap["amlodipina"], Dosage: 5, DosageFrequency: 1},
		{ActiveIngredientID: ingredientMap["colicalciferolo"], Dosage: 40, DosageFrequency: 7},
		{ActiveIngredientID: ingredientMap["doxazosina"], Dosage: 2, DosageFrequency: 1},
		{ActiveIngredientID: ingredientMap["insulina glargine"], Dosage: 16, DosageFrequency: 1},
		{ActiveIngredientID: ingredientMap["metoprololo"], Dosage: 50, DosageFrequency: 1},
		{ActiveIngredientID: ingredientMap["micofenolato mofetile"], Dosage: 1500, DosageFrequency: 1},
		{ActiveIngredientID: ingredientMap["prednisone"], Dosage: 7.5, DosageFrequency: 1},
		{ActiveIngredientID: ingredientMap["zofenopril calcio"], Dosage: 30, DosageFrequency: 1},
	}

	result = db.Create(&prescriptions)
	if result.Error != nil {
		log.Fatalf("failed to populate prescriptions: %v", result.Error)
	}
	log.Printf("Prescriptions populated, %d records inserted", result.RowsAffected)
}

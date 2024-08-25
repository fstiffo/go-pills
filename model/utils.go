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
	populateMedicines(db)
	log.Println("Database populated")
}

func populateActiveIngredients(db *gorm.DB) {
	// Create the active ingredients
	activeIngredients := []ActiveIngredient{
		{Name: "acido acetilsalicilico", ATC: "B01AC06"},
		{Name: "allopurinolo", ATC: "M04AA01"},
		{Name: "amlodipina", ATC: "C08CA01"},
		{Name: "colicalciferolo", ATC: "A11CC05", Unit: ui},
		{Name: "doxazosina", ATC: "C02CA04"},
		{Name: "insulina glargine", ATC: "A10AE04", Unit: ui},
		{Name: "metoprololo", ATC: "C07AB02"},
		{Name: "micofenolato mofetile", ATC: "L04AA06"},
		{Name: "prednisone", ATC: "H02AB07"},
		{Name: "zofenopril calcio", ATC: "C09AA15"},
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
	result := db.Select("ATC", "name").Find(&activeIngredients)
	if result.Error != nil {
		log.Fatalf("failed to recover active ingredients: %v", result.Error)
	}

	ingredientMap := make(map[string]string)
	for _, ingredient := range activeIngredients {
		ingredientMap[ingredient.Name] = ingredient.ATC
	}
	prescriptions := []Prescription{
		{RelatedATC: ingredientMap["acido acetilsalicilico"], Dosage: 100 * 1000, DosageFrequency: 1},
		{RelatedATC: ingredientMap["allopurinolo"], Dosage: 150 * 1000, DosageFrequency: 1},
		{RelatedATC: ingredientMap["amlodipina"], Dosage: 5 * 1000, DosageFrequency: 1},
		{RelatedATC: ingredientMap["colicalciferolo"], Dosage: 10000 * 1000, DosageFrequency: 7},
		{RelatedATC: ingredientMap["doxazosina"], Dosage: 2 * 1000, DosageFrequency: 1},
		{RelatedATC: ingredientMap["insulina glargine"], Dosage: 16 * 1000, DosageFrequency: 1},
		{RelatedATC: ingredientMap["metoprololo"], Dosage: 50 * 1000, DosageFrequency: 1},
		{RelatedATC: ingredientMap["micofenolato mofetile"], Dosage: 1500 * 1000, DosageFrequency: 1},
		{RelatedATC: ingredientMap["prednisone"], Dosage: 7500, DosageFrequency: 1},
		{RelatedATC: ingredientMap["zofenopril calcio"], Dosage: 30 * 1000, DosageFrequency: 1},
	}

	result = db.Create(&prescriptions)
	if result.Error != nil {
		log.Fatalf("failed to populate prescriptions: %v", result.Error)
	}
	log.Printf("Prescriptions populated, %d records inserted", result.RowsAffected)
}

func populateMedicines(db *gorm.DB) {
	// Create the medicines
	medicines := []Medicine{
		{Name: "Acido Acetilsalicilico",
			MAH:        "Mylan S.p.A.",
			RelatedATC: "B01AC06",
			AIC:        "047065014",
			Dosage:     100 * 1000,
			Package:    "blister",
			Form:       "compressa gastroresistente",
			BoxSize:    30},
		{Name: "Allopurinolo Sandoz",
			MAH:        "Sandoz S.p.A.",
			RelatedATC: "M04AA01",
			AIC:        "039060292",
			Dosage:     300 * 1000,
			Package:    "blister",
			Form:       "compressa",
			BoxSize:    30},
		{Name: "Norvasc",
			MAH:        "Viatris Pharms S.r.l.",
			RelatedATC: "C08CA01",
			AIC:        "027428010",
			Dosage:     5 * 1000,
			Package:    "blister",
			Form:       "compressa",
			BoxSize:    28},
		{Name: "Colecalciferolo IPSO Pharma",
			MAH:        "IPSO Pharma S.r.l.",
			RelatedATC: "A11CC05",
			AIC:        "043913019",
			Dosage:     1000000 * 1000,
			Package:    "flacone",
			Form:       "gocce orali, soluzione",
			BoxSize:    1},
		{Name: "Doxazosin Auribindo",
			MAH:        "Auribindo Pharma (Italia) S.r.l.",
			RelatedATC: "C02CA04",
			AIC:        "040243180",
			Dosage:     4 * 1000,
			Package:    "blister",
			Form:       "compressa",
			BoxSize:    20},
		{Name: "Toujeo",
			MAH:        "Sanofi-Aventis Deutschland GMBH",
			RelatedATC: "A10AE04",
			AIC:        "043192347",
			Dosage:     450 * 1000,
			Package:    "penna",
			Form:       "sospensione iniettabile 300 unità/ml",
			BoxSize:    3},
		{Name: "Metoprololo DOC Generici",
			MAH:        "DOC Generici S.r.l.",
			RelatedATC: "C07AB02",
			AIC:        "035054055",
			Dosage:     100 * 1000,
			Package:    "blister PVC/Al",
			Form:       "compressa",
			BoxSize:    30},
		{Name: "Micofenolato Mofetile Tillomed",
			MAH:        "Tillomed Italia S.r.l.",
			RelatedATC: "L04AA06",
			AIC:        "045833011",
			Dosage:     500 * 1000,
			Package:    "blister PVC/Al",
			Form:       "compressa rivestita con film",
			BoxSize:    50},
		{Name: "Deltacortene",
			MAH:        "Bruno Farmaceutici S.p.A.",
			RelatedATC: "H02AB07",
			AIC:        "010089047",
			Dosage:     5 * 1000,
			Package:    "blister",
			Form:       "compressa",
			BoxSize:    20},
		{Name: "Zofenopril Mylan Generics",
			MAH:        "Mylan S.p.A.",
			RelatedATC: "C09AA15",
			AIC:        "040724041",
			Dosage:     30 * 1000,
			Package:    "blister PVC/Aclar/Al",
			Form:       "compressa rivestita con film",
			BoxSize:    28},
	}

	result := db.Create(&medicines)
	if result.Error != nil {
		log.Fatalf("failed to populate medicines: %v", result.Error)
	}
	log.Printf("Medicines populated, %d records inserted", result.RowsAffected)
}

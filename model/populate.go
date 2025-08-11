package model

import (
	"log"
	"math"

	"gorm.io/gorm"
)

// Populate creates the tables in the database and populates them with the necessary data
func Populate(db *gorm.DB, reset bool) {

	// Migrate the schema
	if err := resetSchema(db, reset); err != nil {
		log.Fatalf("failed to reset schema: %v", err)
	}

	// Populate the database with the necessary data
	if err := populateActiveIngredients(db); err != nil {
		log.Fatalf("failed to populate active ingredients: %v", err)
	}
	if err := populatePrescriptions(db); err != nil {
		log.Fatalf("failed to populate prescriptions: %v", err)
	}
	if err := populateMedicines(db); err != nil {
		log.Fatalf("failed to populate medicines: %v", err)
	}

	log.Println("Database populated")
}

// Migrate applies database migrations without dropping existing data.
func Migrate(db *gorm.DB) error {
	return resetSchema(db, false)
}

// resetSchema drops the tables in the database and migrates the schema
func resetSchema(db *gorm.DB, reset bool) error {
	tables := []interface{}{&ActiveIngredient{}, &Medicine{}, &Prescription{}, &StockLog{}}
	if reset {
		for _, t := range tables {
			if db.Migrator().HasTable(t) {
				if err := db.Migrator().DropTable(t); err != nil {
					return err
				}
			}
		}
	}

	// Migrate the schema
	if err := db.AutoMigrate(tables...); err != nil {
		return err
	}
	log.Println("Schema migrated")
	return nil
}

func populateActiveIngredients(db *gorm.DB) error {
	// Create the active ingredients
	activeIngredients := []ActiveIngredient{
		{Name: "acido acetilsalicilico", ATC: "B01AC06"},
		{Name: "allopurinolo", ATC: "M04AA01"},
		{Name: "amlodipina", ATC: "C08CA01"},
		{Name: "colecalciferolo", ATC: "A11CC05", Unit: ui},
		{Name: "doxazosina", ATC: "C02CA04"},
		{Name: "insulina glargine", ATC: "A10AE04", Unit: ui},
		{Name: "metoprololo", ATC: "C07AB02"},
		{Name: "micofenolato mofetile", ATC: "L04AA06"},
		{Name: "prednisone", ATC: "H02AB07"},
		{Name: "zofenopril calcio", ATC: "C09AA15"},
	}

	result := db.Create(&activeIngredients)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Active ingredients populated, %d records inserted", result.RowsAffected)
	return nil

}

func populatePrescriptions(db *gorm.DB) error {
	// Recover the active ingredients
	var activeIngredients []ActiveIngredient
	if result := db.Select("ATC", "name").Find(&activeIngredients); result.Error != nil {
		return result.Error
	}

	ingredientMap := make(map[string]string)
	for _, ingredient := range activeIngredients {
		ingredientMap[ingredient.Name] = ingredient.ATC
	}
	prescriptions := []Prescription{
		{RelatedATC: ingredientMap["acido acetilsalicilico"], Dosage: 100 * 1000, DosingFrequency: 1},
		{RelatedATC: ingredientMap["allopurinolo"], Dosage: 150 * 1000, DosingFrequency: 1},
		{RelatedATC: ingredientMap["amlodipina"], Dosage: 5 * 1000, DosingFrequency: 1},
		{RelatedATC: ingredientMap["colecalciferolo"], Dosage: 10000 * 1000, DosingFrequency: 7},
		{RelatedATC: ingredientMap["doxazosina"], Dosage: 2 * 1000, DosingFrequency: 1},
		{RelatedATC: ingredientMap["insulina glargine"], Dosage: 16 * 1000, DosingFrequency: 1},
		{RelatedATC: ingredientMap["metoprololo"], Dosage: 50 * 1000, DosingFrequency: 1},
		{RelatedATC: ingredientMap["micofenolato mofetile"], Dosage: 1500 * 1000, DosingFrequency: 1},
		{RelatedATC: ingredientMap["prednisone"], Dosage: int64(math.Round(7.5 * 1000)), DosingFrequency: 1},
		{RelatedATC: ingredientMap["zofenopril calcio"], Dosage: 30 * 1000, DosingFrequency: 1},
	}

	result := db.Create(&prescriptions)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Prescriptions populated, %d records inserted", result.RowsAffected)
	return nil
}

func populateMedicines(db *gorm.DB) error {
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
		return result.Error
	}
	log.Printf("Medicines populated, %d records inserted", result.RowsAffected)
	return nil
}

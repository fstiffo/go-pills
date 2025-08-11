package model

import (
	"log"

	"gorm.io/gorm"
)

// MedicineSummary contains prescription data for presentation
type MedicineSummary struct {
	Name       string
	MAH        string
	RelatedATC string
	AIC        string
	Dosage     int64
	Unit       Unit
	Package    string
	Form       string
	BoxSize    int
}

// GetMedicinesSummary returns a summary of all medicines
func GetMedicinesSummary(db *gorm.DB) []MedicineSummary {
	type medicine struct {
		Medicine
		Unit Unit
	}
	var ms []medicine
	result := db.Model(&Medicine{}).
		Select("medicines.*, ai.unit").
		Joins("JOIN active_ingredients ai ON ai.atc = medicines.related_atc").
		Order("name").
		Scan(&ms)
	if result.Error != nil {
		log.Fatalf("failed to get prescriptions: %v", result.Error)
	}

	var summaries []MedicineSummary

	for _, m := range ms {
		summaries = append(summaries, MedicineSummary{
			Name:       m.Name,
			MAH:        m.MAH,
			RelatedATC: m.RelatedATC,
			AIC:        m.AIC,
			Dosage:     m.Dosage,
			Unit:       m.Unit,
			Package:    m.Package,
			Form:       m.Form,
			BoxSize:    m.BoxSize,
		})
	}

	return summaries
}

// InsertActiveIngredient inserts a new ActiveIngredient record.
func InsertActiveIngredient(db *gorm.DB, ai *ActiveIngredient) error {
	return db.Create(ai).Error
}

// InsertMedicine inserts a new Medicine record.
func InsertMedicine(db *gorm.DB, med *Medicine) error {
	return db.Create(med).Error
}

// GetActiveIngredientByATC returns an active ingredient by its ATC code.
func GetActiveIngredientByATC(db *gorm.DB, atc string) (*ActiveIngredient, error) {
	var ai ActiveIngredient
	err := db.Where("atc = ?", atc).First(&ai).Error
	return &ai, err
}

// GetAllMedicines returns all medicines from the database.
func GetAllMedicines(db *gorm.DB) ([]Medicine, error) {
	var medicines []Medicine
	err := db.Find(&medicines).Error
	return medicines, err
}

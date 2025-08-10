package model_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/fstiffo/go-pills/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&model.ActiveIngredient{}, &model.Prescription{}, &model.PrescriptionLog{}, &model.Medicine{}, &model.StockLog{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestUpdateStockedUnitsFromIntake(t *testing.T) {
	t.Run("NoActivePrescriptions", func(t *testing.T) {
		db := setupTestDB(t)
		ai := model.ActiveIngredient{Name: "Test AI", ATC: "A10BA02", StockedUnits: 1000}
		db.Create(&ai)

		err := model.UpdateStockedUnitsFromIntake(db, &ai)
		assert.NoError(t, err)

		var updatedAI model.ActiveIngredient
		db.First(&updatedAI, ai.ID)
		assert.Equal(t, int64(1000), updatedAI.StockedUnits)
		assert.False(t, updatedAI.LastIntakeUpdate.Valid)
	})

	t.Run("SinglePrescriptionNoLogs", func(t *testing.T) {
		db := setupTestDB(t)
		ai := model.ActiveIngredient{Name: "Test AI", ATC: "A10BA02", StockedUnits: 100000}
		db.Create(&ai)

		prescription := model.Prescription{
			RelatedATC:      "A10BA02",
			Dosage:          1000, // 1mg
			DosingFrequency: 1,    // every day
			StartDate:       sql.NullTime{Time: time.Now().Add(-10 * 24 * time.Hour), Valid: true},
		}
		db.Create(&prescription)

		// Set LastIntakeUpdate to 5 days ago
		db.Model(&ai).Update("last_intake_update", time.Now().Add(-5*24*time.Hour))

		err := model.UpdateStockedUnitsFromIntake(db, &ai)
		assert.NoError(t, err)

		var updatedAI model.ActiveIngredient
		db.First(&updatedAI, ai.ID)

		// 5 days * 1 dose/day * 1000 units/dose = 5000 units
		assert.Equal(t, int64(95000), updatedAI.StockedUnits)
		assert.True(t, updatedAI.LastIntakeUpdate.Valid)
		assert.WithinDuration(t, time.Now(), updatedAI.LastIntakeUpdate.Time, 2*time.Second)
	})

	t.Run("SinglePrescriptionWithChangingDosage", func(t *testing.T) {
		db := setupTestDB(t)
		ai := model.ActiveIngredient{Name: "Test AI", ATC: "A10BA02", StockedUnits: 100000}
		db.Create(&ai)
		db.Model(&ai).Update("last_intake_update", time.Now().Add(-10*24*time.Hour))

		prescription := model.Prescription{
			RelatedATC:      "A10BA02",
			Dosage:          2000, // Current dosage is 2mg
			DosingFrequency: 1,
			StartDate:       sql.NullTime{Time: time.Now().Add(-20 * 24 * time.Hour), Valid: true},
		}
		db.Create(&prescription)

		// Log history:
		// At T-15, dosage was set to 1000.
		db.Create(&model.PrescriptionLog{
			PrescriptionID:  prescription.ID,
			UpdatedAt:       time.Now().Add(-15 * 24 * time.Hour),
			Dosage:          1000,
			DosingFrequency: 1,
		})
		// At T-5, dosage was set to 2000.
		db.Create(&model.PrescriptionLog{
			PrescriptionID:  prescription.ID,
			UpdatedAt:       time.Now().Add(-5 * 24 * time.Hour),
			Dosage:          2000,
			DosingFrequency: 1,
		})

		err := model.UpdateStockedUnitsFromIntake(db, &ai)
		assert.NoError(t, err)

		var updatedAI model.ActiveIngredient
		db.First(&updatedAI, ai.ID)

		// Calculation from LastIntakeUpdate (T-10):
		// The state at T-10 is defined by the log at T-15 (Dosage: 1000).
		// Segment 1: [T-10, T-5). Duration: 5 days. Dosage: 1000.
		// Consumption 1: 5 days * 1000 units/day = 5000
		// The state at T-5 is defined by the log at T-5 (Dosage: 2000).
		// Segment 2: [T-5, T-0). Duration: 5 days. Dosage: 2000.
		// Consumption 2: 5 days * 2000 units/day = 10000
		// Total consumption: 15000
		// Expected stock: 100000 - 15000 = 85000

		assert.Equal(t, int64(85000), updatedAI.StockedUnits)
		assert.True(t, updatedAI.LastIntakeUpdate.Valid)
		assert.WithinDuration(t, time.Now(), updatedAI.LastIntakeUpdate.Time, 2*time.Second)
	})
}

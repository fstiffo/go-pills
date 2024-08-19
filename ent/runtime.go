// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fstiffo/pills/ent/activeingredient"
	"fstiffo/pills/ent/medicine"
	"fstiffo/pills/ent/prescription"
	"fstiffo/pills/ent/purchase"
	"fstiffo/pills/ent/schema"
	"fstiffo/pills/ent/stockinglog"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	activeingredientFields := schema.ActiveIngredient{}.Fields()
	_ = activeingredientFields
	// activeingredientDescName is the schema descriptor for name field.
	activeingredientDescName := activeingredientFields[0].Descriptor()
	// activeingredient.NameValidator is a validator for the "name" field. It is called by the builders before save.
	activeingredient.NameValidator = activeingredientDescName.Validators[0].(func(string) error)
	medicineFields := schema.Medicine{}.Fields()
	_ = medicineFields
	// medicineDescName is the schema descriptor for name field.
	medicineDescName := medicineFields[0].Descriptor()
	// medicine.NameValidator is a validator for the "name" field. It is called by the builders before save.
	medicine.NameValidator = medicineDescName.Validators[0].(func(string) error)
	// medicineDescMah is the schema descriptor for mah field.
	medicineDescMah := medicineFields[1].Descriptor()
	// medicine.MahValidator is a validator for the "mah" field. It is called by the builders before save.
	medicine.MahValidator = medicineDescMah.Validators[0].(func(string) error)
	// medicineDescDosage is the schema descriptor for dosage field.
	medicineDescDosage := medicineFields[2].Descriptor()
	// medicine.DosageValidator is a validator for the "dosage" field. It is called by the builders before save.
	medicine.DosageValidator = medicineDescDosage.Validators[0].(func(float64) error)
	// medicineDescUnit is the schema descriptor for unit field.
	medicineDescUnit := medicineFields[3].Descriptor()
	// medicine.UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	medicine.UnitValidator = medicineDescUnit.Validators[0].(func(string) error)
	// medicineDescAtc is the schema descriptor for atc field.
	medicineDescAtc := medicineFields[4].Descriptor()
	// medicine.AtcValidator is a validator for the "atc" field. It is called by the builders before save.
	medicine.AtcValidator = func() func(string) error {
		validators := medicineDescAtc.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(atc string) error {
			for _, fn := range fns {
				if err := fn(atc); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// medicineDescBoxSize is the schema descriptor for box_size field.
	medicineDescBoxSize := medicineFields[7].Descriptor()
	// medicine.BoxSizeValidator is a validator for the "box_size" field. It is called by the builders before save.
	medicine.BoxSizeValidator = medicineDescBoxSize.Validators[0].(func(int) error)
	// medicineDescStock is the schema descriptor for stock field.
	medicineDescStock := medicineFields[8].Descriptor()
	// medicine.DefaultStock holds the default value on creation for the stock field.
	medicine.DefaultStock = medicineDescStock.Default.(float32)
	prescriptionFields := schema.Prescription{}.Fields()
	_ = prescriptionFields
	// prescriptionDescDosage is the schema descriptor for dosage field.
	prescriptionDescDosage := prescriptionFields[0].Descriptor()
	// prescription.DosageValidator is a validator for the "dosage" field. It is called by the builders before save.
	prescription.DosageValidator = prescriptionDescDosage.Validators[0].(func(int) error)
	// prescriptionDescUnit is the schema descriptor for unit field.
	prescriptionDescUnit := prescriptionFields[1].Descriptor()
	// prescription.UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	prescription.UnitValidator = prescriptionDescUnit.Validators[0].(func(string) error)
	// prescriptionDescDosageFrequency is the schema descriptor for dosage_frequency field.
	prescriptionDescDosageFrequency := prescriptionFields[2].Descriptor()
	// prescription.DefaultDosageFrequency holds the default value on creation for the dosage_frequency field.
	prescription.DefaultDosageFrequency = prescriptionDescDosageFrequency.Default.(int)
	// prescription.DosageFrequencyValidator is a validator for the "dosage_frequency" field. It is called by the builders before save.
	prescription.DosageFrequencyValidator = prescriptionDescDosageFrequency.Validators[0].(func(int) error)
	// prescriptionDescStartDate is the schema descriptor for start_date field.
	prescriptionDescStartDate := prescriptionFields[3].Descriptor()
	// prescription.DefaultStartDate holds the default value on creation for the start_date field.
	prescription.DefaultStartDate = prescriptionDescStartDate.Default.(func() time.Time)
	purchaseFields := schema.Purchase{}.Fields()
	_ = purchaseFields
	// purchaseDescQuantity is the schema descriptor for quantity field.
	purchaseDescQuantity := purchaseFields[1].Descriptor()
	// purchase.QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	purchase.QuantityValidator = purchaseDescQuantity.Validators[0].(func(int) error)
	stockinglogFields := schema.StockingLog{}.Fields()
	_ = stockinglogFields
	// stockinglogDescQuantity is the schema descriptor for quantity field.
	stockinglogDescQuantity := stockinglogFields[1].Descriptor()
	// stockinglog.QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	stockinglog.QuantityValidator = stockinglogDescQuantity.Validators[0].(func(int) error)
}

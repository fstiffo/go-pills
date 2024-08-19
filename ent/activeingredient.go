// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"fstiffo/pills/ent/activeingredient"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ActiveIngredient is the model entity for the ActiveIngredient schema.
type ActiveIngredient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActiveIngredientQuery when eager-loading is set.
	Edges        ActiveIngredientEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ActiveIngredientEdges holds the relations/edges for other nodes in the graph.
type ActiveIngredientEdges struct {
	// Medicines holds the value of the medicines edge.
	Medicines []*Medicine `json:"medicines,omitempty"`
	// Prescriptions holds the value of the prescriptions edge.
	Prescriptions []*Prescription `json:"prescriptions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MedicinesOrErr returns the Medicines value or an error if the edge
// was not loaded in eager-loading.
func (e ActiveIngredientEdges) MedicinesOrErr() ([]*Medicine, error) {
	if e.loadedTypes[0] {
		return e.Medicines, nil
	}
	return nil, &NotLoadedError{edge: "medicines"}
}

// PrescriptionsOrErr returns the Prescriptions value or an error if the edge
// was not loaded in eager-loading.
func (e ActiveIngredientEdges) PrescriptionsOrErr() ([]*Prescription, error) {
	if e.loadedTypes[1] {
		return e.Prescriptions, nil
	}
	return nil, &NotLoadedError{edge: "prescriptions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ActiveIngredient) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case activeingredient.FieldID:
			values[i] = new(sql.NullInt64)
		case activeingredient.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ActiveIngredient fields.
func (ai *ActiveIngredient) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case activeingredient.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ai.ID = int(value.Int64)
		case activeingredient.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ai.Name = value.String
			}
		default:
			ai.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ActiveIngredient.
// This includes values selected through modifiers, order, etc.
func (ai *ActiveIngredient) Value(name string) (ent.Value, error) {
	return ai.selectValues.Get(name)
}

// QueryMedicines queries the "medicines" edge of the ActiveIngredient entity.
func (ai *ActiveIngredient) QueryMedicines() *MedicineQuery {
	return NewActiveIngredientClient(ai.config).QueryMedicines(ai)
}

// QueryPrescriptions queries the "prescriptions" edge of the ActiveIngredient entity.
func (ai *ActiveIngredient) QueryPrescriptions() *PrescriptionQuery {
	return NewActiveIngredientClient(ai.config).QueryPrescriptions(ai)
}

// Update returns a builder for updating this ActiveIngredient.
// Note that you need to call ActiveIngredient.Unwrap() before calling this method if this ActiveIngredient
// was returned from a transaction, and the transaction was committed or rolled back.
func (ai *ActiveIngredient) Update() *ActiveIngredientUpdateOne {
	return NewActiveIngredientClient(ai.config).UpdateOne(ai)
}

// Unwrap unwraps the ActiveIngredient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ai *ActiveIngredient) Unwrap() *ActiveIngredient {
	_tx, ok := ai.config.driver.(*txDriver)
	if !ok {
		panic("ent: ActiveIngredient is not a transactional entity")
	}
	ai.config.driver = _tx.drv
	return ai
}

// String implements the fmt.Stringer.
func (ai *ActiveIngredient) String() string {
	var builder strings.Builder
	builder.WriteString("ActiveIngredient(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ai.ID))
	builder.WriteString("name=")
	builder.WriteString(ai.Name)
	builder.WriteByte(')')
	return builder.String()
}

// ActiveIngredients is a parsable slice of ActiveIngredient.
type ActiveIngredients []*ActiveIngredient

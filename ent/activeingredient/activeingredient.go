// Code generated by ent, DO NOT EDIT.

package activeingredient

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the activeingredient type in the database.
	Label = "active_ingredient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeMedicines holds the string denoting the medicines edge name in mutations.
	EdgeMedicines = "medicines"
	// EdgePrescriptions holds the string denoting the prescriptions edge name in mutations.
	EdgePrescriptions = "prescriptions"
	// Table holds the table name of the activeingredient in the database.
	Table = "active_ingredients"
	// MedicinesTable is the table that holds the medicines relation/edge.
	MedicinesTable = "medicines"
	// MedicinesInverseTable is the table name for the Medicine entity.
	// It exists in this package in order to avoid circular dependency with the "medicine" package.
	MedicinesInverseTable = "medicines"
	// MedicinesColumn is the table column denoting the medicines relation/edge.
	MedicinesColumn = "active_ingredient_medicines"
	// PrescriptionsTable is the table that holds the prescriptions relation/edge.
	PrescriptionsTable = "prescriptions"
	// PrescriptionsInverseTable is the table name for the Prescription entity.
	// It exists in this package in order to avoid circular dependency with the "prescription" package.
	PrescriptionsInverseTable = "prescriptions"
	// PrescriptionsColumn is the table column denoting the prescriptions relation/edge.
	PrescriptionsColumn = "active_ingredient_prescriptions"
)

// Columns holds all SQL columns for activeingredient fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the ActiveIngredient queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByMedicinesCount orders the results by medicines count.
func ByMedicinesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMedicinesStep(), opts...)
	}
}

// ByMedicines orders the results by medicines terms.
func ByMedicines(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMedicinesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPrescriptionsCount orders the results by prescriptions count.
func ByPrescriptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPrescriptionsStep(), opts...)
	}
}

// ByPrescriptions orders the results by prescriptions terms.
func ByPrescriptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrescriptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMedicinesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MedicinesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MedicinesTable, MedicinesColumn),
	)
}
func newPrescriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PrescriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PrescriptionsTable, PrescriptionsColumn),
	)
}

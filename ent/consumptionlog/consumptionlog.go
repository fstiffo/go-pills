// Code generated by ent, DO NOT EDIT.

package consumptionlog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the consumptionlog type in the database.
	Label = "consumption_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldConsumedAt holds the string denoting the consumed_at field in the database.
	FieldConsumedAt = "consumed_at"
	// FieldUnits holds the string denoting the units field in the database.
	FieldUnits = "units"
	// EdgePrescription holds the string denoting the prescription edge name in mutations.
	EdgePrescription = "prescription"
	// Table holds the table name of the consumptionlog in the database.
	Table = "consumption_logs"
	// PrescriptionTable is the table that holds the prescription relation/edge.
	PrescriptionTable = "consumption_logs"
	// PrescriptionInverseTable is the table name for the Prescription entity.
	// It exists in this package in order to avoid circular dependency with the "prescription" package.
	PrescriptionInverseTable = "prescriptions"
	// PrescriptionColumn is the table column denoting the prescription relation/edge.
	PrescriptionColumn = "prescription_comsumption_logs"
)

// Columns holds all SQL columns for consumptionlog fields.
var Columns = []string{
	FieldID,
	FieldConsumedAt,
	FieldUnits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "consumption_logs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"active_ingredient_consumption_logs",
	"prescription_comsumption_logs",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultConsumedAt holds the default value on creation for the "consumed_at" field.
	DefaultConsumedAt func() time.Time
	// UnitsValidator is a validator for the "units" field. It is called by the builders before save.
	UnitsValidator func(int) error
)

// OrderOption defines the ordering options for the ConsumptionLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByConsumedAt orders the results by the consumed_at field.
func ByConsumedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConsumedAt, opts...).ToFunc()
}

// ByUnits orders the results by the units field.
func ByUnits(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnits, opts...).ToFunc()
}

// ByPrescriptionField orders the results by prescription field.
func ByPrescriptionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrescriptionStep(), sql.OrderByField(field, opts...))
	}
}
func newPrescriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PrescriptionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PrescriptionTable, PrescriptionColumn),
	)
}
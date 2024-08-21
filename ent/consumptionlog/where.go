// Code generated by ent, DO NOT EDIT.

package consumptionlog

import (
	"fstiffo/pills/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldLTE(FieldID, id))
}

// ConsumedAt applies equality check predicate on the "consumed_at" field. It's identical to ConsumedAtEQ.
func ConsumedAt(v time.Time) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldEQ(FieldConsumedAt, v))
}

// Units applies equality check predicate on the "units" field. It's identical to UnitsEQ.
func Units(v int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldEQ(FieldUnits, v))
}

// ConsumedAtEQ applies the EQ predicate on the "consumed_at" field.
func ConsumedAtEQ(v time.Time) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldEQ(FieldConsumedAt, v))
}

// ConsumedAtNEQ applies the NEQ predicate on the "consumed_at" field.
func ConsumedAtNEQ(v time.Time) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldNEQ(FieldConsumedAt, v))
}

// ConsumedAtIn applies the In predicate on the "consumed_at" field.
func ConsumedAtIn(vs ...time.Time) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldIn(FieldConsumedAt, vs...))
}

// ConsumedAtNotIn applies the NotIn predicate on the "consumed_at" field.
func ConsumedAtNotIn(vs ...time.Time) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldNotIn(FieldConsumedAt, vs...))
}

// ConsumedAtGT applies the GT predicate on the "consumed_at" field.
func ConsumedAtGT(v time.Time) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldGT(FieldConsumedAt, v))
}

// ConsumedAtGTE applies the GTE predicate on the "consumed_at" field.
func ConsumedAtGTE(v time.Time) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldGTE(FieldConsumedAt, v))
}

// ConsumedAtLT applies the LT predicate on the "consumed_at" field.
func ConsumedAtLT(v time.Time) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldLT(FieldConsumedAt, v))
}

// ConsumedAtLTE applies the LTE predicate on the "consumed_at" field.
func ConsumedAtLTE(v time.Time) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldLTE(FieldConsumedAt, v))
}

// ConsumedAtIsNil applies the IsNil predicate on the "consumed_at" field.
func ConsumedAtIsNil() predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldIsNull(FieldConsumedAt))
}

// ConsumedAtNotNil applies the NotNil predicate on the "consumed_at" field.
func ConsumedAtNotNil() predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldNotNull(FieldConsumedAt))
}

// UnitsEQ applies the EQ predicate on the "units" field.
func UnitsEQ(v int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldEQ(FieldUnits, v))
}

// UnitsNEQ applies the NEQ predicate on the "units" field.
func UnitsNEQ(v int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldNEQ(FieldUnits, v))
}

// UnitsIn applies the In predicate on the "units" field.
func UnitsIn(vs ...int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldIn(FieldUnits, vs...))
}

// UnitsNotIn applies the NotIn predicate on the "units" field.
func UnitsNotIn(vs ...int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldNotIn(FieldUnits, vs...))
}

// UnitsGT applies the GT predicate on the "units" field.
func UnitsGT(v int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldGT(FieldUnits, v))
}

// UnitsGTE applies the GTE predicate on the "units" field.
func UnitsGTE(v int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldGTE(FieldUnits, v))
}

// UnitsLT applies the LT predicate on the "units" field.
func UnitsLT(v int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldLT(FieldUnits, v))
}

// UnitsLTE applies the LTE predicate on the "units" field.
func UnitsLTE(v int) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.FieldLTE(FieldUnits, v))
}

// HasPrescription applies the HasEdge predicate on the "prescription" edge.
func HasPrescription() predicate.ConsumptionLog {
	return predicate.ConsumptionLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrescriptionTable, PrescriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrescriptionWith applies the HasEdge predicate on the "prescription" edge with a given conditions (other predicates).
func HasPrescriptionWith(preds ...predicate.Prescription) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(func(s *sql.Selector) {
		step := newPrescriptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ConsumptionLog) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ConsumptionLog) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ConsumptionLog) predicate.ConsumptionLog {
	return predicate.ConsumptionLog(sql.NotPredicates(p))
}
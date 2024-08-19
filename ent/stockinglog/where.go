// Code generated by ent, DO NOT EDIT.

package stockinglog

import (
	"fstiffo/pills/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLTE(FieldID, id))
}

// StockedAt applies equality check predicate on the "stocked_at" field. It's identical to StockedAtEQ.
func StockedAt(v time.Time) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldStockedAt, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldQuantity, v))
}

// StockedAtEQ applies the EQ predicate on the "stocked_at" field.
func StockedAtEQ(v time.Time) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldStockedAt, v))
}

// StockedAtNEQ applies the NEQ predicate on the "stocked_at" field.
func StockedAtNEQ(v time.Time) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNEQ(FieldStockedAt, v))
}

// StockedAtIn applies the In predicate on the "stocked_at" field.
func StockedAtIn(vs ...time.Time) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldIn(FieldStockedAt, vs...))
}

// StockedAtNotIn applies the NotIn predicate on the "stocked_at" field.
func StockedAtNotIn(vs ...time.Time) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNotIn(FieldStockedAt, vs...))
}

// StockedAtGT applies the GT predicate on the "stocked_at" field.
func StockedAtGT(v time.Time) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGT(FieldStockedAt, v))
}

// StockedAtGTE applies the GTE predicate on the "stocked_at" field.
func StockedAtGTE(v time.Time) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGTE(FieldStockedAt, v))
}

// StockedAtLT applies the LT predicate on the "stocked_at" field.
func StockedAtLT(v time.Time) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLT(FieldStockedAt, v))
}

// StockedAtLTE applies the LTE predicate on the "stocked_at" field.
func StockedAtLTE(v time.Time) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLTE(FieldStockedAt, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLTE(FieldQuantity, v))
}

// HasMedicine applies the HasEdge predicate on the "medicine" edge.
func HasMedicine() predicate.StockingLog {
	return predicate.StockingLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MedicineTable, MedicineColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMedicineWith applies the HasEdge predicate on the "medicine" edge with a given conditions (other predicates).
func HasMedicineWith(preds ...predicate.Medicine) predicate.StockingLog {
	return predicate.StockingLog(func(s *sql.Selector) {
		step := newMedicineStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StockingLog) predicate.StockingLog {
	return predicate.StockingLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StockingLog) predicate.StockingLog {
	return predicate.StockingLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StockingLog) predicate.StockingLog {
	return predicate.StockingLog(sql.NotPredicates(p))
}

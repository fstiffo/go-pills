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

// Boxes applies equality check predicate on the "boxes" field. It's identical to BoxesEQ.
func Boxes(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldBoxes, v))
}

// Units applies equality check predicate on the "units" field. It's identical to UnitsEQ.
func Units(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldUnits, v))
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

// StockedAtIsNil applies the IsNil predicate on the "stocked_at" field.
func StockedAtIsNil() predicate.StockingLog {
	return predicate.StockingLog(sql.FieldIsNull(FieldStockedAt))
}

// StockedAtNotNil applies the NotNil predicate on the "stocked_at" field.
func StockedAtNotNil() predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNotNull(FieldStockedAt))
}

// BoxesEQ applies the EQ predicate on the "boxes" field.
func BoxesEQ(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldBoxes, v))
}

// BoxesNEQ applies the NEQ predicate on the "boxes" field.
func BoxesNEQ(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNEQ(FieldBoxes, v))
}

// BoxesIn applies the In predicate on the "boxes" field.
func BoxesIn(vs ...int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldIn(FieldBoxes, vs...))
}

// BoxesNotIn applies the NotIn predicate on the "boxes" field.
func BoxesNotIn(vs ...int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNotIn(FieldBoxes, vs...))
}

// BoxesGT applies the GT predicate on the "boxes" field.
func BoxesGT(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGT(FieldBoxes, v))
}

// BoxesGTE applies the GTE predicate on the "boxes" field.
func BoxesGTE(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGTE(FieldBoxes, v))
}

// BoxesLT applies the LT predicate on the "boxes" field.
func BoxesLT(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLT(FieldBoxes, v))
}

// BoxesLTE applies the LTE predicate on the "boxes" field.
func BoxesLTE(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLTE(FieldBoxes, v))
}

// BoxesIsNil applies the IsNil predicate on the "boxes" field.
func BoxesIsNil() predicate.StockingLog {
	return predicate.StockingLog(sql.FieldIsNull(FieldBoxes))
}

// BoxesNotNil applies the NotNil predicate on the "boxes" field.
func BoxesNotNil() predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNotNull(FieldBoxes))
}

// UnitsEQ applies the EQ predicate on the "units" field.
func UnitsEQ(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldEQ(FieldUnits, v))
}

// UnitsNEQ applies the NEQ predicate on the "units" field.
func UnitsNEQ(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNEQ(FieldUnits, v))
}

// UnitsIn applies the In predicate on the "units" field.
func UnitsIn(vs ...int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldIn(FieldUnits, vs...))
}

// UnitsNotIn applies the NotIn predicate on the "units" field.
func UnitsNotIn(vs ...int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldNotIn(FieldUnits, vs...))
}

// UnitsGT applies the GT predicate on the "units" field.
func UnitsGT(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGT(FieldUnits, v))
}

// UnitsGTE applies the GTE predicate on the "units" field.
func UnitsGTE(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldGTE(FieldUnits, v))
}

// UnitsLT applies the LT predicate on the "units" field.
func UnitsLT(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLT(FieldUnits, v))
}

// UnitsLTE applies the LTE predicate on the "units" field.
func UnitsLTE(v int) predicate.StockingLog {
	return predicate.StockingLog(sql.FieldLTE(FieldUnits, v))
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

// HasActiveIngredient applies the HasEdge predicate on the "active_ingredient" edge.
func HasActiveIngredient() predicate.StockingLog {
	return predicate.StockingLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActiveIngredientTable, ActiveIngredientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActiveIngredientWith applies the HasEdge predicate on the "active_ingredient" edge with a given conditions (other predicates).
func HasActiveIngredientWith(preds ...predicate.ActiveIngredient) predicate.StockingLog {
	return predicate.StockingLog(func(s *sql.Selector) {
		step := newActiveIngredientStep()
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

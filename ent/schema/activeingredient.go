package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ActiveIngredient holds the schema definition for the ActiveIngredient entity.
type ActiveIngredient struct {
	ent.Schema
}

// Fields of the ActiveIngredient.
func (ActiveIngredient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty(),
		field.Int("stock").
			Default(0),
		field.Enum("unit").
			Values("mg", "g", "ml", "u").
			Optional().
			Default("mg"),
		field.Time("last_stocked_at"),
		field.Time("last_consumed_at"),
	}
}

// Edges of the ActiveIngredient.
func (ActiveIngredient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("medicines", Medicine.Type),
		edge.To("prescriptions", Prescription.Type),
		edge.To("stocking_logs", StockingLog.Type),
		edge.To("consumption_logs", ConsumptionLog.Type),
	}
}

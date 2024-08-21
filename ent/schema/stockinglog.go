package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StockingLog holds the schema definition for the StockingLog entity.
type StockingLog struct {
	ent.Schema
}

// Fields of the StockingLog.
func (StockingLog) Fields() []ent.Field {
	return []ent.Field{
		field.Time("stocked_at").
			Default(func() time.Time {
				return time.Now()
			}),
		field.Int("boxes").
			Positive().
			Default(1),
		field.Int("units").
			Positive(),
	}
}

// Edges of the StockingLog.
func (StockingLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("medicine", Medicine.Type).
			Ref("stocking_logs").
			Required().
			Unique(),
		edge.From("active_ingredient", ActiveIngredient.Type).
			Ref("stocking_logs").
			Required().
			Unique(),
	}
}

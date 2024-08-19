package schema

import (
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
		field.Time("stocked_at"),
		field.Int("quantity").Positive(),
	}
}

// Edges of the StockingLog.
func (StockingLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("medicine", Medicine.Type).
			Ref("stocking_logs").
			Unique(),
	}
}

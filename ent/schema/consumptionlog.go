package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ConsumptionLog holds the schema definition for the ConsumptionLog entity.
type ConsumptionLog struct {
	ent.Schema
}

// Fields of the ConsumptionLog.
func (ConsumptionLog) Fields() []ent.Field {
	return []ent.Field{
		field.Time("consumed_at").
			Optional().
			Default(
				func() time.Time {
					return time.Now()
				},
			),
		field.Int("units").
			Positive(),
	}
}

// Edges of the ConsumptionLog.
func (ConsumptionLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prescription", Prescription.Type).
			Ref("comsumption_logs").
			Required().
			Unique(),
		edge.From("active_ingredient", ActiveIngredient.Type).
			Ref("consumption_logs").
			Required().
			Unique(),
	}
}

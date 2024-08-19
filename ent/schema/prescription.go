package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Prescription holds the schema definition for the Prescription entity.
type Prescription struct {
	ent.Schema
}

// Fields of the Prescription.
func (Prescription) Fields() []ent.Field {
	return []ent.Field{
		field.Int("dosage").Positive(),
		field.String("unit").NotEmpty(),
		field.Int("dosage_frequency").
			Optional().
			Positive().
			Default(1), // Every 1 day
		field.Time("start_date").
			Optional().
			Default(func() time.Time {
				return time.Now()
			}),
		field.Time("end_date").Optional(),
	}
}

// Edges of the Prescription.
func (Prescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comsumption_logs", ConsumptionLog.Type),
		edge.From("active_ingredient", ActiveIngredient.Type).
			Ref("prescriptions").
			Unique(),
	}
}

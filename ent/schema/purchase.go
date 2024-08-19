package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Purchase holds the schema definition for the Purchase entity.
type Purchase struct {
	ent.Schema
}

// Fields of the Purchase.
func (Purchase) Fields() []ent.Field {
	return []ent.Field{
		field.Time("puchased_at"),
		field.Int("quantity").
			Positive(),
	}
}

// Edges of the Purchase.
func (Purchase) Edges() []ent.Edge {
	return nil
}

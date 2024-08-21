package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Medicine holds the schema definition for the Medicine entity.
type Medicine struct {
	ent.Schema
}

// Fields of the Medicine.
func (Medicine) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty(),
		field.String("mah").
			NotEmpty(),
		field.Float("dosage").
			Positive(),
		field.String("atc").
			Unique().
			NotEmpty().
			MinLen(7).
			Match(regexp.MustCompile("[A-Z][0-9][0-9][A-Z][A-Z][0-9][0-9].*")),
		field.String("package"),
		field.String("form"),
		field.Int("box_size").
			Positive(),
	}
}

// Edges of the Medicine.
func (Medicine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stocking_logs", StockingLog.Type),
		edge.From("active_ingredient", ActiveIngredient.Type).
			Ref("medicines").
			Required().
			Unique(),
	}
}

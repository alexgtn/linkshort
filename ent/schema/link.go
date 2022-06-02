package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.String("long_uri").NotEmpty().Unique(),
		field.Int("accessed_times").Default(0).NonNegative(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return nil
}

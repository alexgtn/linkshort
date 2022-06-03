package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/alexgtn/go-linkshort/domain/link"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.String("short_path").NotEmpty().Unique().MaxLen(link.MaxLen),
		field.String("long_uri").NotEmpty().Unique().Immutable().MaxLen(link.MaxLen),
		field.Int("accessed_times").Default(0).NonNegative(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return nil
}

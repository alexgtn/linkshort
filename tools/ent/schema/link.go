package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/alexgtn/go-linkshort/internal/domain/link"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		// By default it creates a unique ID int
		field.String("short_path").Optional().Unique().MaxLen(link.MaxLen),
		field.String("long_uri").NotEmpty().Unique().Immutable().MaxLen(link.MaxLen),
		field.Int("accessed_times").Default(0).NonNegative(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return nil
}

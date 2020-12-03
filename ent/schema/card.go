package schema

import (
	"database/sql"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Amount is a custom Go type that's convertible to the basic float64 type.
type Amount float64

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().
			// A ValueScanner type.
			GoType(&sql.NullString{}),
		field.String("number").NotEmpty(),
		field.Time("expired").Immutable(),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("card").Unique().
			Required(),
	}
}

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amout").SchemaType(map[string]string{
			dialect.MySQL:    "decimal(6,2)", // Override MySQL
			dialect.Postgres: "numeric",      // Override Postgres.
		}),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return nil
}

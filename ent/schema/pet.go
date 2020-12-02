package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(25).NotEmpty().Unique().Immutable(),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return nil
}
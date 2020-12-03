package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return nil
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("next", Node.Type).Unique().From("prev").Unique(),
	}
}

package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Positive(),
		field.String("name").Default("a8m"),
		field.String("username").Unique(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
		edge.To("groups", Group.Type),
		edge.To("friends", User.Type),
	}
}

func (User) Index() []ent.Index {
	return []ent.Index{
		index.Fields("age", "name").Unique(),
	}
}

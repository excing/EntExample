package schema

import (
	"net/url"
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Positive().Optional(),
		field.Float("rank").Optional(),
		field.Bool("active").Default(false),
		// Name string `json:"name,omitempty" gqlgen:"gql_name"`
		field.String("name").StorageKey("old_name").StructTag(`gqlgen:"gql_name"`).Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.JSON("url", &url.URL{}).Optional(),
		field.JSON("strings", []string{}).Optional(),
		field.Enum("state").Values("on", "off").Optional(),
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		// 	Nickname *string `json:"nickname,omitempty"
		field.String("nickname").Optional().Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
		edge.From("groups", Group.Type).Ref("users"),
		edge.To("friends", User.Type),
		edge.To("pets", Pet.Type),
		edge.To("card", Card.Type).Unique(),
		edge.To("spouse", User.Type).Unique(),
	}
}

// Index user index
func (User) Index() []ent.Index {
	return []ent.Index{
		index.Fields("age", "name").Unique(),
	}
}

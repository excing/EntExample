package schema

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"old,omitempty"`),
		field.String("name").
			// Regexp validation for group name.
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("group name must begin with uppercase")
				}
				return nil
			}),
		field.String("nickname").Validate(MaxRuneCount(20)).Optional(),
		field.Int("count").Positive().Optional(),
		field.Int("code").Negative().Optional(),
		field.Int("index").NonNegative().Optional(),
		field.Int("min").Min(3).Optional(),
		field.Int("max").Max(9).Optional(),
		field.Int("range").Range(3, 9).Optional(),
		field.String("note").MinLen(10).Optional(),
		field.String("log").MaxLen(200).Optional(),
		field.String("username").NotEmpty().Optional(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}

// MaxRuneCount validates the rune length of a string by using the unicode/utf8 package
func MaxRuneCount(maxLen int) func(s string) error {
	return func(s string) error {
		if utf8.RuneCountInString(s) > maxLen {
			return errors.New("value is more than the max length")
		}
		return nil
	}
}

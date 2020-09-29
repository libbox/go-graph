package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Optional(),
		field.String("name").
			Default("unknown"),
	}
}

type Annotation struct {
	Category string
}

func (a *Annotation) Name() string {
	return "EntGQL"
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("following", User.Type).From("followers"),
		edge.To("requesting", User.Type).From("requesters"),
		edge.To("friend", User.Type),
		//edge.To("likes", Post.Type).From("users"),
	}
}

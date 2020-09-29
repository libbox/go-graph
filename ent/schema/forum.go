package schema

import "github.com/facebook/ent"

// Forum holds the schema definition for the Forum entity.
type Forum struct {
	ent.Schema
}

// Fields of the Forum.
func (Forum) Fields() []ent.Field {
	return nil
}

// Edges of the Forum.
func (Forum) Edges() []ent.Edge {
	return nil
}

package schema

import "github.com/facebook/ent"

// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return nil
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return nil
}

package schema

import "github.com/facebook/ent"

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return nil
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return nil
}

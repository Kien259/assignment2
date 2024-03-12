package schema

import "entgo.io/ent"

// Earthquake holds the schema definition for the Earthquake entity.
type Earthquake struct {
	ent.Schema
}

// Fields of the Earthquake.
func (Earthquake) Fields() []ent.Field {
	return nil
}

// Edges of the Earthquake.
func (Earthquake) Edges() []ent.Edge {
	return nil
}

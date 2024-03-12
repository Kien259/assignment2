package schema

import "entgo.io/ent"

// Coordinate holds the schema definition for the Coordinate entity.
type Coordinate struct {
	ent.Schema
}

// Fields of the Coordinate.
func (Coordinate) Fields() []ent.Field {
	return nil
}

// Edges of the Coordinate.
func (Coordinate) Edges() []ent.Edge {
	return nil
}

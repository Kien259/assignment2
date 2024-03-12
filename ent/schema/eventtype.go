package schema

import "entgo.io/ent"

// EventType holds the schema definition for the EventType entity.
type EventType struct {
	ent.Schema
}

// Fields of the EventType.
func (EventType) Fields() []ent.Field {
	return nil
}

// Edges of the EventType.
func (EventType) Edges() []ent.Edge {
	return nil
}

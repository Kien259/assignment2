package schema

import "entgo.io/ent"

// AssociatedEvent holds the schema definition for the AssociatedEvent entity.
type AssociatedEvent struct {
	ent.Schema
}

// Fields of the AssociatedEvent.
func (AssociatedEvent) Fields() []ent.Field {
	return nil
}

// Edges of the AssociatedEvent.
func (AssociatedEvent) Edges() []ent.Edge {
	return nil
}

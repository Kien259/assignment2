package schema

import "entgo.io/ent"

// FeatureType holds the schema definition for the FeatureType entity.
type FeatureType struct {
	ent.Schema
}

// Fields of the FeatureType.
func (FeatureType) Fields() []ent.Field {
	return nil
}

// Edges of the FeatureType.
func (FeatureType) Edges() []ent.Edge {
	return nil
}

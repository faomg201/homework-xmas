package schema

import "github.com/facebookincubator/ent"

// Research holds the schema definition for the Research entity.
type Research struct {
	ent.Schema
}

// Fields of the Research.
func (Research) Fields() []ent.Field {
	return nil
}

// Edges of the Research.
func (Research) Edges() []ent.Edge {
	return nil
}

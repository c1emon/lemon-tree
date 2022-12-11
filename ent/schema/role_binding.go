package schema

import (
	"entgo.io/ent"
)

// RoleBinding holds the schema definition for the RoleBinding entity.
type RoleBinding struct {
	ent.Schema
}

// Fields of the RoleBinding.
func (RoleBinding) Fields() []ent.Field {
	return nil
}

// Edges of the RoleBinding.
func (RoleBinding) Edges() []ent.Edge {
	return nil
}

func (RoleBinding) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

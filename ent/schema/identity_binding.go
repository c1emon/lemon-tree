package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IdentityBinding holds the schema definition for the IdentityBinding entity.
type IdentityBinding struct {
	ent.Schema
}

// Fields of the IdentityBinding.
func (IdentityBinding) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the IdentityBinding.
func (IdentityBinding) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("identity_bindings").
			Unique(),
		edge.From("identity_provider", IdentityProvider.Type).
			Ref("identity_bindings").
			Unique(),
		edge.From("user", User.Type).
			Ref("identity_bindings").
			Unique(),
	}
}

func (IdentityBinding) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

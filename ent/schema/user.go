package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("username").Unique(),
		field.String("password").Sensitive(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.String("openid"),
		field.Int("age").Optional().Max(120).Min(0),
		field.Enum("gender").Values("male", "female", "unknown").Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("users").
			Unique(),
		edge.From("department", Department.Type).
			Ref("users").
			Unique(),
		edge.To("identity_bindings", IdentityBinding.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

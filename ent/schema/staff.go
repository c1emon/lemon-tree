package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Staff holds the schema definition for the Staff entity.
type Staff struct {
	ent.Schema
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("username"),
		field.String("password"),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.String("openid"),
		field.Int("age").Optional(),
		field.Enum("gender").Values("male", "female", "unknown").Default("unknown"),
	}
}

// Edges of the Staff.
func (Staff) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("staffs").
			Unique(),
		edge.From("department", Department.Type).
			Ref("staffs").
			Unique(),
		edge.To("identity_bindings", IdentityBinding.Type),
	}
}

func (Staff) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

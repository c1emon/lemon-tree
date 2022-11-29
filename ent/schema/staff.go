package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Staff holds the schema definition for the Staff entity.
type Staff struct {
	ent.Schema
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age").Optional(),
		field.Enum("gender").Optional(),
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
	}
}

func (Staff) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

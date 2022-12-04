package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// OauthClient holds the schema definition for the OauthClient entity.
type OauthClient struct {
	ent.Schema
}

// Fields of the OauthClient.
func (OauthClient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("client_id").Unique().Immutable(),
		field.String("client_secret"),
		field.Strings("grant_type").Default("password"),
		field.Strings("redirect_urls"),
		field.Bool("enabled").Default(false),
		field.Bool("internal").Default(false).Immutable(),
	}
}

// Edges of the OauthClient.
func (OauthClient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("oauth_clients").
			Unique(),
		edge.From("creator", Staff.Type).Immutable(),
	}
}

func (OauthClient) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

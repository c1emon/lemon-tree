package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.Enum("grant_type").Values("password", "token", "authorization_code"),
		field.Strings("redirect_urls"),
		field.Bool("enabled").Default(false),
		field.Bool("internal").Default(false).Immutable(),
		field.String("creator").Optional(),
	}
}

// Edges of the OauthClient.
func (OauthClient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("oauth_clients").
			Unique(),
	}
}

func (OauthClient) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// OauthIdentityProvider holds the schema definition for the OauthIdentityProvider entity.
type OauthIdentityProvider struct {
	ent.Schema
}

// Fields of the OauthIdentityProvider.
func (OauthIdentityProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Strings("client_id"),
		field.Strings("client_secret"),
		field.Strings("redirect_urls"),
		field.Strings("type"),
	}
}

// Edges of the OauthIdentityProvider.
func (OauthIdentityProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("oauth_identity_providers").
			Unique(),
	}
}

func (OauthIdentityProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

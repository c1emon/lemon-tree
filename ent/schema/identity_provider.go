package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/c1emon/lemontree/goauth"
)

// IdentityProvider holds the schema definition for the IdentityProvider entity.
type IdentityProvider struct {
	ent.Schema
}

// Fields of the IdentityProvider.
func (IdentityProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("type").Values("oauth", "ldap", "saml"),
		field.JSON("oauth_config", &goauth.OauthBasicConfig{}),
	}
}

// Edges of the IdentityProvider.
func (IdentityProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("identity_providers").
			Unique(),
		edge.To("identity_bindings", IdentityBinding.Type),
	}
}

func (IdentityProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

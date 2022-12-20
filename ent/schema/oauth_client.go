package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OAuthClient holds the schema definition for the OAuthClient entity.
type OAuthClient struct {
	ent.Schema
}

// Fields of the OAuthClient.
func (OAuthClient) Fields() []ent.Field {

	//ClientID       string   `json:"client_id"`
	//Secret         []byte   `json:"client_secret,omitempty"`
	//RotatedSecrets [][]byte `json:"rotated_secrets,omitempty"`
	//RedirectURIs   []string `json:"redirect_uris"`
	//GrantTypes     []string `json:"grant_types"`
	//ResponseTypes  []string `json:"response_types"`
	//Scopes         []string `json:"scopes"`
	//Audience       []string `json:"audience"`
	//Public         bool     `json:"public"`
	//
	//ResponseModes  []ResponseModeType    `json:"response_modes"`
	//TokenLifespans *ClientLifespanConfig `json:"token_lifespans"`

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

// Edges of the OAuthClient.
func (OAuthClient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("oauth_clients").
			Unique(),
	}
}

func (OAuthClient) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

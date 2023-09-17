package oidc

import (
	"time"

	"github.com/c1emon/lemontree/model"
	"github.com/zitadel/oidc/v2/pkg/oidc"
	"github.com/zitadel/oidc/v2/pkg/op"
)

var (
	_ op.Client = &Client{}
)

type Client struct {
	model.BaseFields
	secret                         string
	redirectURIs                   []string
	applicationType                op.ApplicationType
	authMethod                     oidc.AuthMethod
	loginURL                       func(string) string
	responseTypes                  []oidc.ResponseType
	grantTypes                     []oidc.GrantType
	accessTokenType                op.AccessTokenType
	devMode                        bool
	idTokenUserinfoClaimsAssertion bool
	clockSkew                      time.Duration
	postLogoutRedirectURIGlobs     []string
	redirectURIGlobs               []string
}

// AccessTokenType implements op.Client.
func (*Client) AccessTokenType() op.AccessTokenType {
	panic("unimplemented")
}

// ApplicationType implements op.Client.
func (*Client) ApplicationType() op.ApplicationType {
	panic("unimplemented")
}

// AuthMethod implements op.Client.
func (*Client) AuthMethod() oidc.AuthMethod {
	panic("unimplemented")
}

// ClockSkew implements op.Client.
func (*Client) ClockSkew() time.Duration {
	panic("unimplemented")
}

// DevMode implements op.Client.
func (*Client) DevMode() bool {
	panic("unimplemented")
}

// GetID implements op.Client.
func (*Client) GetID() string {
	panic("unimplemented")
}

// GrantTypes implements op.Client.
func (*Client) GrantTypes() []oidc.GrantType {
	panic("unimplemented")
}

// IDTokenLifetime implements op.Client.
func (*Client) IDTokenLifetime() time.Duration {
	panic("unimplemented")
}

// IDTokenUserinfoClaimsAssertion implements op.Client.
func (*Client) IDTokenUserinfoClaimsAssertion() bool {
	panic("unimplemented")
}

// IsScopeAllowed implements op.Client.
func (*Client) IsScopeAllowed(scope string) bool {
	panic("unimplemented")
}

// LoginURL implements op.Client.
func (*Client) LoginURL(string) string {
	panic("unimplemented")
}

// PostLogoutRedirectURIs implements op.Client.
func (*Client) PostLogoutRedirectURIs() []string {
	panic("unimplemented")
}

// RedirectURIs implements op.Client.
func (*Client) RedirectURIs() []string {
	panic("unimplemented")
}

// ResponseTypes implements op.Client.
func (*Client) ResponseTypes() []oidc.ResponseType {
	panic("unimplemented")
}

// RestrictAdditionalAccessTokenScopes implements op.Client.
func (*Client) RestrictAdditionalAccessTokenScopes() func(scopes []string) []string {
	panic("unimplemented")
}

// RestrictAdditionalIdTokenScopes implements op.Client.
func (*Client) RestrictAdditionalIdTokenScopes() func(scopes []string) []string {
	panic("unimplemented")
}

type warpToRedirectGlobs struct {
	*Client
}

// RedirectURIGlobs provide wildcarding for additional valid redirects
func (c warpToRedirectGlobs) RedirectURIGlobs() []string {
	return c.redirectURIGlobs
}

// PostLogoutRedirectURIGlobs provide extra wildcarding for additional valid redirects
func (c warpToRedirectGlobs) PostLogoutRedirectURIGlobs() []string {
	return c.postLogoutRedirectURIGlobs
}

// RedirectGlobsClient wraps the client in a op.HasRedirectGlobs
// only if DevMode is enabled.
func RedirectGlobsClient(client *Client) op.Client {
	if client.devMode {
		return warpToRedirectGlobs{client}
	}
	return client
}

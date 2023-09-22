package oidc

import (
	"time"

	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/zitadel/oidc/v2/pkg/oidc"
	"github.com/zitadel/oidc/v2/pkg/op"
)

var (
	_ op.Client = &Client{}
)

type Client struct {
	gormx.BaseFields
	Name                           string   `json:"name" gorm:"column:name;type:varchar(256);not null"`
	Secret                         string   `json:"secret" gorm:"column:secret;type:varchar(256);not null"`
	redirectURIs                   []string // 回调 回到原请求app
	applicationType                op.ApplicationType
	authMethod                     oidc.AuthMethod
	loginURL                       func(string) string // 登陆页面地址
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
func (c *Client) AccessTokenType() op.AccessTokenType {
	return c.accessTokenType
}

// ApplicationType implements op.Client.
func (c *Client) ApplicationType() op.ApplicationType {
	return c.applicationType
}

// AuthMethod implements op.Client.
func (c *Client) AuthMethod() oidc.AuthMethod {
	return c.authMethod
}

// ClockSkew implements op.Client.
func (c *Client) ClockSkew() time.Duration {
	return c.clockSkew
}

// DevMode implements op.Client.
func (c *Client) DevMode() bool {
	return c.devMode
}

// GetID implements op.Client.
func (c *Client) GetID() string {
	return c.Id
}

// GrantTypes implements op.Client.
func (c *Client) GrantTypes() []oidc.GrantType {
	return c.grantTypes
}

// IDTokenLifetime implements op.Client.
func (c *Client) IDTokenLifetime() time.Duration {
	return 1 * time.Hour
}

// IDTokenUserinfoClaimsAssertion implements op.Client.
func (c *Client) IDTokenUserinfoClaimsAssertion() bool {
	return c.idTokenUserinfoClaimsAssertion
}

// IsScopeAllowed implements op.Client.
func (c *Client) IsScopeAllowed(scope string) bool {
	return scope == "custom_scope"
}

// 重定向到登陆页面
// LoginURL implements op.Client.
func (c *Client) LoginURL(id string) string {
	return c.loginURL(id)
}

// PostLogoutRedirectURIs implements op.Client.
func (c *Client) PostLogoutRedirectURIs() []string {
	return []string{}
}

// RedirectURIs implements op.Client.
func (c *Client) RedirectURIs() []string {
	return c.redirectURIs
}

// ResponseTypes implements op.Client.
func (c *Client) ResponseTypes() []oidc.ResponseType {
	return c.responseTypes
}

// RestrictAdditionalAccessTokenScopes implements op.Client.
func (c *Client) RestrictAdditionalAccessTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

// RestrictAdditionalIdTokenScopes implements op.Client.
func (c *Client) RestrictAdditionalIdTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
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

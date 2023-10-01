package client

import (
	"fmt"
	"sort"
	"time"

	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/zitadel/oidc/v2/pkg/oidc"
	"github.com/zitadel/oidc/v2/pkg/op"
	"golang.org/x/exp/slices"
)

type OidcClientRepository interface {
	gormx.BaseRepository[Client]
	FindByOidAndName(string, string) *Client
}

type Client struct {
	gormx.BaseFields
	Name                           string              `gorm:"column:name;type:varchar(256);uniqueIndex:udx_org_name;not null"`
	OrganizationId                 string              `gorm:"column:oid;type:varchar(256);uniqueIndex:udx_org_name;not null"`
	Secret                         string              `gorm:"column:secret;type:varchar(256);not null"`
	RedirectURIs                   []string            `gorm:"column:redirect_uri;type:text;serializer:json;not null"` // 回调 回到原请求app
	ApplicationType                op.ApplicationType  `gorm:"column:application_type;type:varchar(256);not null"`
	AuthMethod                     oidc.AuthMethod     `gorm:"column:auth_method;type:varchar(256);not null"`
	LoginURL                       string              `gorm:"column:login_uri;type:varchar(256);not null"` // 登陆页面地址
	ResponseTypes                  []oidc.ResponseType `gorm:"column:response_types;type:text;serializer:json;not null"`
	GrantTypes                     []oidc.GrantType    `gorm:"column:grant_types;serializer:json;not null"`
	AccessTokenType                op.AccessTokenType  `gorm:"column:access_token_type;type:varchar(256);not null"`
	DevMode                        bool                `gorm:"column:dev_mode;type:bool;default:false;not null"`
	IdTokenUserinfoClaimsAssertion bool                `gorm:"column:id_token_userinfo_claims_assertion;type:bool;default:false;not null"`
	ClockSkew                      time.Duration       `gorm:"column:clock_skew;type:int;not null"`
	PostLogoutRedirectURIGlobs     []string            `gorm:"column:post_logout_redirect_uri_globs;type:text;serializer:json"`
	RedirectURIGlobs               []string            `gorm:"column:redirect_uri_globs;type:text;serializer:json"`
	AllowdScopes                   []string            `gorm:"column:allowd_scopes;type:text;serializer:json"`
}

func (Client) TableName() string {
	return "oidc_clients"
}

func (c *Client) WarpToOidcClient() op.Client {
	return &OidcClient{entity: c}
}

var _ op.Client = &OidcClient{}

type OidcClient struct {
	entity *Client
}

// AccessTokenType implements op.Client.
func (c *OidcClient) AccessTokenType() op.AccessTokenType {
	return c.entity.AccessTokenType
}

// ApplicationType implements op.Client.
func (c *OidcClient) ApplicationType() op.ApplicationType {
	return c.entity.ApplicationType
}

// AuthMethod implements op.Client.
func (c *OidcClient) AuthMethod() oidc.AuthMethod {
	return c.entity.AuthMethod
}

// ClockSkew implements op.Client.
func (c *OidcClient) ClockSkew() time.Duration {
	return c.entity.ClockSkew
}

// DevMode implements op.Client.
func (c *OidcClient) DevMode() bool {
	return c.entity.DevMode
}

// GetID implements op.Client.
func (c *OidcClient) GetID() string {
	return c.entity.Id
}

// GrantTypes implements op.Client.
func (c *OidcClient) GrantTypes() []oidc.GrantType {
	return c.entity.GrantTypes
}

// IDTokenLifetime implements op.Client.
func (c *OidcClient) IDTokenLifetime() time.Duration {
	return 1 * time.Hour
}

// IDTokenUserinfoClaimsAssertion implements op.Client.
func (c *OidcClient) IDTokenUserinfoClaimsAssertion() bool {
	return c.entity.IdTokenUserinfoClaimsAssertion
}

// IsScopeAllowed implements op.Client.
func (c *OidcClient) IsScopeAllowed(scope string) bool {
	slices.Sort(c.entity.AllowdScopes)
	idx := sort.SearchStrings(c.entity.AllowdScopes, scope)
	return idx < len(c.entity.AllowdScopes) && c.entity.AllowdScopes[idx] == scope
}

// 重定向到登陆页面
// LoginURL implements op.Client.
func (c *OidcClient) LoginURL(id string) string {
	// TODO: precess with net/url
	return fmt.Sprintf("%s?id=%s", c.entity.LoginURL, id)
}

// PostLogoutRedirectURIs implements op.Client.
func (c *OidcClient) PostLogoutRedirectURIs() []string {
	if c.entity.PostLogoutRedirectURIGlobs != nil && len(c.entity.PostLogoutRedirectURIGlobs) > 0 {
		return c.entity.PostLogoutRedirectURIGlobs
	}
	return []string{}
}

// RedirectURIs implements op.Client.
func (c *OidcClient) RedirectURIs() []string {
	return c.entity.RedirectURIs
}

// ResponseTypes implements op.Client.
func (c *OidcClient) ResponseTypes() []oidc.ResponseType {
	return c.entity.ResponseTypes
}

// RestrictAdditionalAccessTokenScopes implements op.Client.
func (c *OidcClient) RestrictAdditionalAccessTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

// RestrictAdditionalIdTokenScopes implements op.Client.
func (c *OidcClient) RestrictAdditionalIdTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

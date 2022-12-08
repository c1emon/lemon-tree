package dingtalk

import (
	"github.com/c1emon/lemontree/goauth"
	"github.com/google/uuid"
	"net/url"
	"strings"
)

var AuthZBaseUrl = "https://login.dingtalk.com/oauth2/auth"

type AuthZBuilder struct {
	config *goauth.AuthConfig
	state  string
}

func NewAuthZBuilder() *AuthZBuilder {
	id, _ := uuid.NewUUID()
	return &AuthZBuilder{
		config: &goauth.AuthConfig{},
		state:  id.String(),
	}
}

func (b *AuthZBuilder) WithConfig(c *goauth.AuthConfig) *AuthZBuilder {
	b.config = c
	return b
}

func (b *AuthZBuilder) SetOrgType(t string) *AuthZBuilder {
	b.config.AuthZParams["org_type"] = t
	return b
}

func (b *AuthZBuilder) SetCropIds(ids []string) *AuthZBuilder {
	b.config.AuthZParams["corpId"] = ids
	return b
}

func (b *AuthZBuilder) IsExclusiveLogin(i bool) *AuthZBuilder {
	b.config.AuthZParams["exclusiveLogin"] = i

	return b
}

func (b *AuthZBuilder) SetExclusiveCorpId(id string) *AuthZBuilder {
	b.config.AuthZParams["exclusiveCorpId"] = id
	return b
}

func (b *AuthZBuilder) SetState(state string) *AuthZBuilder {
	b.state = state
	return b
}

func (b *AuthZBuilder) Build() string {
	bc := b.config

	authZUri, _ := url.Parse(AuthZBaseUrl)
	query := authZUri.Query()

	query.Set("redirect_uri", bc.RedirectUri)
	query.Set("response_type", "code")
	query.Set("client_id", bc.ClientId)
	query.Set("client_secret", bc.ClientSecret)
	query.Set("state", b.state)
	query.Set("prompt", "consent")

	ap := bc.AuthZParams
	if bc.Scope != nil && len(bc.Scope) > 0 {
		scope := strings.Join(bc.Scope, " ")
		query.Set("scope", scope)
		if strings.Contains(scope, "corpid") {
			if v, ok := ap["org_type"]; ok {
				query.Set("org_type", v.(string))
			}
			if v, ok := ap["corpId"]; ok {
				query.Set("corpId", strings.Join(v.([]string), " "))
			}
		}
	}

	if v, ok := ap["exclusiveLogin"]; ok && v.(bool) {
		query.Set("exclusiveLogin", "true")
		if v, ok := ap["exclusiveCorpId"]; ok {
			query.Set("exclusiveCorpId", v.(string))
		}
	}

	authZUri.RawQuery = query.Encode()

	return authZUri.String()
}

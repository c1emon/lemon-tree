package goauth

type AuthConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	StateCache   StateCache
}

type AuthConfigBuilder struct {
	config *AuthConfig
}

func NewAuthConfigBuilder() *AuthConfigBuilder {
	return &AuthConfigBuilder{
		config: &AuthConfig{},
	}
}

func (acb *AuthConfigBuilder) ClientId(id string) *AuthConfigBuilder {
	acb.config.ClientId = id
	return acb
}

func (acb *AuthConfigBuilder) SetClientSecret(secret string) *AuthConfigBuilder {
	acb.config.ClientSecret = secret
	return acb
}

func (acb *AuthConfigBuilder) SetRedirectUri(uri string) *AuthConfigBuilder {
	acb.config.RedirectUri = uri
	return acb
}

func (acb *AuthConfigBuilder) SetStateCache(cache StateCache) *AuthConfigBuilder {
	acb.config.StateCache = cache
	return acb
}

func (acb *AuthConfigBuilder) Build() *AuthConfig {

	return acb.config
}

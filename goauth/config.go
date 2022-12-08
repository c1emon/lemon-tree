package goauth

type AuthConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	Scope        []string
	CustomValues map[string]any
	AuthZParams  map[string]any
}

type AuthConfigBuilder struct {
	config *AuthConfig
}

func NewAuthConfigBuilder() *AuthConfigBuilder {
	return &AuthConfigBuilder{
		config: &AuthConfig{
			CustomValues: make(map[string]any),
		},
	}
}

func (acb *AuthConfigBuilder) SetClientId(id string) *AuthConfigBuilder {
	acb.config.ClientId = id
	return acb
}

func (acb *AuthConfigBuilder) SetScope(s []string) *AuthConfigBuilder {
	acb.config.Scope = s
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

func (acb *AuthConfigBuilder) AddValue(k string, v any) *AuthConfigBuilder {
	acb.config.CustomValues[k] = v
	return acb
}

func (acb *AuthConfigBuilder) Build() *AuthConfig {

	return acb.config
}

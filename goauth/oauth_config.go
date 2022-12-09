package goauth

type OauthBasicConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	Scope        []string
	CustomParams map[string]any
	AuthZParams  map[string]any
}

type OauthBasicConfigBuilder struct {
	config *OauthBasicConfig
}

func NewOauthBasicConfigBuilder() *OauthBasicConfigBuilder {
	return &OauthBasicConfigBuilder{
		config: &OauthBasicConfig{
			CustomParams: make(map[string]any),
			AuthZParams:  make(map[string]any),
		},
	}
}

func (acb *OauthBasicConfigBuilder) SetClientId(id string) *OauthBasicConfigBuilder {
	acb.config.ClientId = id
	return acb
}

func (acb *OauthBasicConfigBuilder) SetScope(s []string) *OauthBasicConfigBuilder {
	acb.config.Scope = s
	return acb
}

func (acb *OauthBasicConfigBuilder) SetClientSecret(secret string) *OauthBasicConfigBuilder {
	acb.config.ClientSecret = secret
	return acb
}

func (acb *OauthBasicConfigBuilder) SetRedirectUri(uri string) *OauthBasicConfigBuilder {
	acb.config.RedirectUri = uri
	return acb
}

func (acb *OauthBasicConfigBuilder) AddValue(k string, v any) *OauthBasicConfigBuilder {
	acb.config.CustomParams[k] = v
	return acb
}

func (acb *OauthBasicConfigBuilder) Build() *OauthBasicConfig {

	return acb.config
}

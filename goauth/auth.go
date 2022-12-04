package goauth

type AbstractAuthHandler interface {
	Authorize(state string) string
	AccessToken(code string) string
	UserInfo()

	Login() string
	Revoke() string
	Refresh() string
}

type AuthRequest struct {
	handler *AbstractAuthHandler
}

type AuthRequestBuilder struct {
	request *AuthRequest
}

func NewAuthRequestBuilder() *AuthRequestBuilder {
	return &AuthRequestBuilder{
		request: &AuthRequest{},
	}
}

func (arb *AuthRequestBuilder) Source(source string) *AuthRequestBuilder {
	switch source {
	case "DingTalk":
		arb.request.handler = nil
	default:
		arb.request.handler = nil
	}
	return arb
}

func (arb *AuthRequestBuilder) Config(config *AuthConfig) *AuthRequestBuilder {

	return arb
}

func (arb *AuthRequestBuilder) Build() *AuthRequest {

	return arb.request
}

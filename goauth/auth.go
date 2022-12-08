package goauth

type AuthHandler interface {
	GetAuthorizeUri(string) string
	ProcessCallback(string, map[string]string)
	GetAccessToken() string
	GetUserInfo()

	Login() string
	Revoke() string
	Refresh() string

	SetConfig(*AuthConfig)
	GetConfig() *AuthConfig
}

func Source(source string) *AuthHandler {
	var h *AuthHandler
	switch source {
	case "DingTalk":
		h = nil
	default:
		h = nil
	}
	return h
}

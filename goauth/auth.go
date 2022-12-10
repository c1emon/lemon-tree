package goauth

type AuthHandler interface {
	GetAccessToken(string) (string, error)
	GetIdentity(string) (*Identity, error)
}

type Identity struct {
	UserName    string
	DisplayName string
	Ids         []string
	UnionIds    []string
	OpenIds     []string
	Email       string
	Phone       string
	AvatarUrl   string
	Others      map[string]any
}

//func GetHandler(source string, config *OauthBasicConfig) *AuthHandler {
//	var h AuthHandler
//	switch source {
//	case "DingTalk":
//		h = dingtalk.NewDingTalkOauthHandler(config)
//	case "DingTalkMiniApp":
//		h = dingtalk.NewDingTalkMiniAppHandler(config)
//	default:
//		h = nil
//	}
//	return &h
//}

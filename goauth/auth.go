package goauth

import "github.com/c1emon/lemontree/goauth/dingtalk"

type AuthHandler interface {
	GetAccessToken(string) string
	GetUserInfo(string) string
}

func GetHandler(source string, config *OauthBasicConfig) *AuthHandler {
	var h AuthHandler
	switch source {
	case "DingTalk":
		h = dingtalk.NewDingTalkOauthHandler(config)
	default:
		h = nil
	}
	return &h
}

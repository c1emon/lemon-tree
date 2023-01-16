package test

import (
	"fmt"
	"github.com/c1emon/lemontree/goauth"
	"github.com/c1emon/lemontree/goauth/dingtalk"
	"testing"
)

func Test_GetUri(t *testing.T) {
	config := goauth.NewOauthBasicConfigBuilder().
		SetClientId("").
		SetClientSecret("").
		SetScope([]string{"openid"}).
		SetRedirectUri("https://baidu.com/oauth/ididid").Build()
	uri := dingtalk.NewOauthAuthZUriBuilder().WithConfig(config).Build()
	fmt.Printf("\n%s\n", uri)
}

func Test_GetIdentity(t *testing.T) {
	config := goauth.NewOauthBasicConfigBuilder().
		SetClientId("").
		SetClientSecret("").
		SetScope([]string{"openid"}).
		SetRedirectUri("https://baidu.com/oauth/ididid").Build()

	h := dingtalk.NewDingTalkOauthHandler(config)

	var code = ""
	token, err := h.GetAccessToken(code)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	i, err := h.GetIdentity(token)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	fmt.Printf("Identity: %+v", i)
}

func Test_GetIdentityMiniApp(t *testing.T) {
	config := goauth.NewOauthBasicConfigBuilder().
		SetClientId("").
		SetClientSecret("").
		Build()

	h := dingtalk.NewDingTalkMiniAppHandler(config)

	var code = ""
	token, err := h.GetAccessToken(code)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	i, err := h.GetIdentity(token)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	fmt.Printf("Identity: %+v", i)

}

package test

import (
	"fmt"
	"github.com/c1emon/lemontree/goauth"
	"github.com/c1emon/lemontree/goauth/dingtalk"
	"testing"
)

func Test_GetUri(t *testing.T) {
	config := goauth.NewOauthBasicConfigBuilder().
		SetClientId("dingdjymrdzdxa191wcz").
		SetClientSecret("-maRQbiGZ4KM7BllyOpmYfggNl5fzrof9XWV7jZm_0ZN-sbUbl_6V-jeDdSaeJ28").
		SetScope([]string{"openid"}).
		SetRedirectUri("https://baidu.com/oauth/ididid").Build()
	uri := dingtalk.NewOauthAuthZUriBuilder().WithConfig(config).Build()
	fmt.Printf("\n%s\n", uri)
}

func Test_GetIdentity(t *testing.T) {
	config := goauth.NewOauthBasicConfigBuilder().
		SetClientId("dingdjymrdzdxa191wcz").
		SetClientSecret("-maRQbiGZ4KM7BllyOpmYfggNl5fzrof9XWV7jZm_0ZN-sbUbl_6V-jeDdSaeJ28").
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
		SetClientId("dingdjymrdzdxa191wcz").
		SetClientSecret("-maRQbiGZ4KM7BllyOpmYfggNl5fzrof9XWV7jZm_0ZN-sbUbl_6V-jeDdSaeJ28").
		Build()

	h := dingtalk.NewDingTalkMiniAppHandler(config)

	var code = "e8248054a40f3cf6a0ff6035f762d9b6"
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

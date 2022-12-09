package test

import (
	"fmt"
	"github.com/c1emon/lemontree/goauth"
	"github.com/c1emon/lemontree/goauth/dingtalk"
	"testing"
)

func Test_A(t *testing.T) {
	config := goauth.NewOauthBasicConfigBuilder().
		SetClientId("dingdjymrdzdxa191wcz").
		SetClientSecret("-maRQbiGZ4KM7BllyOpmYfggNl5fzrof9XWV7jZm_0ZN-sbUbl_6V-jeDdSaeJ28").
		SetScope([]string{"openid"}).
		SetRedirectUri("https://baidu.com/oauth/ididid").Build()
	uri := dingtalk.NewOauthAuthZUriBuilder().WithConfig(config).Build()
	fmt.Printf("\n%s\n", uri)
}

func Test_B(t *testing.T) {
	config := goauth.NewOauthBasicConfigBuilder().
		SetClientId("dingdjymrdzdxa191wcz").
		SetClientSecret("-maRQbiGZ4KM7BllyOpmYfggNl5fzrof9XWV7jZm_0ZN-sbUbl_6V-jeDdSaeJ28").
		SetScope([]string{"openid"}).
		SetRedirectUri("https://baidu.com/oauth/ididid").Build()
	//uri := dingtalk.NewOauthAuthZUriBuilder().WithConfig(config).Build()

	h := dingtalk.NewDingTalkOauthHandler(config)

	//h.GetAccessToken("9824a155b4923f73bf2b718ec22cf2db")
	h.GetUserInfo("b247fdf80a01342f9c09451149121c99")

}

package dingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/c1emon/lemontree/goauth"
	"io"
	"net/http"
)

type OauthHandler struct {
	config *goauth.OauthBasicConfig
}

type CallbackParam struct {
	State string
	Code  string
	Host  string
}

func CallbackPreHandler(r *http.Request) (*CallbackParam, error) {
	if err := r.URL.Query().Get("error"); err != "" {
		return nil, errors.New(err)
	}

	return &CallbackParam{
		State: r.URL.Query().Get("state"),
		Code:  r.URL.Query().Get("authCode"),
		Host:  r.Host,
	}, nil
}

func NewDingTalkOauthHandler(config *goauth.OauthBasicConfig) *OauthHandler {
	return &OauthHandler{
		config: config,
	}
}

var AccessTokenUri = "https://api.dingtalk.com/v1.0/oauth2/userAccessToken"
var UserInfoUri = "https://api.dingtalk.com/v1.0/contact/users/me"

func (d *OauthHandler) GetAccessToken(code string) string {

	v, _ := json.Marshal(struct {
		ClientID     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
		Code         string `json:"code"`
		GrantType    string `json:"grantType"`
		//RefreshToken string `json:"refreshToken"`
	}{
		ClientID:     d.config.ClientId,
		ClientSecret: d.config.ClientSecret,
		Code:         code,
		GrantType:    "authorization_code",
	})

	req, err := http.NewRequest("POST", AccessTokenUri, bytes.NewBuffer(v))
	if err != nil {

	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
	}

	r := struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		ExpireIn     int    `json:"expireIn"`
		CorpID       string `json:"corpId"`
	}{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println(string(body))
	}

	return r.AccessToken
}

func (d *OauthHandler) GetUserInfo(token string) string {
	req, err := http.NewRequest("GET", UserInfoUri, nil)
	if err != nil {

	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-acs-dingtalk-access-token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
	}

	r := struct {
		Nick      string `json:"nick"`
		AvatarURL string `json:"avatarUrl"`
		Mobile    string `json:"mobile"`
		OpenID    string `json:"openId"`
		UnionID   string `json:"unionId"`
		Email     string `json:"email"`
		StateCode string `json:"stateCode"`
	}{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println(string(body))
	}
	fmt.Printf("%s", string(body))

	return r.UnionID
}

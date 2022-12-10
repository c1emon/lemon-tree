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

func (d *OauthHandler) GetAccessToken(code string) (string, error) {

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
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Printf("failed close http body: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	r := struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		ExpireIn     int    `json:"expireIn"`
		CorpID       string `json:"corpId"`
	}{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}

	return r.AccessToken, nil
}

func (d *OauthHandler) GetUserInfo(token string) (*goauth.Identity, error) {
	req, err := http.NewRequest("GET", UserInfoUri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-acs-dingtalk-access-token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Printf("failed close http body: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &goauth.Identity{
		UserName:    r.Nick,
		DisplayName: r.Nick,
		Ids:         []string{r.OpenID},
		UnionIds:    []string{r.UnionID},
		Email:       r.Email,
		Phone:       r.Mobile,
		AvatarUrl:   r.AvatarURL,
	}, nil
}

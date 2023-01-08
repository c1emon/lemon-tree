package dingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/c1emon/lemontree/goauth"
	"io"
	"net/http"
	"net/url"
)

type MiniAppHandler struct {
	config *goauth.OauthBasicConfig
}

func NewDingTalkMiniAppHandler(config *goauth.OauthBasicConfig) *MiniAppHandler {

	return &MiniAppHandler{config: config}
}

func (d *MiniAppHandler) GetAccessToken(code string) (string, error) {
	// code of dingtalk mini app is authCode, which directly get and send by mini app(front end)
	return code, nil
}

var BackEndAccessTokenUri = "https://oapi.dingtalk.com/gettoken"

func getBackEndAccessToken(key, secret string) (string, error) {
	req, err := http.NewRequest("GET", BackEndAccessTokenUri, nil)
	if err != nil {
		return "", err
	}

	req.URL.RawQuery = url.Values{
		"appkey":    {key},
		"appsecret": {secret},
	}.Encode()
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Printf("failed close controller body: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	r := struct {
		ErrCode     int    `json:"errcode"`
		AccessToken string `json:"access_token"`
		ErrMsg      string `json:"errmsg"`
		ExpiresIn   int    `json:"expires_in"`
	}{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}

	if r.ErrCode != 0 {
		return "", errors.New(r.ErrMsg)
	}

	return r.AccessToken, nil
}

var BackEndUserInfoUri = "https://oapi.dingtalk.com/topapi/v2/user/getuserinfo"

func (d *MiniAppHandler) GetIdentity(token string) (*goauth.Identity, error) {

	accessToken, err := getBackEndAccessToken(d.config.ClientId, d.config.ClientSecret)
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(struct {
		Code string `json:"code"`
	}{token})

	req, err := http.NewRequest("POST", BackEndUserInfoUri, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = url.Values{
		"access_token": {accessToken},
	}.Encode()
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Printf("failed close controller body: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := struct {
		ErrCode int `json:"errcode"`
		Result  struct {
			AssociatedUnionId string `json:"associated_unionid"`
			UnionId           string `json:"unionid"`
			DeviceID          string `json:"device_id"`
			SysLevel          int    `json:"sys_level"`
			Name              string `json:"name"`
			Sys               bool   `json:"sys"`
			Userid            string `json:"userid"`
		} `json:"result"`
		ErrMsg string `json:"errmsg"`
	}{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	if r.ErrCode != 0 {
		return nil, errors.New(r.ErrMsg)
	}

	return &goauth.Identity{
		UserName:    r.Result.Name,
		DisplayName: r.Result.Name,
		Ids:         []string{r.Result.Userid},
		UnionIds:    []string{r.Result.UnionId, r.Result.AssociatedUnionId},
	}, nil
}

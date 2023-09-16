package doauth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ory/fosite"
)

type JsonResponseModeHandler struct {
}

func (h *JsonResponseModeHandler) ResponseModes() fosite.ResponseModeTypes {
	return []fosite.ResponseModeType{"json"}
}

func (h *JsonResponseModeHandler) WriteAuthorizeResponse(ctx context.Context, rw http.ResponseWriter, ar fosite.AuthorizeRequester, resp fosite.AuthorizeResponder) {

	j := map[string]any{
		"redirectUri": ar.GetRedirectURI().String(),
	}
	for k, v := range resp.GetParameters() {
		j[k] = v
	}

	b, _ := json.Marshal(j)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")

	_, _ = rw.Write(b)
}

func (h *JsonResponseModeHandler) WriteAuthorizeError(ctx context.Context, rw http.ResponseWriter, ar fosite.AuthorizeRequester, err error) {
	rfcerr := fosite.ErrorToRFC6749Error(err)
	errors := rfcerr.ToValues()
	errors.Set("state", ar.GetState())
	errors.Add("custom_err_param", "bar")

	rw.WriteHeader(http.StatusInternalServerError)
	rw.Header().Set("Content-Type", "application/json")

	b, _ := json.Marshal(errors)
	_, _ = rw.Write(b)

}

type JsonClient struct {
	*fosite.DefaultClient
}

func (c *JsonClient) GetResponseModes() []fosite.ResponseModeType {
	return []fosite.ResponseModeType{"json", "query", "form_post", "", "fragment"}
}

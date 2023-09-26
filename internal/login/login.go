package login

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/c1emon/lemontree/internal/idp"
	"github.com/c1emon/lemontree/pkg/httpx"
	"github.com/gin-gonic/gin"
)

func NewLoginProvider(mgr *idp.IDPManager) *LoginProvider {
	cbUrl := func(ctx context.Context, id string) string {
		return "https://baidu.com"
	}

	return &LoginProvider{
		authCallbackUrl: cbUrl,
		manager:         mgr,
	}
}

type LoginProvider struct {
	authCallbackUrl func(context.Context, string) string
	manager         *idp.IDPManager
	// issuerInterceptor op.IssuerInterceptor
}

func AuthInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		authId := c.Query("auth_id")
		c.Set("auth_id", authId)
		c.Set("oid", authId)
		c.Set("client_id", authId)

		c.Next()
	}
}

func (p *LoginProvider) LoginHandler(c *gin.Context) {

	param := &struct {
		// Id         string          `json:"id,omitempty"`
		// AuthReqId  string          `json:"auth_req_id"`
		ProviderId string          `json:"provider_id"`
		Redirect   bool            `json:"redirect,omitempty"`
		Param      json.RawMessage `json:"param,omitempty"`
	}{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData("json parse err:"))
		return
	}

	// check
	provider, err := p.manager.FindById(c, param.ProviderId)
	if err != nil {

	}

	_, err = provider.GetUser(c, param.Param)
	if err != nil {

	}

	if param.Redirect {
		c.Redirect(http.StatusFound, p.authCallbackUrl(c, "ctx.Id"))
		return
	}

	resp := &struct {
		// Id          string `json:"id"`
		RedirectUri string `json:"redirect_uri"`
	}{RedirectUri: p.authCallbackUrl(c, "ctx.Id")}
	c.JSON(200, httpx.ResponseOK().WithData(resp))
}

package login

import (
	"encoding/json"
	"net/http"

	"github.com/c1emon/lemontree/internal/idp"
	"github.com/c1emon/lemontree/pkg/httpx"
	"github.com/c1emon/lemontree/pkg/logx"
	"github.com/gin-gonic/gin"
)

func NewLoginProvider(mgr *idp.IDPManager, successCbProvider func(*gin.Context, string) string) *LoginProvider {

	return &LoginProvider{
		authCallbackProvider: successCbProvider,
		manager:              mgr,
	}
}

type LoginProvider struct {
	authCallbackProvider func(*gin.Context, string) string
	manager              *idp.IDPManager
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
		Id         string          `json:"id"`
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
		c.JSON(200, httpx.NewResponse(1).WithData("no such idp"))
		return
	}

	user, err := provider.GetUser(c, param.Param)
	if err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData("no such user"))
		return
	}

	logx.GetLogger().Infof("login success for %s", user.Name)

	if param.Redirect {
		c.Redirect(http.StatusFound, p.authCallbackProvider(c, param.Id))
		return
	}

	resp := &struct {
		RedirectUri string `json:"redirect_uri"`
	}{RedirectUri: p.authCallbackProvider(c, param.Id)}

	c.JSON(200, httpx.ResponseOK().WithData(resp))
}

package controller

import (
	"github.com/c1emon/lemontree/goauth"
	"github.com/c1emon/lemontree/goauth/dingtalk"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginHandler struct {
}

func BuildLogin(g *echo.Group) {
	h := &LoginHandler{}

	oauthG := g.Group("/oauth")
	oauthG.Any("/:id", h.oauthHandler)
	oauthG.GET("/authZ/:orgId/:type", h.authZHandler)

}

// oauthHandler handle any controller request to /api/v1/login/oauth/:id
func (h *LoginHandler) oauthHandler(c echo.Context) error {
	id := c.Param("id")

	if c.FormValue("grant_type") == "password" {
		username := c.FormValue("username")
		password := c.FormValue("password")
		//s, err := h.client.Staff.Query().Where(staff.UsernameEQ(username)).Only(context.Background())
		//if err != nil {
		//	//if foundError, ok := err.(ent.NotFoundError); ok {
		//	//
		//	//}
		//	return errors.New("bad username")
		//}

		c.Logger().Infof("id=%s,username=%s,password=%s", id, username, password)
	}

	return c.JSON(200, "hello")
}

func (*LoginHandler) authZHandler(c echo.Context) error {
	t := c.Param("type")
	id := c.Param("orgId")

	// get oauth info by org id and type
	c.Logger().Infof("org: %s type: %s", id, t)

	config := goauth.NewOauthBasicConfigBuilder().
		SetClientId("dingdjymrdzdxa191wcz").
		SetClientSecret("-maRQbiGZ4KM7BllyOpmYfggNl5fzrof9XWV7jZm_0ZN-sbUbl_6V-jeDdSaeJ28").
		SetScope([]string{"openid"}).
		SetRedirectUri("https://baidu.com/oauth/ididid").Build()
	uri := dingtalk.NewOauthAuthZUriBuilder().WithConfig(config).Build()

	if d := c.QueryParam("redirect"); d == "false" {
		return c.JSON(200, uri)
	}
	return c.Redirect(302, uri)

}

func (*LoginHandler) ldapHandler(c echo.Context) error {
	t := c.Param("type")
	return c.String(http.StatusOK, t)
}

func (*LoginHandler) samlHandler(c echo.Context) error {
	t := c.Param("type")
	return c.String(http.StatusOK, t)
}

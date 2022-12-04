package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginHandler struct {
}

type oauthLoginInfo struct {
	GrantType    string `query:"grant_type"`
	ClientId     string `query:"client_id"`
	ClientSecret string `query:"client_secret"`
	Code         string `query:"code"`
	Username     string `query:"username"`
	Password     string `query:"password"`
}

func BuildLogin(g *echo.Group) {
	h := &LoginHandler{}
	g.GET("/login/oauth", h.oauthHandler)
	g.POST("/login/ldap", h.ldapHandler)
	g.POST("/login/saml", h.samlHandler)

	oauthGroup := g.Group("/oauth")
	oauthGroup.Any("/callback/:id", h.oauthCallbackHandler)

}

func (*LoginHandler) oauthCallbackHandler(c echo.Context) error {
	// id := c.Param("id")
	// get code from query
	// find idp by id
	// check login status by idp's api
	// get user info from idp
	// set user status
	// generate token
	// save and return token
	return c.JSON(200, "")
}

func (*LoginHandler) oauthHandler(c echo.Context) error {
	var info oauthLoginInfo
	err := c.Bind(&info)
	if err != nil {
		c.Logger().Info(fmt.Sprint("bad req"))
	}
	return c.JSON(200, info)
}

func (*LoginHandler) ldapHandler(c echo.Context) error {
	t := c.Param("type")
	return c.String(http.StatusOK, t)
}

func (*LoginHandler) samlHandler(c echo.Context) error {
	t := c.Param("type")
	return c.String(http.StatusOK, t)
}

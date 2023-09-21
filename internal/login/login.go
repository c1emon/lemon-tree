package login

import (
	"github.com/c1emon/lemontree/httpx"
	"github.com/c1emon/lemontree/logx"
	"github.com/c1emon/lemontree/model"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	model.BaseFields
	providerType string
	username     string
	code         string
	done         bool
}

type LoginProvider interface {
}

func PreLoginHandler(c *gin.Context) {
	// GET login?id=xxxx
	supportTypes := make([]string, 1)
	supportTypes[0] = "passwd"
	data := struct {
		supportTypes []string
	}{
		supportTypes,
	}

	c.JSON(200, httpx.ResponseOK().WithData(data))
}

func LoginHandler(c *gin.Context) {
	// POST FORM login/{type}
	id := c.PostForm("id")
	username := c.PostForm("username")
	code := c.PostForm("code")

	logx.GetLogger().Infof("[%s] login: %s@%s", id, username, code)

	// c.Redirect(http.StatusFound, "")

	data := struct {
		flag bool
	}{true}

	c.JSON(200, httpx.ResponseOK().WithData(data))

}

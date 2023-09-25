package controller

import (
	"github.com/c1emon/lemontree/internal/factory"
	"github.com/c1emon/lemontree/internal/user"
	"github.com/gin-gonic/gin"
)

func NewUserController() *UserController {
	return &UserController{
		svc: factory.GetUserService(),
	}
}

type UserController struct {
	svc *user.UserService
}

func (u *UserController) CreateUserHandler(c *gin.Context) {
	// param := &struct {
	// 	Username string `json:"provider_id"`
	// 	Redirect bool   `json:"redirect,omitempty"`
	// }{}

	// u.svc.CreateUser(c)
}

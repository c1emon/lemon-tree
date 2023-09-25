package controller

import (
	"github.com/c1emon/lemontree/internal/login"
	"github.com/c1emon/lemontree/pkg/ginx"
)

func CreateLoginRoute() {
	ginx.GetGinEngine().POST("/login", login.NewLoginProvider().LoginHandler)
}

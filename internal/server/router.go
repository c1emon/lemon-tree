package server

import (
	"github.com/c1emon/lemontree/internal/idp"
	"github.com/c1emon/lemontree/internal/login"
	"github.com/c1emon/lemontree/internal/user"
	"github.com/c1emon/lemontree/pkg/ginx"
)

func RegRouter() {
	eng := ginx.GetGinEng()
	userSvc := user.NewUserService()
	idpMgr := idp.NewIDPManager(userSvc)
	loginHandler := login.NewLoginProvider(idpMgr, nil)

	eng.POST("/login", loginHandler.LoginHandler)
}

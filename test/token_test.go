package test

import (
	"github.com/c1emon/lemontree/controller"
	"github.com/c1emon/lemontree/doauth"
	"testing"
)

func Test_GetToken(t *testing.T) {
	h := doauth.NewClemonOAuthEndpoints()

	e := controller.SingletonEchoFactory()

	e.Any("/oauth2/auth", controller.HttpWrapperWithError(h.AuthEndpoint))

	e.Any("/oauth2/token", controller.HttpWrapperWithError(h.TokenEndpoint))

	e.Start(":8080")

}

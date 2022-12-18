package test

import (
	"github.com/c1emon/lemontree/doauth"
	"github.com/c1emon/lemontree/router"
	"testing"
)

func Test_GetToken(t *testing.T) {
	h := doauth.NewClemonOAuthEndpoints()

	e := router.SingletonEchoFactory()

	e.Any("/oauth2/auth", router.HttpWrapperWithError(h.AuthEndpoint))

	e.Any("/oauth2/token", router.HttpWrapperWithError(h.TokenEndpoint))

	e.Start(":8080")

}

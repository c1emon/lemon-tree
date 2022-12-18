package doauth

import (
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/jwt"
	"net/http"
	"time"
)

type HandlerFuncWithError func(http.ResponseWriter, *http.Request) error

type Endpoint interface {
	TokenEndpoint(rw http.ResponseWriter, req *http.Request) error
	AuthEndpoint(rw http.ResponseWriter, req *http.Request) error
	RevokeEndpoint(rw http.ResponseWriter, req *http.Request) error
	IntrospectionEndpoint(rw http.ResponseWriter, req *http.Request) error
}

func newSession(user string) *openid.DefaultSession {
	return &openid.DefaultSession{
		Claims: &jwt.IDTokenClaims{
			Issuer:      "https://fosite.my-application.com",
			Subject:     user,
			Audience:    []string{"https://my-client.my-application.com"},
			ExpiresAt:   time.Now().Add(time.Hour * 6),
			IssuedAt:    time.Now(),
			RequestedAt: time.Now(),
			AuthTime:    time.Now(),
		},
		Headers: &jwt.Headers{
			Extra: make(map[string]interface{}),
		},
	}
}

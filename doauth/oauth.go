package doauth

import (
	"net/http"
)

type HandlerFuncWithError func(http.ResponseWriter, *http.Request) error

type Endpoint interface {
	TokenEndpoint(rw http.ResponseWriter, req *http.Request) error
	AuthEndpoint(rw http.ResponseWriter, req *http.Request) error
	RevokeEndpoint(rw http.ResponseWriter, req *http.Request) error
	IntrospectionEndpoint(rw http.ResponseWriter, req *http.Request) error
}

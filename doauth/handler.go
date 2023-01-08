package doauth

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/c1emon/lemontree/log"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/storage"
	"github.com/ory/fosite/token/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func emptySession(user string) *openid.DefaultSession {
	now := time.Now()
	return &openid.DefaultSession{
		Claims: &jwt.IDTokenClaims{
			Issuer:      "https://fosite.my-application.com",
			Subject:     user,
			Audience:    []string{"https://my-client.my-application.com"},
			ExpiresAt:   now.Add(time.Hour * 6),
			IssuedAt:    now,
			RequestedAt: now,
			AuthTime:    now,
		},
		Headers: &jwt.Headers{
			Extra: make(map[string]interface{}),
		},
	}
}

var _ Endpoint = &DefaultHandler{}

type DefaultHandler struct {
	config *fosite.Config
	oauth2 fosite.OAuth2Provider
	logger *logrus.Logger
}

func NewClemonOAuthEndpoints() *DefaultHandler {
	k, _ := rsa.GenerateKey(rand.Reader, 2048)

	s := storage.NewExampleStore()
	s.Clients["my-client0"] = &JsonClient{
		DefaultClient: &fosite.DefaultClient{
			ID:             "my-client0",
			Secret:         []byte(`$2a$10$IxMdI6d.LIRZPpSfEwNoeu4rY3FhDREsxFJXikcgdRRAStxUlsuEO`),           // = "foobar"
			RotatedSecrets: [][]byte{[]byte(`$2y$10$X51gLxUQJ.hGw1epgHTE5u0bt64xM0COU7K9iAp.OFg8p2pUd.1zC`)}, // = "foobaz",
			RedirectURIs:   []string{"http://localhost:3846/callback"},
			ResponseTypes:  []string{"id_token", "code", "token", "id_token token", "code id_token", "code token", "code id_token token"},
			GrantTypes:     []string{"implicit", "refresh_token", "authorization_code", "password", "client_credentials"},
			Scopes:         []string{"fosite", "openid", "photos", "offline"}},
	}

	return &DefaultHandler{
		oauth2: compose.ComposeAllEnabled(&fosite.Config{
			AccessTokenLifespan:          time.Minute * 30,
			GlobalSecret:                 []byte("some-cool-secret-that-is-32bytes"),
			ResponseModeHandlerExtension: &JsonResponseModeHandler{},
			// ...
		}, s, k),
		logger: log.GetLogger(),
	}
}

func (e *DefaultHandler) AuthEndpoint(rw http.ResponseWriter, req *http.Request) error {
	// This context will be passed to all methods.
	ctx := req.Context()

	// Let's create an AuthorizeRequest object!
	// It will analyze the request and extract important information like scopes, response type and others.
	ar, err := e.oauth2.NewAuthorizeRequest(ctx, req)

	if err != nil {
		e.logger.Errorf("Error occurred in NewAuthorizeRequest: %+v", err)
		e.oauth2.WriteAuthorizeError(ctx, rw, ar, err)
		return nil
	}
	// You have now access to authorizeRequest, Code ResponseTypes, Scopes ...

	var requestedScopes string
	for _, this := range ar.GetRequestedScopes() {
		requestedScopes += fmt.Sprintf(`<li><input type="checkbox" name="scopes" value="%s">%s</li>`, this, this)
	}

	// Normally, this would be the place where you would check if the user is logged in and gives his consent.
	// We're simplifying things and just checking if the request includes a valid username and password
	req.ParseForm()
	if req.PostForm.Get("username") != "peter" {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.Write([]byte(`<h1>Login page</h1>`))
		rw.Write([]byte(fmt.Sprintf(`
			<p>Howdy! This is the log in page. For this example, it is enough to supply the username.</p>
			<form method="post">
				<p>
					By logging in, you consent to grant these scopes:
					<ul>%s</ul>
				</p>
				<input type="text" name="username" /> <small>try peter</small><br>
				<input type="submit">
			</form>
		`, requestedScopes)))
		return nil
	}

	// let's see what scopes the user gave consent to
	for _, scope := range req.PostForm["scopes"] {
		ar.GrantScope(scope)
	}

	// Now that the user is authorized, we set up a session:
	mySessionData := emptySession("peter")

	// When using the HMACSHA strategy you must use something that implements the HMACSessionContainer.
	// It brings you the power of overriding the default values.
	//
	// mySessionData.HMACSession = &strategy.HMACSession{
	//	AccessTokenExpiry: time.Now().Add(time.Day),
	//	AuthorizeCodeExpiry: time.Now().Add(time.Day),
	// }
	//

	// If you're using the JWT strategy, there's currently no distinction between access token and authorize code claims.
	// Therefore, you both access token and authorize code will have the same "exp" claim. If this is something you
	// need let us know on github.
	//
	// mySessionData.JWTClaims.ExpiresAt = time.Now().Add(time.Day)

	// It's also wise to check the requested scopes, e.g.:
	// if ar.GetRequestedScopes().Has("admin") {
	//     controller.Error(rw, "you're not allowed to do that", controller.StatusForbidden)
	//     return
	// }

	// Now we need to get a response. This is the place where the AuthorizeEndpointHandlers kick in and start processing the request.
	// NewAuthorizeResponse is capable of running multiple response type handlers which in turn enables this library
	// to support open id connect.
	response, err := e.oauth2.NewAuthorizeResponse(ctx, ar, mySessionData)

	// Catch any errors, e.g.:
	// * unknown client
	// * invalid redirect
	// * ...
	if err != nil {
		e.logger.Errorf("Error occurred in NewAuthorizeResponse: %+v", err)
		e.oauth2.WriteAuthorizeError(ctx, rw, ar, err)
		return nil
	}

	// Last but not least, send the response!
	e.oauth2.WriteAuthorizeResponse(ctx, rw, ar, response)
	return nil
}

func (e *DefaultHandler) TokenEndpoint(rw http.ResponseWriter, req *http.Request) error {
	// This context will be passed to all methods.
	ctx := req.Context()

	// Create an empty session object which will be passed to the request handlers
	mySessionData := emptySession("")

	// This will create an access request object and iterate through the registered TokenEndpointHandlers to validate the request.
	accessRequest, err := e.oauth2.NewAccessRequest(ctx, req, mySessionData)

	// Catch any errors, e.g.:
	// * unknown client
	// * invalid redirect
	// * ...
	if err != nil {
		e.logger.Errorf("Error occurred in NewAccessRequest: %+v", err)
		e.oauth2.WriteAccessError(ctx, rw, accessRequest, err)
		return nil
	}

	// If this is a client_credentials grant, grant all requested scopes
	// NewAccessRequest validated that all requested scopes the client is allowed to perform
	// based on configured scope matching strategy.
	if accessRequest.GetGrantTypes().ExactOne("client_credentials") {
		for _, scope := range accessRequest.GetRequestedScopes() {
			accessRequest.GrantScope(scope)
		}
	}

	// Next we create a response for the access request. Again, we iterate through the TokenEndpointHandlers
	// and aggregate the result in response.
	response, err := e.oauth2.NewAccessResponse(ctx, accessRequest)
	if err != nil {
		e.logger.Errorf("Error occurred in NewAccessResponse: %+v", err)
		e.oauth2.WriteAccessError(ctx, rw, accessRequest, err)
		return nil
	}

	// All done, send the response.
	e.oauth2.WriteAccessResponse(ctx, rw, accessRequest, response)

	// The client now has a valid access token
	return nil
}

func (e *DefaultHandler) RevokeEndpoint(rw http.ResponseWriter, req *http.Request) error {
	// This context will be passed to all methods.
	ctx := req.Context()

	// This will accept the token revocation request and validate various parameters.
	err := e.oauth2.NewRevocationRequest(ctx, req)

	// All done, send the response.
	e.oauth2.WriteRevocationResponse(ctx, rw, err)
	return err
}

func (e *DefaultHandler) IntrospectionEndpoint(rw http.ResponseWriter, req *http.Request) error {
	ctx := req.Context()
	mySessionData := emptySession("")
	ir, err := e.oauth2.NewIntrospectionRequest(ctx, req, mySessionData)
	if err != nil {
		e.logger.Errorf("Error occurred in NewIntrospectionRequest: %+v", err)
		e.oauth2.WriteIntrospectionError(ctx, rw, err)
		return err
	}

	e.oauth2.WriteIntrospectionResponse(ctx, rw, ir)
	return nil
}

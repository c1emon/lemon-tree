package oidcx

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/c1emon/lemontree/internal/oidcx/client"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v2/pkg/oidc"
	"github.com/zitadel/oidc/v2/pkg/op"
	jose "gopkg.in/square/go-jose.v2"
)

var (
	_ op.Storage                  = &Storage{}
	_ op.ClientCredentialsStorage = &Storage{}
	_ op.SigningKey               = &signingKey{}
)

// accessToken will store an access_token in-memory based on the provided information
func (s *Storage) accessToken(applicationID, refreshTokenID, subject string, audience, scopes []string) (*Token, error) {
	// s.lock.Lock()
	// defer s.lock.Unlock()

	token := &Token{
		ID:             uuid.NewString(),
		ApplicationID:  applicationID,
		RefreshTokenID: refreshTokenID,
		Subject:        subject,
		Audience:       audience,
		Expiration:     time.Now().Add(5 * time.Minute),
		Scopes:         scopes,
	}
	s.tokens[token.ID] = token
	return token, nil
}

type signingKey struct {
	id        string
	algorithm jose.SignatureAlgorithm
	key       *rsa.PrivateKey
}

func (s *signingKey) SignatureAlgorithm() jose.SignatureAlgorithm {
	return s.algorithm
}

func (s *signingKey) Key() interface{} {
	return s.key
}

func (s *signingKey) ID() string {
	return s.id
}

type publicKey struct {
	signingKey
}

func (s *publicKey) ID() string {
	return s.id
}

func (s *publicKey) Algorithm() jose.SignatureAlgorithm {
	return s.algorithm
}

func (s *publicKey) Use() string {
	return "sig"
}

func (s *publicKey) Key() interface{} {
	return &s.key.PublicKey
}

type Service struct {
	keys map[string]*rsa.PublicKey
}

type Storage struct {
	authRequests  map[string]*AuthRequest
	codes         map[string]string
	tokens        map[string]*Token
	refreshTokens map[string]*RefreshToken
	signingKey    signingKey
	serviceUsers  map[string]*client.Client
	services      map[string]Service
	clientSvc     *client.OidcClientService
}

func NewStorage() *Storage {
	return &Storage{}
}

// ClientCredentials implements op.ClientCredentialsStorage.
func (s *Storage) ClientCredentials(ctx context.Context, clientID string, clientSecret string) (op.Client, error) {
	c, err := s.clientSvc.GetById(clientID)
	if err != nil {
		return nil, errors.New("wrong service user or password")
	}

	if c.Secret != clientSecret {
		return nil, errors.New("wrong service user or password")
	}

	return c.WarpToOidcClient(), nil
}

// ClientCredentialsTokenRequest implements op.ClientCredentialsStorage.
func (s *Storage) ClientCredentialsTokenRequest(ctx context.Context, clientID string, scopes []string) (op.TokenRequest, error) {
	client, ok := s.serviceUsers[clientID]
	if !ok {
		return nil, errors.New("wrong service user or password")
	}

	return &oidc.JWTTokenRequest{
		Subject:  client.Id,
		Audience: []string{clientID},
		Scopes:   scopes,
	}, nil
}

// AuthRequestByCode implements op.Storage.
func (s *Storage) AuthRequestByCode(ctx context.Context, code string) (op.AuthRequest, error) {
	requestID, ok := s.codes[code]
	if !ok {
		return nil, fmt.Errorf("code invalid or expired")
	}
	return s.AuthRequestByID(ctx, requestID)
}

// AuthRequestByID implements op.Storage.
func (s *Storage) AuthRequestByID(ctx context.Context, id string) (op.AuthRequest, error) {
	request, ok := s.authRequests[id]
	if !ok {
		return nil, fmt.Errorf("request not found")
	}
	return request, nil
}

// AuthorizeClientIDSecret implements op.Storage.
func (s *Storage) AuthorizeClientIDSecret(ctx context.Context, clientID string, clientSecret string) error {
	// s.lock.Lock()
	// defer s.lock.Unlock()
	c, err := s.clientSvc.GetById(clientID)
	if err != nil {
		return fmt.Errorf("client not found")
	}
	// for this example we directly check the secret
	// obviously you would not have the secret in plain text, but rather hashed and salted (e.g. using bcrypt)
	if c.Secret != clientSecret {
		return fmt.Errorf("invalid secret")
	}
	return nil
}

// createRefreshToken will store a refresh_token in-memory based on the provided information
func (s *Storage) createRefreshToken(accessToken *Token, amr []string, authTime time.Time) (string, error) {
	// s.lock.Lock()
	// defer s.lock.Unlock()
	token := &RefreshToken{
		ID:            accessToken.RefreshTokenID,
		Token:         accessToken.RefreshTokenID,
		AuthTime:      authTime,
		AMR:           amr,
		ApplicationID: accessToken.ApplicationID,
		UserID:        accessToken.Subject,
		Audience:      accessToken.Audience,
		Expiration:    time.Now().Add(5 * time.Hour),
		Scopes:        accessToken.Scopes,
	}
	s.refreshTokens[token.ID] = token
	return token.Token, nil
}

// renewRefreshToken checks the provided refresh_token and creates a new one based on the current
func (s *Storage) renewRefreshToken(currentRefreshToken string) (string, string, error) {
	// s.lock.Lock()
	// defer s.lock.Unlock()
	refreshToken, ok := s.refreshTokens[currentRefreshToken]
	if !ok {
		return "", "", fmt.Errorf("invalid refresh token")
	}
	// deletes the refresh token and all access tokens which were issued based on this refresh token
	delete(s.refreshTokens, currentRefreshToken)
	for _, token := range s.tokens {
		if token.RefreshTokenID == currentRefreshToken {
			delete(s.tokens, token.ID)
			break
		}
	}
	// creates a new refresh token based on the current one
	token := uuid.NewString()
	refreshToken.Token = token
	refreshToken.ID = token
	s.refreshTokens[token] = refreshToken
	return token, refreshToken.ID, nil
}

// CreateAccessAndRefreshTokens implements op.Storage.
func (s *Storage) CreateAccessAndRefreshTokens(ctx context.Context, request op.TokenRequest, currentRefreshToken string) (accessTokenID string, newRefreshTokenID string, expiration time.Time, err error) {
	// generate tokens via token exchange flow if request is relevant
	if teReq, ok := request.(op.TokenExchangeRequest); ok {
		applicationID := teReq.GetClientID()
		authTime := teReq.GetAuthTime()

		refreshTokenID := uuid.NewString()
		accessToken, err := s.accessToken(applicationID, refreshTokenID, request.GetSubject(), request.GetAudience(), request.GetScopes())
		if err != nil {
			return "", "", time.Time{}, err
		}

		refreshToken, err := s.createRefreshToken(accessToken, nil, authTime)
		if err != nil {
			return "", "", time.Time{}, err
		}

		return accessToken.ID, refreshToken, accessToken.Expiration, nil
	}

	var applicationID string
	var authTime time.Time
	var amr []string
	switch req := request.(type) {
	case *AuthRequest:
		applicationID = req.ApplicationID
		authTime = req.authTime
		amr = req.GetAMR()
	case *RefreshTokenRequest:
		applicationID = req.ApplicationID
		authTime = req.AuthTime
		amr = req.AMR
	default:
		applicationID = ""
		authTime = time.Time{}
		amr = nil
	}

	// if currentRefreshToken is empty (Code Flow) we will have to create a new refresh token
	if currentRefreshToken == "" {
		refreshTokenID := uuid.NewString()
		accessToken, err := s.accessToken(applicationID, refreshTokenID, request.GetSubject(), request.GetAudience(), request.GetScopes())
		if err != nil {
			return "", "", time.Time{}, err
		}
		refreshToken, err := s.createRefreshToken(accessToken, amr, authTime)
		if err != nil {
			return "", "", time.Time{}, err
		}
		return accessToken.ID, refreshToken, accessToken.Expiration, nil
	}

	// if we get here, the currentRefreshToken was not empty, so the call is a refresh token request
	// we therefore will have to check the currentRefreshToken and renew the refresh token
	refreshToken, refreshTokenID, err := s.renewRefreshToken(currentRefreshToken)
	if err != nil {
		return "", "", time.Time{}, err
	}
	accessToken, err := s.accessToken(applicationID, refreshTokenID, request.GetSubject(), request.GetAudience(), request.GetScopes())
	if err != nil {
		return "", "", time.Time{}, err
	}
	return accessToken.ID, refreshToken, accessToken.Expiration, nil
}

// CreateAccessToken implements op.Storage.
func (s *Storage) CreateAccessToken(ctx context.Context, request op.TokenRequest) (accessTokenID string, expiration time.Time, err error) {
	var applicationID string
	switch req := request.(type) {
	case *AuthRequest:
		// if authenticated for an app (auth code / implicit flow) we must save the client_id to the token
		applicationID = req.ApplicationID
	case op.TokenExchangeRequest:
		applicationID = req.GetClientID()
	}

	token, err := s.accessToken(applicationID, "", request.GetSubject(), request.GetAudience(), request.GetScopes())
	if err != nil {
		return "", time.Time{}, err
	}
	return token.ID, token.Expiration, nil
}

// CreateAuthRequest implements op.Storage.
func (s *Storage) CreateAuthRequest(ctx context.Context, authReq *oidc.AuthRequest, userID string) (op.AuthRequest, error) {
	// s.lock.Lock()
	// defer s.lock.Unlock()
	// 查看
	if len(authReq.Prompt) == 1 && authReq.Prompt[0] == "none" {
		// With prompt=none, there is no way for the user to log in
		// so return error right away.
		return nil, oidc.ErrLoginRequired()
	}

	// typically, you'll fill your storage / storage model with the information of the passed object
	request := authRequestToInternal(authReq, userID)

	// 生成本次登陆请求的id
	request.SetID(uuid.NewString())

	fmt.Printf("create auth req: %s", request.GetID())
	// and save it in your database (for demonstration purposed we will use a simple map)
	s.authRequests[request.GetID()] = request

	// finally, return the request (which implements the AuthRequest interface of the OP
	return request, nil
}

// DeleteAuthRequest implements op.Storage.
func (s *Storage) DeleteAuthRequest(ctx context.Context, id string) error {
	delete(s.authRequests, id)

	for code, requestID := range s.codes {
		if id == requestID {
			delete(s.codes, code)
			break
		}
	}

	return nil
}

// GetClientByClientID implements op.Storage.
func (s *Storage) GetClientByClientID(ctx context.Context, clientID string) (op.Client, error) {
	// s.lock.Lock()
	// defer s.lock.Unlock()

	c, err := s.clientSvc.GetById(clientID)

	if err != nil {
		return nil, fmt.Errorf("client not found")
	}
	return client.RedirectGlobsClient(c), nil
}

// GetKeyByIDAndClientID implements op.Storage.
func (s *Storage) GetKeyByIDAndClientID(ctx context.Context, keyID string, clientID string) (*jose.JSONWebKey, error) {
	service, ok := s.services[clientID]
	if !ok {
		return nil, fmt.Errorf("clientID not found")
	}
	key, ok := service.keys[keyID]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}
	return &jose.JSONWebKey{
		KeyID: keyID,
		Use:   "sig",
		Key:   key,
	}, nil
}

// GetPrivateClaimsFromScopes implements op.Storage.
func (s *Storage) GetPrivateClaimsFromScopes(ctx context.Context, userID string, clientID string, scopes []string) (map[string]interface{}, error) {

	claims := make(map[string]interface{})

	return claims, nil
}

// GetRefreshTokenInfo implements op.Storage.
func (s *Storage) GetRefreshTokenInfo(ctx context.Context, clientID string, token string) (userID string, tokenID string, err error) {
	refreshToken, ok := s.refreshTokens[token]
	if !ok {
		return "", "", op.ErrInvalidRefreshToken
	}
	return refreshToken.UserID, refreshToken.ID, nil
}

// Health implements op.Storage.
func (s *Storage) Health(ctx context.Context) error {
	return nil
}

// KeySet implements op.Storage.
func (s *Storage) KeySet(ctx context.Context) ([]op.Key, error) {
	return []op.Key{&publicKey{s.signingKey}}, nil
}

// RevokeToken implements op.Storage.
func (s *Storage) RevokeToken(ctx context.Context, tokenOrTokenID string, userID string, clientID string) *oidc.Error {

	accessToken, ok := s.tokens[tokenOrTokenID] // tokenID
	if ok {
		if accessToken.ApplicationID != clientID {
			return oidc.ErrInvalidClient().WithDescription("token was not issued for this client")
		}
		// if it is an access token, just remove it
		// you could also remove the corresponding refresh token if really necessary
		delete(s.tokens, accessToken.ID)
		return nil
	}
	refreshToken, ok := s.refreshTokens[tokenOrTokenID] // token
	if !ok {
		// if the token is neither an access nor a refresh token, just ignore it, the expected behaviour of
		// being not valid (anymore) is achieved
		return nil
	}
	if refreshToken.ApplicationID != clientID {
		return oidc.ErrInvalidClient().WithDescription("token was not issued for this client")
	}
	// if it is a refresh token, you will have to remove the access token as well
	delete(s.refreshTokens, refreshToken.ID)
	for _, accessToken := range s.tokens {
		if accessToken.RefreshTokenID == refreshToken.ID {
			delete(s.tokens, accessToken.ID)
			return nil
		}
	}
	return nil
}

// SaveAuthCode implements op.Storage.
func (s *Storage) SaveAuthCode(ctx context.Context, id string, code string) error {
	s.codes[code] = id
	return nil
}

// SetIntrospectionFromToken implements op.Storage.
func (s *Storage) SetIntrospectionFromToken(ctx context.Context, introspection *oidc.IntrospectionResponse, tokenID, subject, clientID string) error {
	token, ok := s.tokens[tokenID]
	if !ok {
		return fmt.Errorf("token is invalid or has expired")
	}
	// check if the client is part of the requested audience
	for _, aud := range token.Audience {
		if aud == clientID {
			// the introspection response only has to return a boolean (active) if the token is active
			// this will automatically be done by the library if you don't return an error
			// you can also return further information about the user / associated token
			// e.g. the userinfo (equivalent to userinfo endpoint)

			userInfo := new(oidc.UserInfo)
			userInfo.Email = "test@email.ok"
			// err := s.setUserinfo(ctx, userInfo, subject, clientID, token.Scopes)
			// if err != nil {
			// 	return err
			// }
			introspection.SetUserInfo(userInfo)
			//...and also the requested scopes...
			introspection.Scope = token.Scopes
			//...and the client the token was issued to
			introspection.ClientID = token.ApplicationID
			return nil
		}
	}
	return fmt.Errorf("token is not valid for this client")
}

// SetUserinfoFromScopes implements op.Storage.
func (s *Storage) SetUserinfoFromScopes(ctx context.Context, userinfo *oidc.UserInfo, userID string, clientID string, scopes []string) error {
	return nil
}

// SetUserinfoFromToken implements op.Storage.
func (s *Storage) SetUserinfoFromToken(ctx context.Context, userinfo *oidc.UserInfo, tokenID string, subject string, origin string) error {
	userinfo.Email = "SetUserinfoFromToken@email.ok"
	return nil
}

// SignatureAlgorithms implements op.Storage.
func (s *Storage) SignatureAlgorithms(ctx context.Context) ([]jose.SignatureAlgorithm, error) {
	return []jose.SignatureAlgorithm{s.signingKey.algorithm}, nil
}

// SigningKey implements op.Storage.
func (s *Storage) SigningKey(ctx context.Context) (op.SigningKey, error) {
	return &s.signingKey, nil
}

// TerminateSession implements op.Storage.
func (s *Storage) TerminateSession(ctx context.Context, userID string, clientID string) error {
	for _, token := range s.tokens {
		if token.ApplicationID == clientID && token.Subject == userID {
			delete(s.tokens, token.ID)
			delete(s.refreshTokens, token.RefreshTokenID)
		}
	}
	return nil
}

// TokenRequestByRefreshToken implements op.Storage.
func (s *Storage) TokenRequestByRefreshToken(ctx context.Context, refreshToken string) (op.RefreshTokenRequest, error) {
	token, ok := s.refreshTokens[refreshToken]
	if !ok {
		return nil, fmt.Errorf("invalid refresh_token")
	}
	return RefreshTokenRequestFromBusiness(token), nil
}

// ValidateJWTProfileScopes implements op.Storage.
func (s *Storage) ValidateJWTProfileScopes(ctx context.Context, userID string, scopes []string) ([]string, error) {
	allowedScopes := make([]string, 0)
	for _, scope := range scopes {
		if scope == oidc.ScopeOpenID {
			allowedScopes = append(allowedScopes, scope)
		}
	}
	return allowedScopes, nil
}

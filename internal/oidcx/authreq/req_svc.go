package authreq

import (
	"errors"
	"time"

	"github.com/c1emon/lemontree/pkg/cachex"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

func NewAuthReqService(cacher cachex.Cacher) *AuthReqService {
	return &AuthReqService{cacher: cacher}
}

type AuthReqService struct {
	cacher cachex.Cacher
}

func (s *AuthReqService) GetAuthReq(id string) (AuthRequest, error) {
	val, ok := s.cacher.Get(id)
	if !ok {
		return AuthRequest{}, errors.New("no such auth req")
	}
	if req, ok := val.(AuthRequest); ok {
		return req, nil
	} else {
		return AuthRequest{}, errors.New("unknown auth req type")
	}
}

func (s *AuthReqService) SetAuthReq(req AuthRequest) {
	s.cacher.Set(req.Id, req)
}

func (s *AuthReqService) DelAuthReq(id string) {
	s.cacher.Del(id)
}

func promptToInternal(oidcPrompt oidc.SpaceDelimitedArray) []string {
	prompts := make([]string, len(oidcPrompt))
	for _, oidcPrompt := range oidcPrompt {
		switch oidcPrompt {
		case oidc.PromptNone,
			oidc.PromptLogin,
			oidc.PromptConsent,
			oidc.PromptSelectAccount:
			prompts = append(prompts, oidcPrompt)
		}
	}
	return prompts
}

func maxAgeToInternal(maxAge *uint) *time.Duration {
	if maxAge == nil {
		return nil
	}
	dur := time.Duration(*maxAge) * time.Second
	return &dur
}

func AuthRequestToInternal(authReq *oidc.AuthRequest, userID string) AuthRequest {
	return AuthRequest{
		ApplicationID: authReq.ClientID,
		CallbackURI:   authReq.RedirectURI,
		TransferState: authReq.State,
		Prompt:        promptToInternal(authReq.Prompt),
		UiLocales:     authReq.UILocales,
		LoginHint:     authReq.LoginHint,
		MaxAuthAge:    maxAgeToInternal(authReq.MaxAge),
		UserID:        userID,
		Scopes:        authReq.Scopes,
		ResponseType:  authReq.ResponseType,
		Nonce:         authReq.Nonce,
		CodeChallenge: &OIDCCodeChallenge{
			Challenge: authReq.CodeChallenge,
			Method:    string(authReq.CodeChallengeMethod),
		},
	}
}

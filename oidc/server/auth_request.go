package server

import (
	"time"

	"github.com/c1emon/lemontree/model"
	"github.com/zitadel/oidc/v2/pkg/oidc"
	"golang.org/x/text/language"
)

type OIDCCodeChallenge struct {
	Challenge string
	Method    string
}

type AuthRequest struct {
	model.BaseFields
	ApplicationID string
	CallbackURI   string
	TransferState string
	Prompt        []string
	UiLocales     []language.Tag
	LoginHint     string
	MaxAuthAge    *time.Duration
	UserID        string
	Scopes        []string
	ResponseType  oidc.ResponseType
	Nonce         string
	CodeChallenge *OIDCCodeChallenge

	done     bool
	authTime time.Time
}

func (ar *AuthRequest) SetID(id string) {
	ar.Id = id
}

func (ar *AuthRequest) GetID() string {
	return ar.Id
}

func (ar *AuthRequest) GetACR() string {
	return ""
}

func (ar *AuthRequest) GetAMR() []string {
	// this example only uses password for authentication
	if ar.done {
		return []string{"pwd"}
	}
	return nil
}

func (ar *AuthRequest) GetAudience() []string {
	return []string{ar.ApplicationID} // this example will always just use the client_id as audience
}

func (ar *AuthRequest) GetAuthTime() time.Time {
	return ar.authTime
}

func (ar *AuthRequest) GetClientID() string {
	return ar.ApplicationID
}

func (ar *AuthRequest) GetCodeChallenge() *oidc.CodeChallenge {

	if ar.CodeChallenge == nil {
		return nil
	}
	challengeMethod := oidc.CodeChallengeMethodPlain
	if ar.CodeChallenge.Method == "S256" {
		challengeMethod = oidc.CodeChallengeMethodS256
	}
	return &oidc.CodeChallenge{
		Challenge: ar.CodeChallenge.Challenge,
		Method:    challengeMethod,
	}
}

func (ar *AuthRequest) GetNonce() string {
	return ar.Nonce
}

func (ar *AuthRequest) GetRedirectURI() string {
	return ar.CallbackURI
}

func (ar *AuthRequest) GetResponseType() oidc.ResponseType {
	return ar.ResponseType
}

func (ar *AuthRequest) GetResponseMode() oidc.ResponseMode {
	return ""
}

func (ar *AuthRequest) GetScopes() []string {
	return ar.Scopes
}

func (ar *AuthRequest) GetState() string {
	return ar.TransferState
}

func (ar *AuthRequest) GetSubject() string {
	return ar.UserID
}

func (ar *AuthRequest) Done() bool {
	return ar.done
}

func PromptToInternal(oidcPrompt oidc.SpaceDelimitedArray) []string {
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

func MaxAgeToInternal(maxAge *uint) *time.Duration {
	if maxAge == nil {
		return nil
	}
	dur := time.Duration(*maxAge) * time.Second
	return &dur
}

func authRequestToInternal(authReq *oidc.AuthRequest, userID string) *AuthRequest {
	now := time.Now()
	return &AuthRequest{
		BaseFields: model.BaseFields{
			CreateTime: now,
			UpdateTime: now,
		},
		ApplicationID: authReq.ClientID,
		CallbackURI:   authReq.RedirectURI,
		TransferState: authReq.State,
		Prompt:        PromptToInternal(authReq.Prompt),
		UiLocales:     authReq.UILocales,
		LoginHint:     authReq.LoginHint,
		MaxAuthAge:    MaxAgeToInternal(authReq.MaxAge),
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

package authreq

import (
	"time"

	"github.com/zitadel/oidc/v2/pkg/oidc"
	"github.com/zitadel/oidc/v2/pkg/op"
	"golang.org/x/text/language"
)

type OIDCCodeChallenge struct {
	Challenge string
	Method    string
}

// check
var _ op.AuthRequest = &AuthRequest{}

type AuthRequest struct {
	Id            string             `json:"id"`
	ApplicationID string             `json:"app_id"`
	CallbackURI   string             `json:"cb_uri"`
	TransferState string             `json:"transfer_state"`
	Prompt        []string           `json:"promot"`
	UiLocales     []language.Tag     `json:"locales"`
	LoginHint     string             `json:"login_hint"`
	MaxAuthAge    *time.Duration     `json:"auth_age"`
	UserID        string             `json:"uid"`
	Scopes        []string           `json:"scopes"`
	ResponseType  oidc.ResponseType  `json:"response_type"`
	Nonce         string             `json:"nonce"`
	CodeChallenge *OIDCCodeChallenge `json:"code_challenge"`
	Success       bool               `json:"success"`
	AuthTime      time.Time          `json:"auth_time"`
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
	if ar.Success {
		return []string{"pwd"}
	}
	return nil
}

func (ar *AuthRequest) GetAudience() []string {
	return []string{ar.ApplicationID} // this example will always just use the client_id as audience
}

func (ar *AuthRequest) GetAuthTime() time.Time {
	return ar.AuthTime
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
	return ar.Success
}

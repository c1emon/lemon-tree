package client

import (
	"context"
	"time"

	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/pkg/errors"
	"github.com/zitadel/oidc/v2/pkg/oidc"
	"github.com/zitadel/oidc/v2/pkg/op"
)

func NewOidcClientService() *OidcClientService {
	repo := &gormOidcClientRepository{
		db: gormx.GetGormDB(),
	}
	repo.InitDB()
	return &OidcClientService{repository: repo}
}

type OidcClientService struct {
	repository OidcClientRepository
}

func (s *OidcClientService) Create(ctx context.Context) {
	c := &Client{
		Name:                           "test_client_1",
		Secret:                         "secret112233",
		RedirectURIs:                   []string{"http://a.c/red", "http://b.c/redir"},
		ApplicationType:                op.ApplicationTypeWeb,
		AuthMethod:                     oidc.AuthMethodBasic,
		LoginURL:                       "http://a.c/login",
		ResponseTypes:                  []oidc.ResponseType{oidc.ResponseTypeCode, oidc.ResponseTypeIDToken},
		GrantTypes:                     []oidc.GrantType{oidc.GrantTypeCode, oidc.GrantTypeBearer},
		AccessTokenType:                op.AccessTokenTypeBearer,
		DevMode:                        false,
		IdTokenUserinfoClaimsAssertion: false,
		ClockSkew:                      time.Duration(time.Duration.Minutes(5)),
		OrganizationId:                 "12safasf",
	}

	s.repository.CreateOne(ctx, c)
}

func (s *OidcClientService) GetById(id string) (*Client, error) {
	org, err := s.repository.GetOneById(context.Background(), id)
	return org, errors.WithMessage(err, "oidc client")
}

type warpToRedirectGlobs struct {
	*OidcClient
}

// RedirectURIGlobs provide wildcarding for additional valid redirects
func (c warpToRedirectGlobs) RedirectURIGlobs() []string {
	return c.entity.RedirectURIGlobs
}

// PostLogoutRedirectURIGlobs provide extra wildcarding for additional valid redirects
func (c warpToRedirectGlobs) PostLogoutRedirectURIGlobs() []string {
	return c.entity.PostLogoutRedirectURIGlobs
}

// RedirectGlobsClient wraps the client in a op.HasRedirectGlobs
// only if DevMode is enabled.
func RedirectGlobsClient(client *Client) op.Client {
	oidcClient := &OidcClient{entity: client}
	if client.DevMode {
		return warpToRedirectGlobs{oidcClient}
	}
	return oidcClient
}

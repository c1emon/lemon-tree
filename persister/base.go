package persister

import (
	"context"
	"github.com/c1emon/lemontree/ent"
)

type InitDB interface {
	Init()
}

type OrganizationRepository interface {
	SetClient(any)
	Create(context.Context, ent.Organization) (*ent.Organization, error)
	Delete(context.Context, string) error
	Update(context.Context, string, ent.Organization) (*ent.Organization, error)
	GetOneById(context.Context, string) (*ent.Organization, error)
}

type OAuthClientRepository interface {
	SetClient(any)
	Create(context.Context, ent.OAuthClient) (*ent.OAuthClient, error)
	Delete(context.Context, string, ent.OAuthClient) error
	Update(context.Context, string, ent.OAuthClient) (*ent.OAuthClient, error)
	GetOneById(context.Context, string) (*ent.OAuthClient, error)
}

type OIDCClientRepository interface {
}

type UserRepository interface {
	SetClient(any)
	Create(context.Context, ent.User) (*ent.User, error)
	Delete(context.Context, string, ent.User) error
	Update(context.Context, string, ent.User) (*ent.User, error)
	GetOneById(context.Context, string) (*ent.User, error)
}

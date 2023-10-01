package client

import (
	"context"
	"fmt"

	"github.com/c1emon/lemontree/pkg/errorx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// check
var _ OidcClientRepository = &gormOidcClientRepository{}

type gormOidcClientRepository struct {
	db *gorm.DB
}

// CreateOne implements OidcClientRepository.
func (r *gormOidcClientRepository) CreateOne(ctx context.Context, client *Client) error {
	err := r.db.WithContext(ctx).Create(client).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("oidc client %s", client.Id))
}

// DeleteOneById implements OidcClientRepository.
func (r *gormOidcClientRepository) DeleteOneById(context.Context, string) error {
	panic("unimplemented")
}

// FindByOidAndName implements OidcClientRepository.
func (r *gormOidcClientRepository) FindByOidAndName(string, string) *Client {
	panic("unimplemented")
}

// GetOneById implements OidcClientRepository.
func (r *gormOidcClientRepository) GetOneById(ctx context.Context, id string) (*Client, error) {
	client := &Client{}
	client.Id = id
	res := r.db.WithContext(ctx).First(client)
	return client, errors.Wrap(errorx.From(res.Error), fmt.Sprintf("id %s", id))
}

// InitDB implements OidcClientRepository.
func (r *gormOidcClientRepository) InitDB() error {
	return r.db.AutoMigrate(&Client{})
}

// UpdateOneById implements OidcClientRepository.
func (r *gormOidcClientRepository) UpdateOneById(context.Context, string, *Client) error {
	panic("unimplemented")
}

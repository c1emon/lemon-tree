package oidc

import (
	"context"
	"fmt"

	"github.com/c1emon/lemontree/model"
	"github.com/c1emon/lemontree/pkg/errorx"
	"github.com/c1emon/lemontree/pkg/httpx"
	"github.com/c1emon/lemontree/pkg/persister"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DefaultClientRepository interface {
	model.BaseRepository[Client]
	GetOneByName(context.Context, string) (*Client, error)
	GetAllByName(context.Context, httpx.Pageable, string) []Client
}

var _ DefaultClientRepository = &GormClientRepository{}

type GormClientRepository struct {
	db *gorm.DB
}

func NewGormClientRepository() *GormClientRepository {
	r := &GormClientRepository{
		db: persister.GetDB(),
	}
	_ = r.InitDB()
	return r
}

// CreateOne implements DefaultClientRepository.
func (r *GormClientRepository) CreateOne(ctx context.Context, client *Client) error {
	err := r.db.WithContext(ctx).Create(client).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("oidc client %s", client.Id))
}

// DeleteOneById implements DefaultClientRepository.
func (r *GormClientRepository) DeleteOneById(context.Context, string) error {
	panic("unimplemented")
}

// GetAllByName implements DefaultClientRepository.
func (r *GormClientRepository) GetAllByName(context.Context, httpx.Pageable, string) []Client {
	panic("unimplemented")
}

// GetOneById implements DefaultClientRepository.
func (r *GormClientRepository) GetOneById(context.Context, string) (*Client, error) {
	panic("unimplemented")
}

// GetOneByName implements DefaultClientRepository.
func (r *GormClientRepository) GetOneByName(context.Context, string) (*Client, error) {
	panic("unimplemented")
}

// InitDB implements DefaultClientRepository.
func (r *GormClientRepository) InitDB() error {
	return r.db.AutoMigrate(&Client{})
}

// UpdateOneById implements DefaultClientRepository.
func (r *GormClientRepository) UpdateOneById(context.Context, string, *Client) error {
	panic("unimplemented")
}

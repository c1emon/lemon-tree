package idp

import (
	"context"
	"fmt"

	"github.com/c1emon/lemontree/internal/user"
	"github.com/c1emon/lemontree/pkg/errorx"
	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type IDProvider interface {
	GetUser(context.Context, any) (*user.User, error)
	Support(string) bool
}

type identityProviderConfig struct {
	gormx.BaseFields
	Name         string         `json:"name" gorm:"column:name;type:varchar(256);uniqueIndex:udx_org_user;not null"`
	Oid          string         `json:"oid" gorm:"column:oid;type:varchar(256);uniqueIndex:udx_org_user;not null"`
	ProviderType string         `json:"provider_type" gorm:"column:provider_type;not null"`
	Config       datatypes.JSON `json:"config" gorm:"column:config"`
}

func (identityProviderConfig) TableName() string {
	return "identity_provider_configs"
}

type identityProviderConfigRepository interface {
	gormx.BaseRepository[identityProviderConfig]
}

// check
var _ identityProviderConfigRepository = &gormIdentityProviderConfigRepository{}

type gormIdentityProviderConfigRepository struct {
	db *gorm.DB
}

// CreateOne implements identityProviderConfigRepository.
func (r *gormIdentityProviderConfigRepository) CreateOne(ctx context.Context, config *identityProviderConfig) error {
	err := r.db.WithContext(ctx).Create(config).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("provider config id %s", config.Id))
}

// DeleteOneById implements identityProviderConfigRepository.
func (r *gormIdentityProviderConfigRepository) DeleteOneById(context.Context, string) error {
	panic("unimplemented")
}

// GetOneById implements identityProviderConfigRepository.
func (r *gormIdentityProviderConfigRepository) GetOneById(ctx context.Context, id string) (*identityProviderConfig, error) {
	config := &identityProviderConfig{}
	config.Id = id
	res := r.db.WithContext(ctx).First(config)
	return config, errors.Wrap(errorx.From(res.Error), fmt.Sprintf("id %s", id))
}

// InitDB implements identityProviderConfigRepository.
func (r *gormIdentityProviderConfigRepository) InitDB() error {
	return r.db.AutoMigrate(&identityProviderConfig{})
}

// UpdateOneById implements identityProviderConfigRepository.
func (r *gormIdentityProviderConfigRepository) UpdateOneById(context.Context, string, *identityProviderConfig) error {
	panic("unimplemented")
}

func NewIDPManager(userSvc *user.UserService) *IDPManager {

	repo := &gormIdentityProviderConfigRepository{
		db: gormx.GetDB(),
	}
	repo.InitDB()

	return &IDPManager{
		idpConfigRepository: repo,
		userService:         userSvc,
	}
}

type IDPManager struct {
	idpConfigRepository identityProviderConfigRepository
	userService         *user.UserService
}

func (m *IDPManager) FindById(ctx context.Context, id string) (IDProvider, error) {

	conf, err := m.idpConfigRepository.GetOneById(ctx, id)
	if err != nil {
		return nil, err
	}

	switch conf.ProviderType {
	case "password":
		return NewPasswdIDP(m.userService, nil), nil
	}

	return nil, nil
}

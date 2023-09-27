package user

import (
	"context"
	"fmt"

	"github.com/c1emon/lemontree/pkg/errorx"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// check
var _ UserIdentityRepository = &gormUserIdentityRepository{}

type gormUserIdentityRepository struct {
	db *gorm.DB
}

// Validate implements UserIdentityRepository.
func (r *gormUserIdentityRepository) Validate(ctx context.Context, oid string, builder func(builder func() *datatypes.JSONQueryExpression) []any) (string, error) {
	identity := &UserIdentity{}
	identity.Oid = oid

	jsonConds := builder(func() *datatypes.JSONQueryExpression {
		return datatypes.JSONQuery("identity")
	})

	res := r.db.First(identity, jsonConds...)
	return identity.Uid, errorx.From(res.Error)
}

// CreateOne implements UserIdentityRepository.
func (r *gormUserIdentityRepository) CreateOne(ctx context.Context, identity *UserIdentity) error {
	err := r.db.WithContext(ctx).Create(identity).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("uid %s", identity.Uid))
}

// DeleteOneById implements UserIdentityRepository.
func (r *gormUserIdentityRepository) DeleteOneById(context.Context, string) error {
	panic("unimplemented")
}

// FindByUidAndIdpId implements UserIdentityRepository.
func (r *gormUserIdentityRepository) FindByUidAndIdpId(ctx context.Context, uid string, idpId string) (*UserIdentity, error) {
	panic("unimplemented")
}

// GetOneById implements UserIdentityRepository.
func (r *gormUserIdentityRepository) GetOneById(context.Context, string) (*UserIdentity, error) {
	panic("unimplemented")
}

// InitDB implements UserIdentityRepository.
func (r *gormUserIdentityRepository) InitDB() error {
	return r.db.AutoMigrate(&UserIdentity{})
}

// UpdateOneById implements UserIdentityRepository.
func (r *gormUserIdentityRepository) UpdateOneById(context.Context, string, *UserIdentity) error {
	panic("unimplemented")
}

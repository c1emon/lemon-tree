package user

import (
	"context"
	"fmt"

	"github.com/c1emon/lemontree/pkg/errorx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// check
var _ UserRepository = &gormUserRepository{}

type gormUserRepository struct {
	db *gorm.DB
}

// CheckPasswd implements UserRepository.
func (r *gormUserRepository) CheckPasswd(string) bool {
	panic("unimplemented")
}

// CreateOne implements UserRepository.
func (r *gormUserRepository) CreateOne(ctx context.Context, u *User) error {
	err := r.db.WithContext(ctx).Create(u).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("username %s", u.UserName))
}

// DeleteOneById implements UserRepository.
func (r *gormUserRepository) DeleteOneById(context.Context, string) error {
	panic("unimplemented")
}

// FindByOidAndName implements UserRepository.
func (r *gormUserRepository) FindByOidAndName(string, string) (*User, error) {
	panic("unimplemented")
}

// GetOneById implements UserRepository.
func (r *gormUserRepository) GetOneById(ctx context.Context, id string) (*User, error) {
	user := &User{}
	user.Id = id
	res := r.db.WithContext(ctx).First(user)
	return user, errors.Wrap(errorx.From(res.Error), fmt.Sprintf("id %s", id))
}

// InitDB implements UserRepository.
func (r *gormUserRepository) InitDB() error {
	return r.db.AutoMigrate(&User{})
}

// UpdateOneById implements UserRepository.
func (r *gormUserRepository) UpdateOneById(context.Context, string, *User) error {
	panic("unimplemented")
}

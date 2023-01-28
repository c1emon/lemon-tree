package repository

import (
	"context"
	"fmt"
	"github.com/c1emon/lemontree/errorc"
	"github.com/c1emon/lemontree/model"
	"github.com/c1emon/lemontree/persister"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DefaultOrganizationRepository interface {
	BaseRepository[model.Organization]
	AddDepartment(context.Context, model.Department) error
}

// check
var _ DefaultOrganizationRepository = &GormOrganizationRepository{}

type GormOrganizationRepository struct {
	db *gorm.DB
}

func NewGormOrganizationRepository() *GormOrganizationRepository {
	r := &GormOrganizationRepository{
		db: persister.GetDB(),
	}
	_ = r.InitDB()
	return r
}

func (r *GormOrganizationRepository) AddDepartment(ctx context.Context, department model.Department) error {
	return nil
}

func (r *GormOrganizationRepository) CreateOne(ctx context.Context, org *model.Organization) error {
	err := r.db.WithContext(ctx).Create(org).Error
	return errors.Wrap(errorc.From(err), fmt.Sprintf("name %s", org.Name))
}

func (r *GormOrganizationRepository) GetOneById(ctx context.Context, id string) (*model.Organization, error) {
	org := &model.Organization{}
	org.Id = id
	res := r.db.WithContext(ctx).First(org)
	return org, errors.Wrap(errorc.From(res.Error), fmt.Sprintf("id %s", id))
}

func (r *GormOrganizationRepository) UpdateOneById(ctx context.Context, id string, org *model.Organization) (*model.Organization, error) {

	o, err := r.GetOneById(ctx, id)
	if err != nil {
		return nil, errors.Wrap(errorc.From(err), fmt.Sprintf("id %s", id))
	}
	res := r.db.WithContext(ctx).Model(o).Updates(*org)
	o, err = r.GetOneById(ctx, id)
	if err != nil {
		return nil, errors.Wrap(errorc.From(err), fmt.Sprintf("id %s", id))
	}
	return o, errors.Wrap(errorc.From(res.Error), fmt.Sprintf("id %s", id))
}

func (r *GormOrganizationRepository) DeleteOneById(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&model.Organization{}, id).Error
	return errors.Wrap(errorc.From(err), fmt.Sprintf("id %s", id))
}

func (r *GormOrganizationRepository) InitDB() error {
	return r.db.AutoMigrate(&model.Organization{})
}

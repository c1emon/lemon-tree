package org

import (
	"context"
	"fmt"

	"github.com/c1emon/lemontree/pkg/errorx"
	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/c1emon/lemontree/pkg/httpx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// check
var _ OrganizationRepository = &gormOrganizationRepository{}

type gormOrganizationRepository struct {
	db *gorm.DB
}

func NewGormOrganizationRepository() *gormOrganizationRepository {
	r := &gormOrganizationRepository{
		db: gormx.GetGormDB(),
	}
	_ = r.InitDB()
	return r
}

func (r *gormOrganizationRepository) CreateOne(ctx context.Context, org *Organization) error {
	err := r.db.WithContext(ctx).Create(org).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("name %s", org.Name))
}

func (r *gormOrganizationRepository) GetOneById(ctx context.Context, id string) (*Organization, error) {
	org := &Organization{}
	org.Id = id
	res := r.db.WithContext(ctx).First(org)
	return org, errors.Wrap(errorx.From(res.Error), fmt.Sprintf("id %s", id))
}

func (r *gormOrganizationRepository) GetOneByName(ctx context.Context, name string) (*Organization, error) {
	org := &Organization{}
	res := r.db.WithContext(ctx).Where("name = ?", name).First(org)
	return org, errors.Wrap(errorx.From(res.Error), fmt.Sprintf("name %s", name))
}

func (r *gormOrganizationRepository) GetAllByName(ctx context.Context, pageable httpx.Pageable, name string) []Organization {

	var orgs []Organization
	var total int64 = 0

	query := r.db.WithContext(ctx).
		Limit(pageable.GetPageSize()).
		Offset(pageable.GetOffset()).
		Where(" name LIKE % ? % ", name)
	for _, s := range pageable.GetSorts() {
		query.Order(s.Sql())
	}
	query.Find(&orgs)

	r.db.WithContext(ctx).
		Model(&Organization{}).
		Where(" name LIKE % ? % ", name).
		Count(&total)
	pageable.SetTotal(total)

	//if err != nil {
	//	return
	//}
	return orgs
}

func (r *gormOrganizationRepository) UpdateOneById(ctx context.Context, id string, org *Organization) error {

	err := r.db.WithContext(ctx).
		Model(&Organization{BaseFields: gormx.BaseFields{Id: id}}).
		Updates(*org).Error
	if err != nil {
		return errors.Wrap(errorx.From(err), fmt.Sprintf("id %s", id))
	}

	return nil
}

func (r *gormOrganizationRepository) DeleteOneById(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).
		Delete(&Organization{BaseFields: gormx.BaseFields{Id: id}}).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("id %s", id))
}

func (r *gormOrganizationRepository) InitDB() error {
	return r.db.AutoMigrate(&Organization{})
}

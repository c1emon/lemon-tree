package repository

import (
	"context"
	"fmt"
	"github.com/c1emon/lemontree/errorc"
	"github.com/c1emon/lemontree/httpx"
	"github.com/c1emon/lemontree/model"
	"github.com/c1emon/lemontree/persister"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DefaultOrganizationRepository interface {
	BaseRepository[model.Organization]
	AddDepartment(context.Context, model.Department) error
	GetOneByName(context.Context, string) (*model.Organization, error)
	GetAllByName(context.Context, httpx.Pageable, string) []model.Organization
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

func (r *GormOrganizationRepository) GetOneByName(ctx context.Context, name string) (*model.Organization, error) {
	org := &model.Organization{}
	res := r.db.WithContext(ctx).Where("name = ?", name).First(org)
	return org, errors.Wrap(errorc.From(res.Error), fmt.Sprintf("name %s", name))
}

func (r *GormOrganizationRepository) GetAllByName(ctx context.Context, pageable httpx.Pageable, name string) []model.Organization {

	var orgs []model.Organization
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
		Model(&model.Organization{}).
		Where(" name LIKE % ? % ", name).
		Count(&total)
	pageable.SetTotal(total)

	//if err != nil {
	//	return
	//}
	return orgs
}

func (r *GormOrganizationRepository) UpdateOneById(ctx context.Context, id string, org *model.Organization) error {

	err := r.db.WithContext(ctx).
		Model(&model.Organization{BaseField: model.BaseField{Id: id}}).
		Updates(*org).Error
	if err != nil {
		return errors.Wrap(errorc.From(err), fmt.Sprintf("id %s", id))
	}

	return nil
}

func (r *GormOrganizationRepository) DeleteOneById(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).
		Delete(&model.Organization{BaseField: model.BaseField{Id: id}}).Error
	return errors.Wrap(errorc.From(err), fmt.Sprintf("id %s", id))
}

func (r *GormOrganizationRepository) InitDB() error {
	return r.db.AutoMigrate(&model.Organization{})
}
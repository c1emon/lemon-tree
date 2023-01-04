package persister

import (
	"context"
	"github.com/c1emon/lemontree/model"
	"gorm.io/gorm"
)

// check
var _ model.OrganizationRepository = &GormOrganizationRepository{}

type GormOrganizationRepository struct {
	db *gorm.DB
}

func NewDefaultOrganizationRepository() *GormOrganizationRepository {
	r := &GormOrganizationRepository{
		db: GetDB(),
	}
	_ = r.InitDB()
	return r
}

func (r *GormOrganizationRepository) AddDepartment(ctx context.Context, department model.Department) error {
	return nil
}

func (r *GormOrganizationRepository) CreateOne(ctx context.Context, org *model.Organization) (*model.Organization, error) {

	return org, r.db.WithContext(ctx).Create(org).Error
}

func (r *GormOrganizationRepository) GetOneById(ctx context.Context, id string) (*model.Organization, error) {
	org := &model.Organization{}
	org.Id = id
	res := r.db.WithContext(ctx).First(org)
	return org, res.Error
}

func (r *GormOrganizationRepository) UpdateOneById(ctx context.Context, id string, org *model.Organization) (*model.Organization, error) {
	o, err := r.GetOneById(ctx, id)
	if err != nil {

	}
	res := r.db.WithContext(ctx).Model(o).Updates(*org)
	o, err = r.GetOneById(ctx, id)
	if err != nil {

	}
	return o, res.Error
}

func (r *GormOrganizationRepository) DeleteOneById(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Organization{}, id).Error
}

func (r *GormOrganizationRepository) InitDB() error {
	return r.db.AutoMigrate(&model.Organization{})
}

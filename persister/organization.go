package persister

import (
	"context"
	"github.com/c1emon/lemontree/model"
	"gorm.io/gorm"
)

// check
var _ model.OrganizationRepository = &DefaultOrganizationRepository{}

type DefaultOrganizationRepository struct {
	db *gorm.DB
}

func NewDefaultOrganizationRepository() *DefaultOrganizationRepository {
	r := &DefaultOrganizationRepository{
		db: GetDB(),
	}
	r.InitDB()
	return r
}

func (r *DefaultOrganizationRepository) AddDepartment(ctx context.Context, department model.Department) error {
	return nil
}

func (r *DefaultOrganizationRepository) CreateOne(ctx context.Context, org *model.Organization) (*model.Organization, error) {

	return org, r.db.Create(org).Error
}

func (r *DefaultOrganizationRepository) GetOneById(ctx context.Context, id string) (*model.Organization, error) {

	return nil, nil
}

func (r *DefaultOrganizationRepository) UpdateOneById(ctx context.Context, id string, org *model.Organization) (*model.Organization, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DefaultOrganizationRepository) DeleteOneById(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *DefaultOrganizationRepository) InitDB() error {
	return r.db.AutoMigrate(&model.Organization{})
}

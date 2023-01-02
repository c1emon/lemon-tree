package persister

import (
	"context"
	"github.com/c1emon/lemontree/model"
	"github.com/jmoiron/sqlx"
)

// check
var _ model.OrganizationRepository = &DefaultOrganizationRepository{}

type DefaultOrganizationRepository struct {
	db *sqlx.DB
}

func (r *DefaultOrganizationRepository) AddDepartment(ctx context.Context, department model.Department) error {
	return nil
}

func (r *DefaultOrganizationRepository) CreateOne(ctx context.Context, org model.Organization) (model.Organization, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DefaultOrganizationRepository) GetOneById(ctx context.Context, id string) (model.Organization, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DefaultOrganizationRepository) UpdateOneById(ctx context.Context, id string, org model.Organization) (model.Organization, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DefaultOrganizationRepository) DeleteOneById(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

package persister

import (
	"context"
	"github.com/c1emon/lemontree/model"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

// check
var _ model.OrganizationRepository = &DefaultOrganizationRepository{}

type DefaultOrganizationRepository struct {
	db *sqlx.DB
}

func NewDefaultOrganizationRepository() *DefaultOrganizationRepository {
	return &DefaultOrganizationRepository{db: GetDB()}
}

func (r *DefaultOrganizationRepository) AddDepartment(ctx context.Context, department model.Department) error {
	return nil
}

func (r *DefaultOrganizationRepository) CreateOne(ctx context.Context, org model.Organization) (*model.Organization, error) {
	sql, args := sqlbuilder.NewInsertBuilder().InsertInto(r.TableName()).
		Cols("id", "name", "times").
		Values(1, "test", 1234567890).
		Build()
	_, err := db.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *DefaultOrganizationRepository) GetOneById(ctx context.Context, id string) (*model.Organization, error) {

	return nil, nil
}

func (r *DefaultOrganizationRepository) UpdateOneById(ctx context.Context, id string, org model.Organization) (*model.Organization, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DefaultOrganizationRepository) DeleteOneById(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *DefaultOrganizationRepository) TableName() string {
	return ""
}

func (r *DefaultOrganizationRepository) TableSchema() string {
	return `
CREATE TABLE IF NOT EXISTS organization (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    eid CHAR(10) UNIQUE,
    name VARCHAR(255) UNIQUE
    )ENGINE=InnoDB DEFAULT CHARSET=utf8;
`
}

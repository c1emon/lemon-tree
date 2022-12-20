package persister

import (
	"context"
	"github.com/c1emon/lemontree/ent"
	"github.com/c1emon/lemontree/ent/organization"
)

type EntOrganizationRepository struct {
	client *ent.Client
}

func (r *EntOrganizationRepository) SetClient(c any) {
	r.client = c.(*ent.Client)
}

func (r *EntOrganizationRepository) Create(ctx context.Context, o ent.Organization) (*ent.Organization, error) {

	return r.client.Organization.Create().SetName(o.Name).Save(ctx)
}

func (r *EntOrganizationRepository) Delete(ctx context.Context, id string) error {

	_, err := r.client.Organization.Query().Where(organization.ExternalIDEQ(id)).Only(ctx)
	if err != nil {
		return err
	}
	_, err = r.client.Organization.Delete().Where(organization.ExternalIDEQ(id)).Exec(ctx)

	return err
}

func (r *EntOrganizationRepository) Update(ctx context.Context, id string, o ent.Organization) (*ent.Organization, error) {

	err := r.client.Organization.Update().Where(organization.ExternalIDEQ(id)).SetName(o.Name).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return r.client.Organization.Query().Where(organization.ExternalIDEQ(id)).Only(ctx)
}

func (r *EntOrganizationRepository) GetOneById(ctx context.Context, id string) (*ent.Organization, error) {
	return r.client.Organization.Query().Where(organization.ExternalIDEQ(id)).Only(ctx)
}

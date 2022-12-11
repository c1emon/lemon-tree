package service

import (
	"context"
	"fmt"
	"github.com/c1emon/lemontree/dao"
	"github.com/c1emon/lemontree/ent"
	"github.com/c1emon/lemontree/ent/organization"
	"github.com/c1emon/lemontree/util"
)

type OrganizationService struct {
	client *ent.Client
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{client: dao.GetEntClient()}
}

func (s *OrganizationService) CreateOrganization() {

}

func (s *OrganizationService) GetIdByExternalId(eid string) (int64, error) {
	o, err := s.client.Organization.Query().Where(organization.ExternalIDEQ(eid)).Only(context.Background())
	if err != nil {
		switch e := err.(type) {
		case *ent.NotFoundError:
			fmt.Printf("%s", e.Error())
			// return my error

			return -1, util.LemonNotFoundError("ss")
		case *ent.NotSingularError:
			return -1, nil
		}
	}
	return o.ID, nil
}

func (s *OrganizationService) GetExternalIdById(id int64) string {
	o, _ := s.client.Organization.Query().Where(organization.IDEQ(id)).Only(context.Background())
	return o.ExternalID
}

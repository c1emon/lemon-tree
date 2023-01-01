package service

import (
	"context"
	"github.com/c1emon/lemontree/model"
	"github.com/c1emon/lemontree/persister"
)

type OrganizationService struct {
	r model.OrganizationRepository
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{
		r: &persister.DefaultOrganizationRepository{},
	}
}

func (s *OrganizationService) CreateOrganization() {
	s.r.CreateOne(context.Background(), model.Organization{})
}

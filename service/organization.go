package service

import (
	"github.com/c1emon/lemontree/model"
	"github.com/c1emon/lemontree/persister"
)

type OrganizationService struct {
	r model.OrganizationRepository
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{
		r: &persister.GormOrganizationRepository{},
	}
}

func (s *OrganizationService) CreateOrganization() {
}

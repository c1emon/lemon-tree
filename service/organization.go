package service

import (
	"context"
	"fmt"
	"github.com/c1emon/lemontree/model"
	"github.com/c1emon/lemontree/persister"
)

type OrganizationService struct {
	organizationRepository model.OrganizationRepository
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{
		organizationRepository: &persister.GormOrganizationRepository{},
	}
}

func (s *OrganizationService) CreateOrganization() {
	org, err := s.organizationRepository.CreateOne(context.Background(), nil)
	if err != nil {

	}
	fmt.Println(org)

}

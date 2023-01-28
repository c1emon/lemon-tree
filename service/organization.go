package service

import "github.com/c1emon/lemontree/repository"

type OrganizationService struct {
	repository repository.DefaultOrganizationRepository
}

func (s *OrganizationService) Create() error {
	//s.repository.CreateOne()
	return nil
}

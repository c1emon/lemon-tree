package service

import (
	"context"
	"github.com/c1emon/lemontree/httpx"
	"github.com/c1emon/lemontree/model"
	"github.com/c1emon/lemontree/repository"
	"github.com/pkg/errors"
)

type OrganizationService struct {
	repository repository.DefaultOrganizationRepository
}

func (s *OrganizationService) Create(name string) (*model.Organization, error) {
	err := s.repository.CreateOne(context.Background(), &model.Organization{Name: name})
	if err != nil {
		return nil, errors.WithMessage(err, "name "+name)
	}
	org, err := s.repository.GetOneByName(context.Background(), name)
	return org, errors.WithMessage(err, "name "+name)
}

func (s *OrganizationService) GetById(id string) (*model.Organization, error) {

	org, err := s.repository.GetOneById(context.Background(), id)
	return org, errors.WithMessage(err, "organization")
}

func (s *OrganizationService) GetByName(name string) (*model.Organization, error) {

	org, err := s.repository.GetOneByName(context.Background(), name)
	return org, errors.WithMessage(err, "organization")
}

func (s *OrganizationService) GetByNameAll(pageable httpx.Pageable, name string) []model.Organization {

	return s.repository.GetAllByName(context.Background(), pageable, name)
}

func (s *OrganizationService) DeleteOne(id string) error {
	return errors.WithMessage(s.repository.DeleteOneById(context.Background(), id), "organization")
}

func (s *OrganizationService) UpdateOne() error {
	return nil
}

func (s *OrganizationService) FindAll() {

}

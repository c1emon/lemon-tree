package org

import (
	"context"

	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/c1emon/lemontree/pkg/httpx"
	"github.com/pkg/errors"
)

func NewOrganizationService() *OrganizationService {
	repo := &gormOrganizationRepository{
		db: gormx.GetGormDB(),
	}
	repo.InitDB()
	return &OrganizationService{repository: repo}
}

type OrganizationService struct {
	repository OrganizationRepository
}

func (s *OrganizationService) Create(name string) (*Organization, error) {
	err := s.repository.CreateOne(context.Background(), &Organization{Name: name})
	if err != nil {
		return nil, errors.WithMessage(err, "name "+name)
	}
	org, err := s.repository.GetOneByName(context.Background(), name)
	return org, errors.WithMessage(err, "name "+name)
}

func (s *OrganizationService) GetById(id string) (*Organization, error) {

	org, err := s.repository.GetOneById(context.Background(), id)
	return org, errors.WithMessage(err, "organization")
}

func (s *OrganizationService) GetByName(name string) (*Organization, error) {

	org, err := s.repository.GetOneByName(context.Background(), name)
	return org, errors.WithMessage(err, "organization")
}

func (s *OrganizationService) GetByNameAll(pageable httpx.Pageable, name string) []Organization {

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

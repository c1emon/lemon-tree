package service

import (
	"github.com/c1emon/lemontree/dao"
	"github.com/c1emon/lemontree/ent"
)

type IdpService struct {
	client *ent.Client
}

func NewIdpService() *IdpService {
	return &IdpService{client: dao.GetEntClient()}
}

func (s *IdpService) CreateIdp() {

}

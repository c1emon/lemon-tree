package service

import (
	"github.com/c1emon/lemontree/ent"
	"github.com/c1emon/lemontree/persister"
)

type IdpService struct {
	client *ent.Client
}

func NewIdpService() *IdpService {
	return &IdpService{client: persister.GetEntClient()}
}

func (s *IdpService) CreateIdp() {

}

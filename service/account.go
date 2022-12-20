package service

import (
	"context"
	"github.com/c1emon/lemontree/ent"
	"github.com/c1emon/lemontree/ent/organization"
	"github.com/c1emon/lemontree/ent/staff"
	"github.com/c1emon/lemontree/persister"
	"github.com/c1emon/lemontree/util"
)

type AccountService struct {
	client *ent.Client
}

func NewAccountService() *AccountService {
	return &AccountService{client: persister.GetEntClient()}
}

func (s *AccountService) CreateAccount(oid int64, username, password string) {

	p, _ := util.GetHashedPasswd(password)

	s.client.Staff.Create().
		SetUsername(username).
		SetPassword(p).
		SetOrganizationID(oid).
		Save(context.Background())

}

func (s *AccountService) CheckPasswd(oid int64, username, password string) bool {
	a, _ := s.client.Organization.
		Query().Where(organization.ID(oid)).
		QueryStaffs().Where(staff.UsernameEQ(username)).Only(context.Background())
	return util.CheckPasswd(a.Password, password)
}

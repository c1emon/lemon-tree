package user

import (
	"context"
	"encoding/json"

	"github.com/c1emon/lemontree/pkg/logx"
	"github.com/c1emon/lemontree/pkg/persister"
	"gorm.io/datatypes"
)

func NewUserService() *UserService {
	repo := &gormUserRepository{
		db: persister.GetDB(),
	}
	repo.InitDB()
	return &UserService{repository: repo}
}

type UserService struct {
	repository UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, u *User) {
	err := s.repository.CreateOne(ctx, u)
	if err != nil {
		logx.GetLogger().Errorf("failed create user, %s", err)
	}
}

func (s *UserService) FindByOidAndName(oid string, name string) (*User, error) {
	return s.repository.FindByOidAndName(oid, name)
}

func NewUserIdentityService() *UserIdentityService {
	repo := &gormUserIdentityRepository{
		db: persister.GetDB(),
	}
	repo.InitDB()
	return &UserIdentityService{repository: repo, userSvc: nil}
}

type UserIdentityService struct {
	repository UserIdentityRepository
	userSvc    *UserService
}

func (s *UserIdentityService) CreateIdentity(ctx context.Context, oid, uid, idpId string, identityParam map[string]string) {
	i, _ := json.Marshal(identityParam)
	identity := &UserIdentity{
		Uid:      uid,
		Oid:      oid,
		IdpId:    idpId,
		Identity: i,
	}

	err := s.repository.CreateOne(ctx, identity)
	if err != nil {
		logx.GetLogger().Errorf("failed create identity, %s", err)
	}
}

func (s *UserIdentityService) Validate(ctx context.Context, oid string, builder func(*datatypes.JSONQueryExpression) *datatypes.JSONQueryExpression) (string, bool) {

	uid, err := s.repository.Validate(ctx, oid, builder)
	if err != nil {
		return "", false
	}
	return uid, true
}

package user

import (
	"context"
	"encoding/json"

	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/c1emon/lemontree/pkg/logx"
	"gorm.io/datatypes"
)

func NewUserService() *UserService {
	userRepo := &gormUserRepository{
		db: gormx.GetGormDB(),
	}
	userRepo.InitDB()

	userIdRepo := &gormUserIdentityRepository{
		db: gormx.GetGormDB(),
	}
	userIdRepo.InitDB()

	return &UserService{userRepo: userRepo, userIdRepo: userIdRepo}
}

type UserService struct {
	userRepo   UserRepository
	userIdRepo UserIdentityRepository
}

func (s *UserService) CreateUser(ctx context.Context, u *User) {
	err := s.userRepo.CreateOne(ctx, u)
	if err != nil {
		logx.GetLogger().Errorf("failed create user, %s", err)
	}
}

func (s *UserService) FindByOidAndName(oid string, name string) (*User, error) {
	return s.userRepo.FindByOidAndName(oid, name)
}

func (s *UserService) CreateIdentity(ctx context.Context, oid, uid, idpId string, identityParam map[string]string) {
	i, _ := json.Marshal(identityParam)
	identity := &UserIdentity{
		Uid:      uid,
		Oid:      oid,
		IdpId:    idpId,
		Identity: i,
	}

	err := s.userIdRepo.CreateOne(ctx, identity)
	if err != nil {
		logx.GetLogger().Errorf("failed create identity, %s", err)
	}
}

func (s *UserService) Validate(ctx context.Context, oid string, builder func(*datatypes.JSONQueryExpression) *datatypes.JSONQueryExpression) (string, bool) {

	uid, err := s.userIdRepo.Validate(ctx, oid, builder)
	if err != nil {
		return "", false
	}
	return uid, true
}

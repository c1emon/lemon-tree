package idp

import (
	"context"
	"encoding/json"

	"github.com/c1emon/lemontree/internal/user"
	"github.com/c1emon/lemontree/pkg/logx"
	"gorm.io/datatypes"
)

var _ IDProvider = &PasswdIDP{}

// var _ user.ValidaterBuilder = &PasswdIDP{}

type PasswdIDP struct {
	userService         *user.UserService
	userIdentityService *user.UserIdentityService
}

func NewPasswdIDP(userSvc *user.UserService, identitySvc *user.UserIdentityService, config any) *PasswdIDP {
	// user idps
	return &PasswdIDP{
		userService:         userSvc,
		userIdentityService: identitySvc,
	}
}

// Suooprt implements IDProvider.
func (*PasswdIDP) Support(provider string) bool {
	return provider == "passwd"
}

// GetUserInfo implements IDProvider.
func (p *PasswdIDP) GetUser(ctx context.Context, data any) (*user.User, error) {
	// getConfig
	param := struct {
		Oid      string `json:"oid"`
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
	}{}
	if b, ok := data.(json.RawMessage); ok {
		json.Unmarshal(b, &param)
	}
	param.Oid = "clcgbaky00000ze5jztbggr8b"

	uid, ok := p.userIdentityService.Validate(ctx, param.Oid, func() func(*datatypes.JSONQueryExpression) *datatypes.JSONQueryExpression {

		return func(builder *datatypes.JSONQueryExpression) *datatypes.JSONQueryExpression {

			return builder.Equals("username", param.Username)
		}
	}())
	if !ok {
		logx.GetLogger().Infof("login fail for %s", param.Username)
	}

	logx.GetLogger().Infof("login success for %s", uid)

	return nil, nil
}

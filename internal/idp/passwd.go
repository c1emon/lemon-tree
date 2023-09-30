package idp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/c1emon/lemontree/internal/user"
	"github.com/c1emon/lemontree/pkg/logx"
	"gorm.io/datatypes"
)

var _ IDProvider = &PasswdIDP{}

type PasswdIDP struct {
	userService *user.UserService
}

func NewPasswdIDP(userSvc *user.UserService) *PasswdIDP {
	// user idps
	return &PasswdIDP{
		userService: userSvc,
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

	uid, ok := p.userService.Validate(ctx, param.Oid, func(builder func() *datatypes.JSONQueryExpression) []any {
		var exps []any
		exps = append(exps, builder().Equals(param.Username, "username"))
		exps = append(exps, builder().Equals(param.Passwd, "passwd"))
		return exps
	})
	if !ok {
		logx.GetLogger().Infof("login fail for %s", param.Username)
		return nil, fmt.Errorf("login fail for %s", param.Username)
	}

	logx.GetLogger().Infof("login success for %s", uid)

	return p.userService.FindUser(ctx, uid)
}

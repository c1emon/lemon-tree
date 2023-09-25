package user

import (
	"context"

	"github.com/c1emon/lemontree/pkg/gormx"
	"gorm.io/datatypes"
)

type User struct {
	gormx.BaseFields
	UserName string `json:"username" gorm:"column:name;type:varchar(256);uniqueIndex:udx_org_user"`
	Name     string `json:"name" gorm:"column:name;type:varchar(256);uniqueIndex:udx_org_name"`
	Oid      string `json:"oid" gorm:"column:oid;type:varchar(256);uniqueIndex:udx_org_user;uniqueIndex:udx_org_name"`
}

func (User) TableName() string {
	return "users"
}

type UserRepository interface {
	gormx.BaseRepository[User]
	FindByOidAndName(string, string) (*User, error)
}

type UserIdentity struct {
	gormx.BaseFields
	Uid      string         `json:"uid" gorm:"column:uid;type:varchar(256)"`
	Oid      string         `json:"oid" gorm:"column:uid;type:varchar(256)"`
	IdpId    string         `json:"idp_id" gorm:"column:idp_id;type:varchar(256)"`
	Identity datatypes.JSON `json:"identity" gorm:"column:identity;not null"`
}

func (UserIdentity) TableName() string {
	return "user_identities"
}

type UserIdentityRepository interface {
	gormx.BaseRepository[UserIdentity]
	FindByUidAndIdpId(context.Context, string, string) (*UserIdentity, error)
	Validate(context.Context, string, func(*datatypes.JSONQueryExpression) *datatypes.JSONQueryExpression) (string, error)
}

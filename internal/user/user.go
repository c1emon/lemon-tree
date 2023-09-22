package user

import "github.com/c1emon/lemontree/pkg/gormx"

type User struct {
	gormx.BaseFields
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type UserRepository interface {
	gormx.BaseRepository[User]
}

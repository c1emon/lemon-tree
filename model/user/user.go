package model

import "github.com/c1emon/lemontree/model"

type User struct {
	model.BaseFields
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type UserRepository interface {
	model.BaseRepository[User]
}

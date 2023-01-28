package model

import "github.com/c1emon/lemontree/repository"

type User struct {
	BaseField
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type UserRepository interface {
	repository.BaseRepository[User]
}

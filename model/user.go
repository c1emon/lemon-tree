package model

type User struct {
	BaseFields
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type UserRepository interface {
	BaseRepository[User]
}

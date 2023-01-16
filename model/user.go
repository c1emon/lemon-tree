package model

type User struct {
	BaseField
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type UserRepository interface {
	BaseRepository[User]
}

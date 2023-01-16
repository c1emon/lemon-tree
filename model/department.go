package model

type Department struct {
	BaseField
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type DepartmentRepository interface {
	BaseRepository[Department]
}

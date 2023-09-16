package model

type Department struct {
	BaseFields
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type DepartmentRepository interface {
	BaseRepository[Department]
}

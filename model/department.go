package model

import "github.com/c1emon/lemontree/repository"

type Department struct {
	BaseField
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type DepartmentRepository interface {
	repository.BaseRepository[Department]
}

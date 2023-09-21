package org

import "github.com/c1emon/lemontree/model"

type Department struct {
	model.BaseFields
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type DepartmentRepository interface {
	model.BaseRepository[Department]
}

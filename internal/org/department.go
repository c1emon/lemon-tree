package org

import "github.com/c1emon/lemontree/pkg/gormx"

type Department struct {
	gormx.BaseFields
	Name           string `json:"name"`
	OrganizationId string `json:"oid"`
}

type DepartmentRepository interface {
	gormx.BaseRepository[Department]
}

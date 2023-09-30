package org

import (
	"context"

	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/c1emon/lemontree/pkg/httpx"
)

type Organization struct {
	gormx.BaseFields
	Name string `json:"name" gorm:"column:name;type:varchar(256);uniqueIndex;not null"`
}

func (Organization) TableName() string {
	return "organizations"
}

type OrganizationRepository interface {
	gormx.BaseRepository[Organization]
	GetOneByName(context.Context, string) (*Organization, error)
	GetAllByName(context.Context, httpx.Pageable, string) []Organization
}

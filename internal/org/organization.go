package org

import "github.com/c1emon/lemontree/pkg/gormx"

type Organization struct {
	gormx.BaseFields
	Name string `json:"name" gorm:"column:name;type:varchar(256);uniqueIndex;not null"`
}

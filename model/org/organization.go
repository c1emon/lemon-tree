package org

import "github.com/c1emon/lemontree/model"

type Organization struct {
	model.BaseFields
	Name string `json:"name" gorm:"column:name;type:varchar(256);uniqueIndex;not null"`
}

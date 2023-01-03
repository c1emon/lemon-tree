package model

import "context"

type Organization struct {
	BaseField
	Name string `json:"name" gorm:"column:name;type:varchar(256);uniqueIndex;not null"`
}

type OrganizationRepository interface {
	BaseRepository[Organization]
	AddDepartment(context.Context, Department) error
}

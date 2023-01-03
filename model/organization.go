package model

import "context"

type Organization struct {
	BaseField
	Name string `json:"name" db:"name"`
}

type OrganizationRepository interface {
	BaseRepository[Organization]
	AddDepartment(context.Context, Department) error
}

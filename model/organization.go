package model

import "context"

type Organization struct {
	BaseField
	Name string `json:"name,omitempty"`
}

type OrganizationRepository interface {
	BaseRepository[Organization]
	AddDepartment(context.Context, Department) error
}

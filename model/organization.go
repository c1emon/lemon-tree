package model

import "context"

type Organization struct {
	BaseField
	Name string
}

type OrganizationRepository interface {
	AddDepartment(context.Context, Department) error
}

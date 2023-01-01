package model

type Department struct {
	BaseField
	OrganizationId string
	Name           string
}

func NewDepartment() *Department {
	return &Department{BaseField: BaseField{}}
}

func CreateDepartment() *Department {
	return &Department{BaseField: CreateBaseField()}
}

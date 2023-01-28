package model

type Organization struct {
	BaseField
	Name string `json:"name" gorm:"column:name;type:varchar(256);uniqueIndex;not null"`
}

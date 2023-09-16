package model

type Organization struct {
	BaseFields
	Name string `json:"name" gorm:"column:name;type:varchar(256);uniqueIndex;not null"`
}

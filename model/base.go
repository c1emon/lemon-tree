package model

import (
	"context"
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
	"time"
)

type BaseField struct {
	Id         string    `json:"id" gorm:"column:id;type:char(25);primaryKey;<-:create"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;<-:create;autoUpdateTime"`
}

func (f *BaseField) BeforeCreate(tx *gorm.DB) (err error) {
	f.Id = cuid.New()
	return nil
}

type MetaField struct {
}

func CreateBaseField() BaseField {
	return BaseField{Id: cuid.New()}
}

func (f *BaseField) GetId() string {
	return f.Id
}

func (f *BaseField) GetCreatedTime() time.Time {
	return f.CreateTime
}

func (f *BaseField) GetUpdatedTime() time.Time {
	return f.UpdateTime
}

func (f *BaseField) SetUpdatedTime(t time.Time) {
	f.UpdateTime = t
}

type BaseRepository[T any] interface {
	CreateOne(context.Context, *T) (*T, error)
	GetOneById(context.Context, string) (*T, error)
	UpdateOneById(context.Context, string, *T) (*T, error)
	DeleteOneById(context.Context, string) error
	InitDB() error
}

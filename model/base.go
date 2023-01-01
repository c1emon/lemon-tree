package model

import (
	"context"
	"github.com/lucsky/cuid"
	"time"
)

type BaseField struct {
	id         int64
	EId        string    `json:"id" db:"eid"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

func CreateBaseField() BaseField {
	n := time.Now()
	return BaseField{EId: cuid.New(), CreateTime: n, UpdateTime: n}
}

func (f *BaseField) GetEId() string {
	return f.EId
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
	CreateOne(context.Context, T) (T, error)
	GetOneById(context.Context, string) (T, error)
	UpdateOneById(context.Context, string, T) (T, error)
	DeleteOneById(context.Context, string) error
}

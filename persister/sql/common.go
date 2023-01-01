package sql

import "github.com/huandu/go-sqlbuilder"

func GetById(id string) string {
	return sqlbuilder.NewSelectBuilder().Select("*").From("").String()
}

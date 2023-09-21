package oidc

import (
	"github.com/c1emon/lemontree/model"
	"github.com/c1emon/lemontree/model/oidc"
)

type SigninRequest struct {
	model.BaseFields
	oidc.Aa
}

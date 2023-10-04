package token

import "github.com/c1emon/lemontree/pkg/cachex"

func NewTokenService(cacher cachex.Cacher) *TokenService {
	return &TokenService{cacher: cacher}
}

type TokenService struct {
	cacher cachex.Cacher
}

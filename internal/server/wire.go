package server

import (
	"github.com/c1emon/lemontree/internal/setting"
)

// var wireExtsSet = wire.NewSet(
// 	setting.New,
// 	ProvideHttpServer,
// )

func Initialize() (*Server, error) {
	s := setting.New()
	srv, _ := ProvideHttpServer(s)
	return New(s, srv)
}

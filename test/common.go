package test

import (
	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/c1emon/lemontree/pkg/logx"
)

func start() {
	logx.Init("debug")
}

func stop() {
	if err := gormx.DisConnect(); err != nil {
		logx.GetLogger().Warnf("unable close db: %s", err)
	}
}

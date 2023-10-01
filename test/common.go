package test

import (
	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/c1emon/lemontree/pkg/logx"
)

func start() {
	logx.Init("debug")
	gormx.Initialize("postgres", "host=10.10.0.70 port=5432 user=postgres dbname=lemon_tree password=123456")
}

func stop() {
	if err := gormx.DisConnect(); err != nil {
		logx.GetLogger().Warnf("unable close db: %s", err)
	}
}

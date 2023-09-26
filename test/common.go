package test

import (
	"github.com/c1emon/lemontree/pkg/logx"
	"github.com/c1emon/lemontree/pkg/persister"
)

func start() {
	logx.Init("debug")
}

func stop() {
	if err := persister.DisConnect(); err != nil {
		logx.GetLogger().Warnf("unable close db: %s", err)
	}
}

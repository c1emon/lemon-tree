package test

import (
	"github.com/c1emon/lemontree/config"
	"github.com/c1emon/lemontree/log"
	"github.com/c1emon/lemontree/persister"
)

func start() {
	config.SetConfig(8080, "postgres", "host=10.0.0.70 port=5432 user=postgres dbname=lemon_tree password=123456 sslmode=disable")
	log.Init("debug")
}

func stop() {
	if err := persister.DisConnect(); err != nil {
		log.GetLogger().Warnf("unable close db: %s", err)
	}
}

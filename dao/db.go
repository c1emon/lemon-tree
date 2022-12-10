package dao

import (
	"context"
	"github.com/c1emon/lemontree/config"
	"github.com/c1emon/lemontree/ent"
	"github.com/c1emon/lemontree/log"
	"sync"

	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}
var c *ent.Client

func GetEntClient() *ent.Client {
	logger := log.GetLogger()
	lock.Lock()
	defer lock.Unlock()
	if c == nil {
		logger.Debug("connect db")
		conf := config.GetConfig()
		var err error
		c, err = ent.Open(conf.DbDriver, conf.DbSource)
		if err != nil {
			if err := c.Close(); err != nil {
				logger.Warnf("unable close db client: %s", err)
			}
			logger.Panicf("failed connect db: %s", err)
		}

		if err := c.Schema.Create(context.Background()); err != nil {
			if err := c.Close(); err != nil {
				logger.Warnf("unable close db client: %s", err)
			}
			logger.Fatalf("failed create schema resources: %s", err)
		}

	}

	return c
}

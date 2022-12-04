package dao

import (
	"context"
	"github.com/c1emon/lemontree/config"
	"github.com/c1emon/lemontree/ent"
	"github.com/c1emon/lemontree/log"
	"sync"
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
			logger.Panic("failed connect db")
		}

		if err := c.Schema.Create(context.Background()); err != nil {
			logger.Fatalf("failed create schema resources: %v", err)
		}

	}

	return c
}

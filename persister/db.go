package persister

import (
	"github.com/c1emon/lemontree/config"
	"github.com/c1emon/lemontree/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sync"
)

var lock = &sync.Mutex{}

var db *sqlx.DB

func Connect(driverName, dataSourceName string) *sqlx.DB {
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		log.GetLogger().Panicf("unable connect to %s: %s", driverName, err)
	}
	return db
}

func GetDB() *sqlx.DB {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		c := config.GetConfig()
		db = Connect(c.DbDriver, c.DbSource)
	}
	return db
}

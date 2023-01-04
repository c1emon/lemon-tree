package persister

import (
	"github.com/c1emon/lemontree/config"
	"github.com/c1emon/lemontree/log"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"sync"
)

var once = sync.Once{}
var DB *gorm.DB

type DriverType int

const (
	Unknown DriverType = iota - 1
	Postgres
	Mysql
	Sqlite
)

func (d DriverType) String() string {
	switch d {
	case Postgres:
		return "postgres"
	case Mysql:
		return "mysql"
	case Sqlite:
		return "sqlite"
	default:
		return "unknown"
	}
}

func ParseDriverType(dt string) DriverType {
	switch strings.ToLower(dt) {
	case "postgres":
		return Postgres
	case "mysql":
		return Mysql
	case "sqlite":
		return Sqlite
	default:
		return Unknown
	}
}

func Connect(driverName DriverType, dsn string) {
	var dialector gorm.Dialector
	switch driverName {
	case Postgres:
		dialector = postgres.Open(dsn)
	case Mysql:
		dialector = mysql.Open(dsn)
	case Sqlite:
		dialector = sqlite.Open(dsn)
	case Unknown:
	default:
		log.GetLogger().Panicf("unknown driver type: %s", driverName)
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: log.GetGormLogrusLogger(),
	})

	if err != nil {
		log.GetLogger().Panicf("unable connect to %s: %s", driverName, err)
	}
	DB = db
}

func GetDB() *gorm.DB {

	once.Do(func() {
		c := config.GetConfig()
		Connect(ParseDriverType(c.DbDriver), c.DbSource)
	})

	return DB
}

func DisConnect() error {
	if DB != nil {
		d, err := DB.DB()
		if err != nil {
			return nil
		}
		return d.Close()
	}
	return nil
}

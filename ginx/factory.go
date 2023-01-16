package ginx

import (
	"github.com/c1emon/lemontree/log"
	"github.com/gin-gonic/gin"
	"sync"
)

var e *gin.Engine
var once = &sync.Once{}

func SingletonGinFactory() *gin.Engine {

	once.Do(func() {

		g := gin.New()
		g.Use(LogrusLogger(log.GetLogger()))
		e = g
	})

	return e
}

func GetGinEngine() *gin.Engine {
	return SingletonGinFactory()
}

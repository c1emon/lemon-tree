package ginx

import (
	"sync"

	"github.com/c1emon/lemontree/pkg/logx"
	"github.com/gin-gonic/gin"
)

var e *gin.Engine
var once = &sync.Once{}

func SingletonGinFactory() *gin.Engine {

	once.Do(func() {

		g := gin.New()

		g.Use(LogrusLogger(logx.GetLogger()), ErrorHandler(), Recovery(logx.GetLogger()))

		e = g
	})

	return e
}

func GetGinEngine() *gin.Engine {
	return SingletonGinFactory()
}

// TODO: split
//type GinMiddleware interface {
//	HandleRequest()
//	HandleResponse()
//}

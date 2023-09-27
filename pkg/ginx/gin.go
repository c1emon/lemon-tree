package ginx

import (
	"sync"

	"github.com/c1emon/lemontree/pkg/logx"
	"github.com/gin-gonic/gin"
)

var e *gin.Engine
var once = &sync.Once{}

func singletonGinFactory() *gin.Engine {

	once.Do(func() {
		mode := gin.DebugMode
		gin.SetMode(mode)
		g := gin.New()

		g.Use(LogrusLogger(logx.GetLogger()), ErrorHandler(), Recovery(logx.GetLogger()))

		e = g
	})

	return e
}

func GetGinEng() *gin.Engine {
	return singletonGinFactory()
}

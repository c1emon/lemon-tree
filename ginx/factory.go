package ginx

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var e *gin.Engine
var once = &sync.Once{}

func SingletonGinFactory() *gin.Engine {

	once.Do(func() {

		gin.Logger()

		g := gin.New()
		g.Use()

		e = g
	})

	return e
}

func GetGinEngine() *gin.Engine {
	return e
}

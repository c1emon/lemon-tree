package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogrusLogger(logger *logrus.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		// Start timer
		//start := time.Now()
		method := c.Request.Method
		uri := c.Request.RequestURI

		// Process request
		c.Next()
		//latency := time.Now().Sub(start)
		status := c.Writer.Status()
		logger.Infof("[%s %d] %s", method, status, uri)

	}
}

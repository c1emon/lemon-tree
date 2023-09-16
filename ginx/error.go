package ginx

import (
	"github.com/c1emon/lemontree/errorc"
	http2 "github.com/c1emon/lemontree/httpx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Process request
		c.Next()
		for _, e := range c.Errors {
			if err, ok := e.Err.(errorc.ErrorX); ok {
				c.JSON(err.HttpStatus(), http2.NewResponse(err.Code()).WithError(err.Error()))
			} else {
				c.JSON(http.StatusBadRequest, http2.NewResponse(1001).WithError(e.Error()))
			}
			return
		}

	}
}

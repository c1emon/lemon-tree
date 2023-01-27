package ginx

import (
	"github.com/c1emon/lemontree/controller"
	"github.com/c1emon/lemontree/errorc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Process request
		c.Next()
		for _, e := range c.Errors {
			if err, ok := e.Err.(errorc.ErrorX); ok {
				c.JSON(err.HttpStatus(), controller.NewResponse(err.Code()).WithError(err.Error()))
			} else {
				c.JSON(http.StatusBadRequest, controller.NewResponse(1001).WithError(e.Error()))
			}
			return
		}

	}
}

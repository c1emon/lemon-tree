package controller

import (
	"github.com/c1emon/lemontree/errorc"
	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	if err, ok := err.(errorc.Error); ok {

		err := c.JSON(err.HttpStatus(), FromError(err))
		if err != nil {
			c.Logger().Warnf("")
		}
		return
	}
	c.Logger().Warnf("")

	return
}

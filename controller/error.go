package controller

import (
	"github.com/c1emon/lemontree/errorc"
	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(errRaw error, c echo.Context) {
	var err errorc.ErrorX

	defer func() {
		if e := c.JSON(err.HttpStatus(), NewResponse(err.Code()).WithError(errRaw.Error())); e != nil {
			c.Logger().Warnf("send error response failed: %s", e)
		}
	}()

	if e, ok := errRaw.(errorc.ErrorX); ok {
		c.Logger().Infof("checked error: %v", e)
		err = e
	} else {
		c.Logger().Warnf("unchecked error: %+v", errRaw)
		err = errorc.ErrUnknown
	}

}

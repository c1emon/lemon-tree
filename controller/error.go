package controller

import (
	"fmt"
	"github.com/c1emon/lemontree/errorc"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func HTTPErrorHandler(errRaw error, c echo.Context) {
	var err errorc.ErrorX

	defer func() {
		if e := c.JSON(err.HttpStatus(), NewResponse(err.Code()).WithError(errRaw.Error())); e != nil {
			c.Logger().Warnf("send error response failed: %s", e)
		}
	}()

	route := fmt.Sprintf("[%s] %s", c.Request().Method, c.Request().URL)

	if e, ok := errors.Cause(errRaw).(errorc.ErrorX); ok {
		c.Logger().Infof("checked error of %s:\n	%s", route, errRaw)
		err = e
	} else {
		c.Logger().Warnf("unchecked error %s:\n	%+v", route, errRaw)
		err = errorc.ErrUnknown
	}

}

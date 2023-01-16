package test

import (
	"github.com/c1emon/lemontree/controller"
	"github.com/c1emon/lemontree/errorc"
	"github.com/c1emon/lemontree/errorc/parser"
	"github.com/c1emon/lemontree/log"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"testing"
)

func Check() error {
	return errors.WithStack(errorc.From(gorm.ErrRecordNotFound))
}

func Test_Err(t *testing.T) {
	log.Init("debug")
	parser.NewGormParser()
	e := controller.SingletonEchoFactory()
	e.HTTPErrorHandler = controller.HTTPErrorHandler
	e.GET("/test", func(c echo.Context) error {

		return errors.WithMessage(Check(), "get org by username clemon")
	})

	e.Start(":8080")
}

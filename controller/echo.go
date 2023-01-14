package controller

import (
	"github.com/c1emon/lemontree/doauth"
	"github.com/c1emon/lemontree/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"sync"
)

var e *echo.Echo
var once = &sync.Once{}

func SingletonEchoFactory() *echo.Echo {

	once.Do(func() {
		e = echo.New()

		e.Logger = log.GetEchoLogrusLogger()

		e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:    true,
			LogMethod: true,
			LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
				e.Logger.Infof("[%s] %s%s", values.Method, values.URI)
				return nil
			},
		}))
	})

	return e
}

func HttpWrapperWithError(h doauth.HandlerFuncWithError) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.Response().Writer, c.Request())
	}
}

func HttpWrapper(h http.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h(c.Response().Writer, c.Request())
		return nil
	}
}

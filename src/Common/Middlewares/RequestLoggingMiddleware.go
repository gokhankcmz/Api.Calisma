package Middlewares

import (
	"Api.Calisma/src/Common/Logging/LogBody"
	"Api.Calisma/src/Common/Logging/LogrusAdapter"
	"github.com/labstack/echo/v4"
)

func RequestLoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			le := LogBody.GetDynamicLog().AddTimeInfo().AddRequestInfo(c)
			go LogrusAdapter.Create(le).LogInfo()
			return next(c)
		}
	}
}

func UseLogrusRequestLogging(e *echo.Echo) {
	e.Use(RequestLoggingMiddleware())
}

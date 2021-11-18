package Logger

import (
	"github.com/labstack/echo/v4"
)

func RequestLoggingMiddleware(applicationName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			go StartBuildingLogrus().AddRequestInfo(c).AddApplicationInfo(applicationName).FinishBuilding().Info()
			return next(c)
		}
	}
}

func UseLogrusRequestLogging(e *echo.Echo, applicationName string){
	e.Use(RequestLoggingMiddleware(applicationName))
}


package Middlewares

import (
	"Api.Calisma/src/Common/Logger"
	"Api.Calisma/src/Common/Models/ErrorModels"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PanicHandlerMiddleware(applicationName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					switch v := err.(type) {
					case *ErrorModels.Error:
						go Logger.StartBuildingLogrus().AddErrorInfo(err).AddRequestInfo(c).AddApplicationInfo(applicationName).FinishBuilding().Info()
						c.JSON(v.StatusCode, v.PublicError)
					case ErrorModels.Error:
						go Logger.StartBuildingLogrus().AddErrorInfo(err).AddRequestInfo(c).AddApplicationInfo(applicationName).FinishBuilding().Info()
						c.JSON(v.StatusCode, v.PublicError)
					default:
						go Logger.StartBuildingLogrus().AddErrorInfo(err).AddRequestInfo(c).AddApplicationInfo(applicationName).FinishBuilding().Error()
						c.NoContent(http.StatusInternalServerError)
						fmt.Println(err)
					}
				}
			}()
			return next(c)
		}
	}
}

func UsePanicHandlerMiddleware(e *echo.Echo, applicationName string) {
	e.Use(PanicHandlerMiddleware(applicationName))
}

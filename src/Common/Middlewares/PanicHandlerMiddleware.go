package Middlewares

import (
	"Api.Calisma/src/Common/Logging/LogBody"
	"Api.Calisma/src/Common/Logging/LogrusAdapter"
	"Api.Calisma/src/Common/Models/ErrorModels"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PanicHandlerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					le := LogBody.GetDynamicLog().AddTimeInfo().AddRequestInfo(c)
					switch v := err.(type) {
					case *ErrorModels.Error:
						go LogrusAdapter.Create(le.AddStruct(*v)).LogInfo()
						c.JSON(v.StatusCode, v.PublicError)
					case ErrorModels.Error:
						go LogrusAdapter.Create(le.AddStruct(v)).LogInfo()
						c.JSON(v.StatusCode, v.PublicError)
					default:
						go LogrusAdapter.Create(le).LogError()
						fmt.Println(err)
						c.JSON(http.StatusInternalServerError, err)
					}
				}
			}()
			return next(c)
		}
	}
}

func UsePanicHandlerMiddleware(e *echo.Echo) {
	e.Use(PanicHandlerMiddleware())
}

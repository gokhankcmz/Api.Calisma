package Middlewares

import (
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Token"
	"github.com/labstack/echo/v4"
)

func TokenValidatorMiddleware(failMsg string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id := c.Param("id")
			claims := Token.ValidateAndGetClaims(c.Request().Header.Get("Authorization"))
			if claims.ID != id {
				panic(ErrorModels.UnauthorizedRequest.SetPublicDetail(failMsg))
			}
			return next(c)
		}
	}
}

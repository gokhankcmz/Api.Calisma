package Controllers

import (
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Redirection"
	"github.com/labstack/echo/v4"
)

func CreateToken(ctx echo.Context) error{
	return Redirection.Redirect(ctx,Constants.CustomerServiceUri)
}


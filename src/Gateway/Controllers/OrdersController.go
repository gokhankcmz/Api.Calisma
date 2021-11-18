package Controllers

import (
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Redirection"
	"github.com/labstack/echo/v4"
)

func GetAllOrders(ctx echo.Context) error{
	return Redirection.Redirect(ctx,Constants.OrderServiceUri)
}

func CreateOrder(ctx echo.Context) error{
	return Redirection.Redirect(ctx,Constants.OrderServiceUri)
}

func DeleteOrder(ctx echo.Context) error {
	return Redirection.Redirect(ctx,Constants.OrderServiceUri)
}

func GetAOrder(ctx echo.Context) error{
	return Redirection.Redirect(ctx,Constants.OrderServiceUri)
}

func UpdateOrder(ctx echo.Context) error{
	return Redirection.Redirect(ctx,Constants.OrderServiceUri)
}


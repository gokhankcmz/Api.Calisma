package Routing

import (
	Controller "Api.Calisma/src/OrderService/Controller"
	"github.com/labstack/echo/v4"
)

func RouteOrder(e *echo.Echo){
	e.GET("/orders", Controller.GetAllOrders)
	e.POST("/orders", Controller.CreateOrder)
	e.GET("/orders/:id", Controller.GetAOrder)
	e.PUT("/orders/:id", Controller.UpdateOrder)
	e.DELETE("/orders/:id", Controller.DeleteOrder)
}

func RouteProduct(e *echo.Echo){
	e.GET("/orders/:id/product",Controller.GetOrderProduct)
}
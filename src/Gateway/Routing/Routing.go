package Routing

import (
	"Api.Calisma/src/Gateway/Controllers"
	"github.com/labstack/echo/v4"
)

func RouteControllers(e *echo.Echo) {
	RouteCustomer(e)
	RouteOrder(e)
	RouteAddress(e)
	RouteProduct(e)
	RouteToken(e)
}

func RouteCustomer(e *echo.Echo) {
	e.GET("/customers", Controllers.GetAllCustomers)
	e.POST("/customers", Controllers.CreateCustomer)
	e.GET("/customers/:id", Controllers.GetACustomer)
	e.PUT("/customers/:id", Controllers.UpdateCustomer)
	e.DELETE("/customers/:id", Controllers.DeleteCustomer)
}

func RouteAddress(e *echo.Echo) {
	e.GET("/customers/:id/address", Controllers.GetCustomerAddresses)
}

func RouteToken(e *echo.Echo) {
	e.POST("/token", Controllers.CreateToken)
}
func RouteOrder(e *echo.Echo) {
	e.GET("/orders", Controllers.GetAllOrders)
	e.POST("/orders", Controllers.CreateOrder)
	e.GET("/orders/:id", Controllers.GetAOrder)
	e.PUT("/orders/:id", Controllers.UpdateOrder)
	e.DELETE("/orders/:id", Controllers.DeleteOrder)
}

func RouteProduct(e *echo.Echo) {
	e.GET("/orders/:id/product", Controllers.GetOrderProduct)
}

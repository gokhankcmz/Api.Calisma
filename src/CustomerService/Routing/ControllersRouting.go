package Routing

import (
	Controller "Api.Calisma/src/CustomerService/Controller"
	"github.com/labstack/echo/v4"
)

func RouteCustomer(e *echo.Echo){
	e.GET("/customers", Controller.GetAllCustomers)
	e.POST("/customers", Controller.CreateCustomer)
	e.GET("/customers/:id", Controller.GetACustomer)
	e.PUT("/customers/:id", Controller.UpdateCustomer)
	e.DELETE("/customers/:id", Controller.DeleteCustomer)
}

func RouteAddress(e *echo.Echo){
	e.GET("/customers/:id/address", Controller.GetCustomerAddresses)
}

func RouteToken(e *echo.Echo){
	e.POST("/token", Controller.CreateToken)
}
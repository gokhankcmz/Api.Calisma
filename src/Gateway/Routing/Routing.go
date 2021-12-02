package Routing

import (
	"Api.Calisma/src/Common/Middlewares"
	"Api.Calisma/src/Gateway/Handlers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RouteControllers(e *echo.Echo) {
	RouteCustomer(e)
	RouteOrder(e)
}

func RouteCustomer(e *echo.Echo) {
	customers := e.Group("/customers")
	customers.GET("", Handlers.GetAllCustomers)
	customers.POST("", Handlers.CreateCustomer)
	customers.GET("/:id", Handlers.GetACustomer, Middlewares.TokenValidatorMiddleware("User cannot request another user's data."))
	customers.PUT("/:id", Handlers.UpdateCustomer, Middlewares.TokenValidatorMiddleware("User cannot update another user's data."))
	customers.DELETE("/:id", Handlers.DeleteCustomer, Middlewares.TokenValidatorMiddleware("User cannot delete another user."))
	customers.GET("/:id/address", Handlers.GetCustomerAddresses, Middlewares.TokenValidatorMiddleware("User cannot request another user's data."))
	customers.GET("/:id/address", Handlers.GetCustomerAddresses)
	e.POST("/token", Handlers.CreateToken)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func RouteOrder(e *echo.Echo) {
	g := e.Group("/orders")
	g.GET("", Handlers.GetAOrder)
	g.POST("", Handlers.CreateOrder)
	g.GET("/:id", Handlers.GetAOrder)
	g.PUT("/:id", Handlers.UpdateOrder)
	g.DELETE("/:id", Handlers.DeleteOrder)
	g.GET("/:id/product", Handlers.GetOrderProduct)
}

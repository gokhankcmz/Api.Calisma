package Handlers

import (
	"Api.Calisma/src/Common/Middlewares"
	CustomerRepository "Api.Calisma/src/CustomerService/Repository"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handler struct {
	Repository CustomerRepository.IRepository
}

func NewHandler(e *echo.Echo, repository CustomerRepository.IRepository) *Handler {
	handler := &Handler{Repository: repository}
	g := e.Group("/customers")
	g.GET("", handler.GetCustomers)
	g.POST("", handler.CreateCustomer)
	g.GET("/:id", handler.GetACustomer, Middlewares.TokenValidatorMiddleware("User cannot request another user's data."))
	g.PUT("/:id", handler.UpdateCustomer, Middlewares.TokenValidatorMiddleware("User cannot update another user's data."))
	g.DELETE("/:id", handler.DeleteCustomer, Middlewares.TokenValidatorMiddleware("User cannot delete another user."))
	g.GET("/:id/address", handler.GetCustomerAddresses, Middlewares.TokenValidatorMiddleware("User cannot request another user's data."))

	e.POST("/token", handler.CreateToken)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return handler
}

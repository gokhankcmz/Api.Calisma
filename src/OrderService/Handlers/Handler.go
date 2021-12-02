package Handlers

import (
	OrderRepository "Api.Calisma/src/OrderService/Repository"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handler struct {
	Repository OrderRepository.IOrderRepository
}

func NewHandler(e *echo.Echo, repository OrderRepository.IOrderRepository) *Handler {
	handler := &Handler{Repository: repository}
	g := e.Group("/orders")
	g.GET("", handler.GetAOrder)
	g.POST("", handler.CreateOrder)
	g.GET("/:id", handler.GetAOrder)
	g.PUT("/:id", handler.UpdateOrder)
	g.DELETE("/:id", handler.DeleteOrder)
	g.GET("/:id/product", handler.GetOrderProduct)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return handler
}

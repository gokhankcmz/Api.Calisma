package Handlers

import (
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Redirection"
	"github.com/labstack/echo/v4"
)

// GetOrderProduct godoc
// @Summary Get Order Product.
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Param Token header string false "Authorization"
// @Router /orders/:id/product [get]
// @Tags Orders
// @Success 200 {object} Entities.Product
func GetOrderProduct(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.OrderServiceUri)
}

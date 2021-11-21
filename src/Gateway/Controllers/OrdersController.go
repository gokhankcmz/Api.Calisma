package Controllers

import (
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Redirection"
	"github.com/labstack/echo/v4"
)

// GetAllOrders godoc
// @Summary Get All Orders.
// @Accept  json
// @Produce  json
// @Router /orders [get]
// @Param fields query bool false "Detailed Fields"
// @Tags Orders
// @Success 200
func GetAllOrders(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.OrderServiceUri)
}

// CreateOrder godoc
// @Summary Create an Order.
// @Accept  json
// @Produce  json
// @Router /orders [post]
// @Tags Orders
// @Success 200
func CreateOrder(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.OrderServiceUri)
}

// DeleteOrder godoc
// @Summary Delete an Order.
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Param Authorization header string false "Authorization"
// @Router /orders/{id} [delete]
// @Tags Orders
// @Success 200
func DeleteOrder(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.OrderServiceUri)
}

// GetAOrder godoc
// @Summary Get an Order .
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Param Authorization header string false "Authorization"
// @Router /orders/{id} [get]
// @Tags Orders
// @Success 200 {object} Entities.Product
// @Param detail query bool false "Detailed Fields"
func GetAOrder(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.OrderServiceUri)
}

// UpdateOrder godoc
// @Summary Get Order Product.
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Param Authorization header string false "Authorization"
// @Router /orders/{id} [put]
// @Tags Orders
// @Success 200
func UpdateOrder(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.OrderServiceUri)
}

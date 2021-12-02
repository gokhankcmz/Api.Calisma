package Handlers

import (
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Redirection"
	"github.com/labstack/echo/v4"
)

// GetCustomerAddresses godoc
// @Summary Get Customer address.
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID"
// @Param Authorization header string false "Authorization"
// @Router /customers/{id}/address [get]
// @Tags Customers
// @Success 200 {object} Entities.Address
func GetCustomerAddresses(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.CustomerServiceUri)
}

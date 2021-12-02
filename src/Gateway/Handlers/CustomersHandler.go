package Handlers

import (
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Redirection"
	"github.com/labstack/echo/v4"
)

// GetAllCustomers godoc
// @Summary Get All Customer .
// @Accept  json
// @Produce  json
// @Router /customers [get]
// @Param fields query bool false "Detailed Fields"
// @Tags Customers
// @Success 200
func GetAllCustomers(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.CustomerServiceUri)
}

// CreateCustomer godoc
// @Summary Create Customer
// @Accept  json
// @Produce  json
// @Param Customer body RequestModels.CreateCustomerDto true "Customer to create."
// @Router /customers [post]
// @Tags Customers
// @Success 201
func CreateCustomer(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.CustomerServiceUri)
}

// DeleteCustomer godoc
// @Summary Delete a Customer.
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID"
// @Param Authorization header string false "Authorization"
// @Router /customers/{id} [delete]
// @Tags Customers
// @Success 200
func DeleteCustomer(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.CustomerServiceUri)
}

// GetACustomer godoc
// @Summary Get a Customer.
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID"
// @Param Authorization header string false "Authorization"
// @Router /customers/{id} [get]
// @Tags Customers
// @Success 200 {object} Entities.Customer
// @Param detail query bool false "Detailed Fields"
func GetACustomer(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.CustomerServiceUri)
}

// UpdateCustomer godoc
// @Summary Updates a Customer.
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID"
// @Param Authorization header string false "Authorization"
// @Router /customers/{id} [put]
// @Tags Customers
// @Success 200
// @Param Customer body RequestModels.UpdateCustomerDto true "Customer to update."
func UpdateCustomer(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.CustomerServiceUri)
}

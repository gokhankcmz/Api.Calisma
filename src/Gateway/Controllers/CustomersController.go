package Controllers

import (
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Token"
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
// @Summary Get all Customers.
// @Accept  json
// @Produce  json
// @Param Customer body RequestModels.CreateCustomerDto true "Customer to create."
// @Router /customers [post]
// @Tags Customers
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
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id {
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("User cannot delete another user."))
	}
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
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id {
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("The user cannot request another user's data."))
	}
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
func UpdateCustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id {
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("The user cannot update another user's data."))
	}
	return Redirection.Redirect(ctx, Constants.CustomerServiceUri)
}

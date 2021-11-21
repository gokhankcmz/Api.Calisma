package Controllers

import (
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Token"
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
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id {
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("The address does not belong to this person."))
	}
	return Redirection.Redirect(ctx, Constants.CustomerServiceUri)
}

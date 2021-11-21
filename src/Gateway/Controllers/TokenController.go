package Controllers

import (
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Redirection"
	"github.com/labstack/echo/v4"
)

// CreateToken godoc
// @Summary Create a JWT Token.
// @Accept  json
// @Produce  json
// @Param Customer body RequestModels.TokenCredentials true "Customer Credentials."
// @Router /token [post]
// @Tags Token
// @Success 200
func CreateToken(ctx echo.Context) error {
	return Redirection.Redirect(ctx, Constants.CustomerServiceUri)
}

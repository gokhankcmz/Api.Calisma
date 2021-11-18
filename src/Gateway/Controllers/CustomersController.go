package Controllers

import (
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Token"
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Redirection"
	"github.com/labstack/echo/v4"
)

func GetAllCustomers(ctx echo.Context) error{
	return Redirection.Redirect(ctx,Constants.CustomerServiceUri)
}

func CreateCustomer(ctx echo.Context) error{
	return Redirection.Redirect(ctx,Constants.CustomerServiceUri)
}

func DeleteCustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id{
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("User cannot delete another user."))
	}
	return Redirection.Redirect(ctx,Constants.CustomerServiceUri)

}

func GetACustomer(ctx echo.Context) error{
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id{
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("The user cannot request another user's data."))
	}
	return Redirection.Redirect(ctx,Constants.CustomerServiceUri)
}



func UpdateCustomer(ctx echo.Context) error{
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id{
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("The user cannot update another user's data."))
	}
	return Redirection.Redirect(ctx,Constants.CustomerServiceUri)
}


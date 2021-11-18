package Controller

import (
	"Api.Calisma/src/Common/Token"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetOrderProduct(ctx echo.Context) error{
	orderID := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	OrderEntity := PanicIfUnauthorizedRequest(orderID,claims.ID,"Users cannot see the product information of an order that does not belong to them.")
	return ctx.JSON(http.StatusOK,OrderEntity.Product)
}

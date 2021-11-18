package Controller

import (
	"Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Token"
	"Api.Calisma/src/CustomerService/Repository"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCustomerAddresses(ctx echo.Context) error{
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id{
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("The address does not belong to this person."))
	}
	response:= CustomerRepository.GetCustomer(id)
	CustomerEntity := Entities.Customer{}
	_ = json.Unmarshal(response,&CustomerEntity)
	return ctx.JSON(http.StatusOK, CustomerEntity.Address)
}
package Controller

import (
	"Api.Calisma/src/Common/Helpers"
	"Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Models/RequestModels"
	"Api.Calisma/src/Common/Models/ResponseModels"
	"Api.Calisma/src/Common/Token"
	"Api.Calisma/src/CustomerService/Mapper"
	"Api.Calisma/src/CustomerService/Repository"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllCustomers(ctx echo.Context) error{
	fields := ctx.QueryParam("fields")
	if fields == "true"{
		response := CustomerRepository.GetAllCustomers()
		return ctx.JSONBlob(http.StatusOK, response)
	}
	response := CustomerRepository.GetAllCustomerIds()
	return ctx.JSONBlob(http.StatusOK, response)
}

func CreateCustomer(ctx echo.Context) error{
	defer ctx.Request().Body.Close()
	var createDTO RequestModels.CreateCustomerDto
	_= json.NewDecoder(ctx.Request().Body).Decode(&createDTO)
	Helpers.ValidateModelOrPanic(createDTO)
	customerEntity := Mapper.CreateCustomerDtoToCustomer(createDTO)
	response := CustomerRepository.CreateCustomer(&customerEntity)
	return ctx.JSONBlob(http.StatusCreated, response)
}

func DeleteCustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id{
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("User cannot delete another user."))
	}
	response := CustomerRepository.DeleteCustomer(id)
	return ctx.JSONBlob(http.StatusOK, response)

}

func GetACustomer(ctx echo.Context) error{
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id{
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("The user cannot request another user's data."))
	}
	addressRequested := ctx.QueryParam("detail")
	response:= CustomerRepository.GetCustomer(id)
	if addressRequested=="true" {
		return ctx.JSONBlob(http.StatusOK, response)
	}
	var CustomerResponseDto = ResponseModels.CustomerResponseDto{}
	err := json.Unmarshal(response,&CustomerResponseDto)
	if err != nil{
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, CustomerResponseDto)
}



func UpdateCustomer(ctx echo.Context) error{
	id := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	if claims.ID != id{
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail("The user cannot update another user's data."))
	}
	var updateDTO RequestModels.UpdateCustomerDto
	defer ctx.Request().Body.Close()
	_= json.NewDecoder(ctx.Request().Body).Decode(&updateDTO)
	Helpers.ValidateModelOrPanic(updateDTO)
	customerEntityBytes := CustomerRepository.GetCustomer(id)
	var customerEntity Entities.Customer
	err := json.Unmarshal(customerEntityBytes, &customerEntity)
	if err != nil{
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	Mapper.AppylChangesToCustomer(updateDTO, &customerEntity)
	response:= CustomerRepository.UpdateCustomer(&customerEntity)
	return ctx.JSONBlob(http.StatusOK, response)
}

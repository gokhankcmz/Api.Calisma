package Controller

import (
	"Api.Calisma/src/Common/Helpers"
	"Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Models/RequestModels"
	"Api.Calisma/src/Common/Token"
	"Api.Calisma/src/OrderService/Constants"
	"Api.Calisma/src/OrderService/Mapper"
	"Api.Calisma/src/OrderService/Repository"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"time"
)

func GetAllOrders(ctx echo.Context) error{
	fields := ctx.QueryParam("fields")
	if fields == "true"{
		response := OrderRepository.GetAllOrders()
		return ctx.JSONBlob(http.StatusOK, response)
	}
	response := OrderRepository.GetAllOrderIds()
	return ctx.JSONBlob(http.StatusOK, response)
}

func CreateOrder(ctx echo.Context) error{
	defer ctx.Request().Body.Close()
	var createDTO RequestModels.CreateOrderDto
	_= json.NewDecoder(ctx.Request().Body).Decode(&createDTO)
	Helpers.ValidateModelOrPanic(createDTO)
	OrderEntity := Mapper.MapOrderDtoToOrder(createDTO)
	token := ctx.Request().Header.Get("Authorization")
	claims := Token.ValidateToken(token)
	OrderEntity.CustomerID = claims.ID
	OrderEntity.Address = GetCustomerAddress(OrderEntity.CustomerID,token)
	response := OrderRepository.CreateOrder(&OrderEntity)
	return ctx.JSONBlob(http.StatusOK, response)
}

func DeleteOrder(ctx echo.Context) error {
	orderID := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	_ = PanicIfUnauthorizedRequest(orderID,claims.ID,"Users cannot delete an order that does not belong to them.")
	response := OrderRepository.DeleteOrder(orderID)
	return ctx.JSONBlob(http.StatusOK, response)

}

func GetAOrder(ctx echo.Context) error{
	orderID := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	OrderEntity := PanicIfUnauthorizedRequest(orderID,claims.ID,"Users cannot request an order that does not belong to them.")
	Detail := ctx.QueryParam("detail")
	if Detail=="true" {
		return ctx.JSON(http.StatusOK, OrderEntity)
	}
	return ctx.JSON(http.StatusCreated, Mapper.MapOrderToOrderResponseDto(OrderEntity))
}



func UpdateOrder(ctx echo.Context) error{
	defer ctx.Request().Body.Close()
	orderID := ctx.Param("id")
	claims := Token.ValidateToken(ctx.Request().Header.Get("Authorization"))
	OrderEntity := PanicIfUnauthorizedRequest(orderID,claims.ID,"Users cannot request an order that does not belong to them.")
	var updateDTO RequestModels.UpdateOrderDto
	_= json.NewDecoder(ctx.Request().Body).Decode(&updateDTO)
	Helpers.ValidateModelOrPanic(updateDTO)
	Mapper.ApplyChangesToOrder(&updateDTO, OrderEntity)
	response:= OrderRepository.UpdateOrder(OrderEntity)
	return ctx.JSONBlob(http.StatusOK, response)
}

func PanicIfUnauthorizedRequest(orderID, customerID, PanicText string) *Entities.Order{
	response:= OrderRepository.GetOrder(orderID)
	var OrderEntity = Entities.Order{}
	err := json.Unmarshal(response,&OrderEntity)
	if err != nil{
		fmt.Println(err)
	}
	if customerID != OrderEntity.CustomerID{
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail(PanicText))
	}
	return &OrderEntity
}

func GetCustomerAddress(customerID,token string) Entities.Address{
	req,err := http.NewRequest("GET", Constants.CustomerServiceUri + "/customers/" + customerID + "/address",nil)
	if err != nil{
		fmt.Println(err)
	}
	if token != ""{
		req.Header.Set("Authorization", token)
	}
	client:= http.DefaultClient
	client.Timeout = time.Second*10
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var Address = Entities.Address{}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println(err)
	}
	err = json.Unmarshal(body,&Address)

	if err!=nil{
		fmt.Println(err)
	}
	return Address
}

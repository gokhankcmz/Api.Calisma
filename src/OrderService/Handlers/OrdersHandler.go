package Handlers

import (
	"Api.Calisma/src/Common/Helpers"
	"Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Models/RequestModels"
	"Api.Calisma/src/Common/Token"
	"Api.Calisma/src/OrderService/Constants"
	"Api.Calisma/src/OrderService/Mapper"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"time"
)

//TODO: PAGINATION
func (h Handler) GetAllOrders(ctx echo.Context) error {
	fields := ctx.QueryParam("fields")
	if fields == "true" {
		response := h.Repository.GetAllOrders()
		return ctx.JSON(http.StatusOK, response)
	}
	response, _ := h.Repository.GetAllOrderIds()
	return ctx.JSON(http.StatusOK, response)
}

func (h Handler) CreateOrder(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var createDTO RequestModels.CreateOrderDto
	_ = json.NewDecoder(ctx.Request().Body).Decode(&createDTO)
	Helpers.ValidateModelOrPanic(createDTO)
	OrderEntity := Mapper.MapOrderDtoToOrder(createDTO)
	token := ctx.Request().Header.Get("Authorization")
	claims := Token.ValidateAndGetClaims(token)
	OrderEntity.CustomerID = claims.ID
	OrderEntity.Address = h.GetCustomerAddress(OrderEntity.CustomerID, token)
	response := h.Repository.CreateOrder(&OrderEntity)
	return ctx.JSON(http.StatusCreated, response)
}

func (h Handler) DeleteOrder(ctx echo.Context) error {
	orderID := ctx.Param("id")
	claims := Token.ValidateAndGetClaims(ctx.Request().Header.Get("Authorization"))
	_ = h.PanicIfUnauthorizedRequest(orderID, claims.ID, "Users cannot delete an order that does not belong to them.")
	response := h.Repository.DeleteOrder(orderID)
	return ctx.JSON(http.StatusOK, response)
}

//TODO: REPODAN OBJE DÖNSÜN
func (h Handler) GetAOrder(ctx echo.Context) error {
	orderID := ctx.Param("id")
	claims := Token.ValidateAndGetClaims(ctx.Request().Header.Get("Authorization"))
	OrderEntity := h.PanicIfUnauthorizedRequest(orderID, claims.ID, "Users cannot request an order that does not belong to them.")
	Detail := ctx.QueryParam("detail")
	if Detail == "true" {
		return ctx.JSON(http.StatusOK, OrderEntity)
	}
	return ctx.JSON(http.StatusOK, Mapper.MapOrderToOrderResponseDto(OrderEntity))
}

func (h Handler) UpdateOrder(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	orderID := ctx.Param("id")
	claims := Token.ValidateAndGetClaims(ctx.Request().Header.Get("Authorization"))
	OrderEntity := h.PanicIfUnauthorizedRequest(orderID, claims.ID, "Users cannot request an order that does not belong to them.")
	var updateDTO RequestModels.UpdateOrderDto
	_ = json.NewDecoder(ctx.Request().Body).Decode(&updateDTO)
	Helpers.ValidateModelOrPanic(updateDTO)
	Mapper.ApplyChangesToOrder(&updateDTO, OrderEntity)
	response := h.Repository.UpdateOrder(OrderEntity)
	return ctx.JSON(http.StatusOK, response)
}

func (h Handler) PanicIfUnauthorizedRequest(orderID, customerID, PanicText string) *Entities.Order {
	response := h.Repository.GetOrder(orderID)
	if customerID != response.CustomerID {
		panic(ErrorModels.UnauthorizedRequest.SetPublicDetail(PanicText))
	}
	return &response
}

func (h Handler) GetCustomerAddress(customerID, token string) Entities.Address {
	//todo: Fasthttp
	req, err := http.NewRequest("GET", Constants.CustomerServiceUri+"/customers/"+customerID+"/address", nil)
	if err != nil {
		fmt.Println(err)
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	client := http.DefaultClient
	client.Timeout = time.Second * 10
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var Address = Entities.Address{}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &Address)

	if err != nil {
		fmt.Println(err)
	}
	return Address
}

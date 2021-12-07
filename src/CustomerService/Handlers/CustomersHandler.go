package Handlers

import (
	"Api.Calisma/src/Common/Helpers"
	"Api.Calisma/src/Common/Models/RequestModels"
	"Api.Calisma/src/CustomerService/Mapper"
	"Api.Calisma/src/CustomerService/MongoFilters"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) GetCustomers(ctx echo.Context) error {
	fields := ctx.QueryParam("detail")
	paginator := MongoFilters.GetPaginator(ctx)
	filter := MongoFilters.GetSearchFilter(ctx)
	if fields == "true" {
		response, _ := h.Repository.GetCustomers(paginator, filter)
		return ctx.JSON(http.StatusOK, response)
	}
	response, _ := h.Repository.GetAllCustomerIds()
	return ctx.JSON(http.StatusOK, response)
}

func (h Handler) CreateCustomer(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var createDTO RequestModels.CreateCustomerDto
	_ = json.NewDecoder(ctx.Request().Body).Decode(&createDTO)
	Helpers.ValidateModelOrPanic(createDTO)
	customerEntity := Mapper.CreateCustomerFromDto(&createDTO)
	response := h.Repository.CreateCustomer(&customerEntity)
	return ctx.JSON(http.StatusCreated, response)
}

func (h Handler) DeleteCustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	response := h.Repository.DeleteCustomer(id)
	return ctx.JSON(http.StatusOK, response)

}

func (h Handler) GetACustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	detailRequested := ctx.QueryParam("detail")
	response := h.Repository.GetCustomer(id)
	if detailRequested == "true" {
		return ctx.JSON(http.StatusOK, response)
	}
	var CustomerResponseDto = Mapper.GetCustomerResponse(response)
	return ctx.JSON(http.StatusOK, CustomerResponseDto)
}

func (h Handler) UpdateCustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	customerEntity := h.Repository.GetCustomer(id)
	json.NewDecoder(ctx.Request().Body).Decode(&customerEntity)
	defer ctx.Request().Body.Close()
	Helpers.ValidateModelOrPanic(customerEntity)
	response := h.Repository.UpdateCustomer(&customerEntity)
	return ctx.JSON(http.StatusOK, response)
}

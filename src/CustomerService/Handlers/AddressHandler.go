package Handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) GetCustomerAddresses(ctx echo.Context) error {
	id := ctx.Param("id")
	response := h.Repository.GetCustomer(id)
	return ctx.JSON(http.StatusOK, response.Address)
}

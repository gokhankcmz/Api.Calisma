package Handlers

import (
	"Api.Calisma/src/Common/Helpers"
	"Api.Calisma/src/Common/Models/RequestModels"
	"Api.Calisma/src/Common/Token"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) CreateToken(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var tc RequestModels.TokenCredentials
	_ = json.NewDecoder(ctx.Request().Body).Decode(&tc)
	Helpers.ValidateModelOrPanic(tc)
	h.Repository.CheckCustomerIfExist(&tc)
	token := Token.CreateToken(&tc)
	return ctx.JSON(http.StatusOK, token)
}

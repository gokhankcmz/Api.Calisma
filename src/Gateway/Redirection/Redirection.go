package Redirection

import (
	"Api.Calisma/src/Gateway/Constants"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func Redirect(ctx echo.Context, serviceBaseUri string) error{
	defer ctx.Request().Body.Close()

	uri := ctx.Request().RequestURI
	method := ctx.Request().Method
	req,err := http.NewRequest(method,serviceBaseUri + uri, ctx.Request().Body)
	if err != nil{
		fmt.Println(err)
	}

	token := ctx.Request().Header.Get("Authorization")
	if token != ""{
		req.Header.Set("Authorization", token)
	}

	client:= http.DefaultClient
	client.Timeout = Constants.ClientTimeout
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}

	return ctx.JSONBlob(resp.StatusCode,body)
}

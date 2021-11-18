package main

import (
	"Api.Calisma/src/Common/Logger"
	"Api.Calisma/src/Common/Middlewares"
	"Api.Calisma/src/CustomerService/Constants"
	"Api.Calisma/src/CustomerService/Routing"
	"fmt"
	"github.com/labstack/echo/v4"
)
func main(){
	fmt.Println("Application starting...")
	e := echo.New()

	Middlewares.UsePanicHandlerMiddleware(e, Constants.ApplicationName)
	Logger.UseLogrusRequestLogging(e, Constants.ApplicationName)
	Routing.RouteCustomer(e)
	Routing.RouteAddress(e)
	Routing.RouteToken(e)

	//e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Start(":8000")
}

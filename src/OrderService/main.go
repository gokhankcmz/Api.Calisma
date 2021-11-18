package main

import (
	"Api.Calisma/src/Common/Logger"
	"Api.Calisma/src/Common/Middlewares"
	"Api.Calisma/src/OrderService/Constants"
	"Api.Calisma/src/OrderService/Routing"
	"fmt"
	"github.com/labstack/echo/v4"
)
func main(){
	fmt.Println("Application starting...")
	e := echo.New()

	Middlewares.UsePanicHandlerMiddleware(e, Constants.ApplicationName)
	Logger.UseLogrusRequestLogging(e, Constants.ApplicationName)
	Routing.RouteOrder(e)
	Routing.RouteProduct(e)
	e.Start(":8001")
}

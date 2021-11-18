package main

import (
	"Api.Calisma/src/Common/Logger"
	"Api.Calisma/src/Common/Middlewares"
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Routing"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main(){
	fmt.Println("Application starting...")
	e := echo.New()

	Middlewares.UsePanicHandlerMiddleware(e, Constants.ApplicationName)
	Logger.UseLogrusRequestLogging(e, Constants.ApplicationName)
	Routing.RouteControllers(e)
	e.Start(":80")
}


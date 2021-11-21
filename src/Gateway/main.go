package main

import (
	"Api.Calisma/src/Common/Logger"
	"Api.Calisma/src/Common/Middlewares"
	"Api.Calisma/src/Gateway/Constants"
	"Api.Calisma/src/Gateway/Routing"
	_ "Api.Calisma/src/Gateway/docs"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"os/exec"
)

func main() {
	fmt.Println("Application starting...")
	e := echo.New()

	Middlewares.UsePanicHandlerMiddleware(e, Constants.ApplicationName)
	Logger.UseLogrusRequestLogging(e, Constants.ApplicationName)
	Routing.RouteControllers(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	_ = exec.Command("cmd", "/C", "start", "chrome.exe", "localhost/swagger/index.html").Run()
	e.Start(":" + Constants.ApplicationPort)
}

// @title Gateway
// @version 1.0
// @description OrderCase Gateway Documentation.

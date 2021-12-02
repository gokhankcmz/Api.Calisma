package main

import (
	"Api.Calisma/src/Common/Logging/LogBody"
	"Api.Calisma/src/Common/Middlewares"
	"Api.Calisma/src/OrderService/Handlers"
	"Api.Calisma/src/OrderService/Mongo"
	OrderRepository "Api.Calisma/src/OrderService/Repository"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Application starting...")
	e := echo.New()
	LogBody.CreateStaticBody().AddApplicationInfo("CustomerService-Api").AddHostName()
	Middlewares.UsePanicHandlerMiddleware(e)
	Middlewares.UseLogrusRequestLogging(e)
	repository := OrderRepository.NewRepository(Mongo.GetMongoSingletonCollection())
	Handlers.NewHandler(e, repository)
	e.Start(":8001")
}

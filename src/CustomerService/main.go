package main

import (
	"Api.Calisma/src/Common/Logging/LogBody"
	"Api.Calisma/src/Common/Middlewares"
	"Api.Calisma/src/CustomerService/Handlers"
	"Api.Calisma/src/CustomerService/Mongo"
	CustomerRepository "Api.Calisma/src/CustomerService/Repository"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Application starting...")
	e := echo.New()

	//KafkaProducer.CreateProducer("localhost:9092")
	LogBody.CreateStaticBody().AddApplicationInfo("CustomerService-Api").AddHostName()
	Middlewares.UsePanicHandlerMiddleware(e)
	Middlewares.UseLogrusRequestLogging(e)
	repo := CustomerRepository.NewRepository(Mongo.GetMongoSingletonCollection())
	Handlers.NewHandler(e, repo)
	//e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Start(":8000")
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/app/handlers"
	"github.com/shaikrasheed99/broker-service/app/routes"
	"github.com/shaikrasheed99/broker-service/app/services"
)

func main() {
	app := gin.New()

	bs := services.NewBrokerService()

	handler := handlers.NewBrokerHandler(bs)

	routes.RegisterRoutes(app, handler)

	port := ":8000"

	app.Run(port)
}

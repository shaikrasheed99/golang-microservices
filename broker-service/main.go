package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/app/handlers"
	"github.com/shaikrasheed99/broker-service/app/routes"
)

func main() {
	app := gin.New()

	handler := handlers.NewBrokerHandler()

	routes.RegisterRoutes(app, handler)

	port := ":8000"

	app.Run(port)
}

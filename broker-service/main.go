package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/handlers"
	"github.com/shaikrasheed99/broker-service/routes"
)

func main() {
	app := gin.New()

	handler := handlers.NewBrokerHandler()

	routes.RegisterRoutes(app, handler)

	port := ":8000"

	app.Run(port)
}

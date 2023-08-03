package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/routes"
)

func main() {
	app := gin.New()

	routes.RegisterRoutes(app)

	port := ":8000"

	app.Run(port)
}

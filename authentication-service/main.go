package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/authentication-service/handlers"
	"github.com/shaikrasheed99/authentication-service/routes"
)

func main() {
	app := gin.New()

	handler := handlers.NewAuthHandler()

	routes.RegisterRoutes(app, handler)

	port := ":8001"

	app.Run(port)
}

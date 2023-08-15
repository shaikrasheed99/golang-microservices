package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/authentication-service/app/handlers"
	"github.com/shaikrasheed99/authentication-service/app/routes"
	"github.com/shaikrasheed99/authentication-service/app/services"
	"github.com/shaikrasheed99/authentication-service/configs"
	"github.com/shaikrasheed99/authentication-service/persistence/database"
	"github.com/shaikrasheed99/authentication-service/persistence/repositories"
)

func main() {
	err := configs.LoadConfigs()
	if err != nil {
		log.Println("[Main]", err.Error())
		return
	}

	db, err := database.InitDatabase()
	if err != nil {
		log.Println("[Main]", err.Error())
		return
	}

	ar := repositories.NewAuthRepository(db)

	as := services.NewAuthService(ar)

	handler := handlers.NewAuthHandler(as)

	app := gin.New()

	routes.RegisterRoutes(app, handler)

	port := ":8001"

	app.Run(port)
}

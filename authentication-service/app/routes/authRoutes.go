package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/authentication-service/app/handlers"
	"github.com/shaikrasheed99/authentication-service/utils/constants"
)

func RegisterRoutes(engine *gin.Engine, ah handlers.IAuthHandler) {
	log.Println("[RegisterRoutes] Registering routes")

	engine.POST(constants.SignupUserEndpoint, ah.SignupHandler)
	engine.POST(constants.LoginUserEndpoint, ah.LoginHandler)
	engine.GET(constants.HealthEndpoint, ah.Health)
}

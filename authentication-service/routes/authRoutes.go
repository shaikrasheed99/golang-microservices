package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/authentication-service/constants"
	"github.com/shaikrasheed99/authentication-service/handlers"
)

func RegisterRoutes(engine *gin.Engine, h handlers.IAuthHandler) {
	log.Println("[RegisterRoutes] Registering routes")

	engine.GET(constants.HealthEndpoint, h.Health)
}

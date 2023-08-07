package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/constants"
	"github.com/shaikrasheed99/broker-service/handlers"
)

func RegisterRoutes(engine *gin.Engine, h handlers.IBrokerHandler) {
	log.Println("[RegisterRoutes] Registering routes")

	engine.GET(constants.HealthEndpoint, h.Health)
}

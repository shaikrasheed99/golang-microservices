package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/app/handlers"
	"github.com/shaikrasheed99/broker-service/utils/constants"
)

func RegisterRoutes(engine *gin.Engine, bh handlers.IBrokerHandler) {
	log.Println("[RegisterRoutes] Registering routes")

	engine.POST(constants.HandleEndpoint, bh.Handle)
	engine.GET(constants.HealthEndpoint, bh.Health)
}

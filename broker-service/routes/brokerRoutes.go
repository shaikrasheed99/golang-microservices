package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/constants"
	"github.com/shaikrasheed99/broker-service/handlers"
)

func RegisterRoutes(engine *gin.Engine, h handlers.IBrokerHandler) {
	engine.GET(constants.HealthEndpoint, h.Health)
}

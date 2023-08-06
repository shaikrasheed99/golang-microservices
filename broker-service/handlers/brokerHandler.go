package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/helpers"
)

type IBrokerHandler interface {
	Health(*gin.Context)
}

type brokerHandler struct{}

func NewBrokerHandler() IBrokerHandler {
	return &brokerHandler{}
}

func (br *brokerHandler) Health(c *gin.Context) {
	log.Println("[HealthHandler] Hitting health handler")

	res := helpers.CreateSuccessResponse(http.StatusOK, "UP", nil)

	c.JSON(http.StatusOK, res)
}

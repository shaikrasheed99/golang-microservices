package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/utils/constants"
	"github.com/shaikrasheed99/broker-service/utils/helpers"
)

type IBrokerHandler interface {
	Health(*gin.Context)
}

type brokerHandler struct{}

func NewBrokerHandler() IBrokerHandler {
	log.Println("[NewBrokerHandler] Creating new broker handler")

	return &brokerHandler{}
}

func (br *brokerHandler) Health(c *gin.Context) {
	log.Println("[HealthHandler] Hitting health handler")

	res := helpers.CreateSuccessResponse(http.StatusOK, constants.UP, nil)

	c.JSON(http.StatusOK, res)
}

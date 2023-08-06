package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{
		"message": "UP",
	})
}

package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	log.Println("[HealthHandler] Hitting health handler")

	c.JSON(http.StatusOK, gin.H{
		"message": "UP",
	})
}

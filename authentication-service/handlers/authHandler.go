package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/authentication-service/constants"
	"github.com/shaikrasheed99/authentication-service/helpers"
)

type IAuthHandler interface {
	Health(*gin.Context)
}

type authHandler struct{}

func NewAuthHandler() IAuthHandler {
	log.Println("[NewAuthHandler] Creating new auth handler")

	return &authHandler{}
}

func (br *authHandler) Health(c *gin.Context) {
	log.Println("[HealthHandler] Hitting health handler")

	res := helpers.CreateSuccessResponse(http.StatusOK, constants.UP, nil)

	c.JSON(http.StatusOK, res)
}

package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/authentication-service/constants"
	"github.com/shaikrasheed99/authentication-service/helpers"
	"github.com/shaikrasheed99/authentication-service/requests"
	"github.com/shaikrasheed99/authentication-service/services"
)

type IAuthHandler interface {
	SignupHandler(*gin.Context)
	Health(*gin.Context)
}

type authHandler struct {
	as services.IAuthService
}

func NewAuthHandler(as services.IAuthService) IAuthHandler {
	log.Println("[NewAuthHandler] Creating new auth handler")

	return &authHandler{
		as: as,
	}
}

func (ah *authHandler) SignupHandler(c *gin.Context) {
	log.Println("[SignupHandler] Hitting signup handler function in auth handler")

	var req *requests.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("[SignupHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	savedUser, err := ah.as.SaveUser(req)
	if err != nil {
		log.Println("[SignupHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully saved user details", savedUser.Username)

	log.Println("[SignupHandler] Finished execution of signup handler")
	c.JSON(http.StatusCreated, res)
}

func (br *authHandler) Health(c *gin.Context) {
	log.Println("[HealthHandler] Hitting health handler")

	res := helpers.CreateSuccessResponse(http.StatusOK, constants.UP, nil)

	c.JSON(http.StatusOK, res)
}

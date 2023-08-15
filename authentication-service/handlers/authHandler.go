package handlers

import (
	"errors"
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
	LoginHandler(*gin.Context)
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

	_, err := ah.as.SaveUser(req)
	if err != nil {
		log.Println("[SignupHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully saved user details", nil)

	log.Println("[SignupHandler] Finished execution of signup handler")
	c.JSON(http.StatusCreated, res)
}

func (ah *authHandler) LoginHandler(c *gin.Context) {
	log.Println("[LoginHandler] Hitting login handler function in auth handler")

	var req *requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("[LoginHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user, err := ah.as.LoginUser(req.Username)
	if err != nil {
		log.Println("[LoginHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	if req.Password != user.Password {
		err := errors.New(constants.ErrPasswordMismatch)
		log.Println("[LoginHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully logged in", nil)

	log.Println("[LoginHandler] Finished execution of login handler")
	c.JSON(http.StatusCreated, res)
}

func (br *authHandler) Health(c *gin.Context) {
	log.Println("[HealthHandler] Hitting health handler")

	res := helpers.CreateSuccessResponse(http.StatusOK, constants.UP, nil)

	c.JSON(http.StatusOK, res)
}

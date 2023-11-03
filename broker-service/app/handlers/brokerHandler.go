package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/app/models/requests"
	"github.com/shaikrasheed99/broker-service/app/services"
	"github.com/shaikrasheed99/broker-service/utils/constants"
	"github.com/shaikrasheed99/broker-service/utils/helpers"
)

type IBrokerHandler interface {
	Handle(*gin.Context)
	Health(*gin.Context)
}

type brokerHandler struct {
	bs services.IBrokerService
}

func NewBrokerHandler(bs services.IBrokerService) IBrokerHandler {
	log.Println("[NewBrokerHandler] Creating new broker handler")

	return &brokerHandler{
		bs: bs,
	}
}

func (bh *brokerHandler) Handle(c *gin.Context) {
	log.Println("[HandleHandler] Hitting handler handler")

	var req *requests.RequestPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("[HandleHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	var err error
	var serviceRes string
	switch req.Action {
	case constants.AuthSignupAction:
		serviceRes, err = bh.bs.HandleAuthSignup(&req.Signup)
	case constants.AuthLoginAction:
		serviceRes, err = bh.bs.HandleAuthLogin(&req.Login)
	default:
		err = errors.New(constants.ErrInvalidAction)
	}

	if err != nil {
		errMes := errors.New(constants.ErrAuthService)
		log.Println("[HandleHandler]", err.Error())
		errRes := helpers.CreateServiceErrorResponse(http.StatusInternalServerError, errMes.Error(), err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	res := helpers.CreateSuccessResponse(http.StatusOK, serviceRes, nil)

	c.JSON(http.StatusOK, res)
}

func (bh *brokerHandler) Health(c *gin.Context) {
	log.Println("[HealthHandler] Hitting health handler")

	res := helpers.CreateSuccessResponse(http.StatusOK, constants.UP, nil)

	c.JSON(http.StatusOK, res)
}

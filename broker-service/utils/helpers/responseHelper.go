package helpers

import (
	"log"
	"net/http"

	"github.com/shaikrasheed99/broker-service/app/models/responses"
	"github.com/shaikrasheed99/broker-service/utils/constants"
)

func CreateSuccessResponse(code int, message string, data interface{}) responses.SuccessResponse {
	log.Println("[CreateSuccessResponseHelper] Creating success response")

	return responses.SuccessResponse{
		Status:  constants.Success,
		Code:    http.StatusText(code),
		Message: message,
		Data:    data,
	}
}

package helpers

import (
	"net/http"

	"github.com/shaikrasheed99/broker-service/constants"
	"github.com/shaikrasheed99/broker-service/responses"
)

func CreateSuccessResponse(code int, message string, data interface{}) responses.SuccessResponse {
	return responses.SuccessResponse{
		Status:  constants.Success,
		Code:    http.StatusText(code),
		Message: message,
		Data:    data,
	}
}

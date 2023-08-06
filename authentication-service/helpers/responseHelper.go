package helpers

import (
	"net/http"

	"github.com/shaikrasheed99/authentication-service/constants"
	"github.com/shaikrasheed99/authentication-service/responses"
)

func CreateSuccessResponse(code int, message string, data interface{}) responses.SuccessResponse {
	return responses.SuccessResponse{
		Status:  constants.Success,
		Code:    http.StatusText(code),
		Message: message,
		Data:    data,
	}
}

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

func CreateErrorResponse(code int, message string) responses.ErrorResponse {
	log.Println("[CreateErrorResponseHelper] Creating error response")

	return responses.ErrorResponse{
		Status:  constants.Error,
		Code:    http.StatusText(code),
		Message: message,
	}
}

func CreateServiceErrorResponse(code int, message string, reason string) responses.ServiceErrorResponse {
	log.Println("[CreateServiceErrorResponseHelper] Creating service error response")

	return responses.ServiceErrorResponse{
		Status:  constants.Error,
		Code:    http.StatusText(code),
		Message: message,
		Reason:  reason,
	}
}

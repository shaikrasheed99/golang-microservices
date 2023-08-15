package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/app/models/requests"
	"github.com/shaikrasheed99/broker-service/app/models/responses"
	mocks "github.com/shaikrasheed99/broker-service/mocks/app/services"
	"github.com/shaikrasheed99/broker-service/utils/constants"
	"github.com/stretchr/testify/assert"
)

func TestBrokerHandler_Handle(t *testing.T) {
	req := &requests.RequestPayload{
		Action: constants.AuthAction,
		Signup: requests.AuthSignupPayload{
			Username: "test",
			Password: "test",
			Email:    "test@gmail.com",
		},
	}
	serviceRes := "response message from auth service"

	t.Run("should be able to handle auth signup api", func(t *testing.T) {
		mockService := new(mocks.IBrokerService)
		mockService.On("HandleAuthSignup", &req.Signup).Return(serviceRes, nil)

		handler := NewBrokerHandler(mockService)

		router := gin.Default()
		router.POST(constants.HandleEndpoint, handler.Handle)

		body, _ := json.Marshal(req)
		req, _ := http.NewRequest("POST", constants.HandleEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.SuccessResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Success, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusOK), resBody.Code)
		assert.Equal(t, serviceRes, resBody.Message)
		assert.Equal(t, nil, resBody.Data)
		mockService.AssertExpectations(t)
	})

	t.Run("should not be able to handle auth signup api when request payload is invalid", func(t *testing.T) {
		invalidReq := &requests.RequestPayload{
			Action: constants.AuthAction,
			Signup: requests.AuthSignupPayload{
				Username: "test",
				Password: "test",
			},
		}

		mockService := new(mocks.IBrokerService)

		handler := NewBrokerHandler(mockService)

		router := gin.Default()
		router.POST(constants.HandleEndpoint, handler.Handle)

		body, _ := json.Marshal(invalidReq)
		req, _ := http.NewRequest("POST", constants.HandleEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.ErrorResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), resBody.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("should be able to handle auth signup api when error from the auth service", func(t *testing.T) {
		serviceError := errors.New("error from auth service")
		mockService := new(mocks.IBrokerService)
		mockService.On("HandleAuthSignup", &req.Signup).Return("", serviceError)

		handler := NewBrokerHandler(mockService)

		router := gin.Default()
		router.POST(constants.HandleEndpoint, handler.Handle)

		body, _ := json.Marshal(req)
		req, _ := http.NewRequest("POST", constants.HandleEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.ServiceErrorResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), resBody.Code)
		assert.Equal(t, constants.ErrAuthService, resBody.Message)
		assert.Equal(t, serviceError.Error(), resBody.Reason)
		mockService.AssertExpectations(t)
	})
}

func TestBrokerHandler_Health(t *testing.T) {
	t.Run("should be able to up and running", func(t *testing.T) {
		mockService := new(mocks.IBrokerService)
		handler := NewBrokerHandler(mockService)

		router := gin.Default()
		router.GET(constants.HealthEndpoint, handler.Health)

		req, _ := http.NewRequest("GET", constants.HealthEndpoint, nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.SuccessResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Success, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusOK), resBody.Code)
		assert.Equal(t, constants.UP, resBody.Message)
		assert.Equal(t, nil, resBody.Data)
	})
}

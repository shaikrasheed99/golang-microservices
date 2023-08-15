package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/authentication-service/app/models/requests"
	"github.com/shaikrasheed99/authentication-service/app/models/responses"
	mocks "github.com/shaikrasheed99/authentication-service/mocks/app/services"
	"github.com/shaikrasheed99/authentication-service/persistence/models"
	"github.com/shaikrasheed99/authentication-service/utils/constants"
	"gotest.tools/assert"
)

func TestAuthHandler_SignupHandler(t *testing.T) {
	reqMock := &requests.SignupRequest{
		Username: "test",
		Password: "test",
		Email:    "test@gmail.com",
	}
	userMock := &models.User{
		ID:       1,
		Username: reqMock.Username,
		Password: reqMock.Password,
		Email:    reqMock.Email,
	}

	t.Run("should be able to signup a user", func(t *testing.T) {
		mockAuthService := new(mocks.IAuthService)
		mockAuthService.On("SaveUser", reqMock).Return(userMock, nil)

		handler := NewAuthHandler(mockAuthService)

		router := gin.Default()
		router.POST(constants.SignupUserEndpoint, handler.SignupHandler)

		body, _ := json.Marshal(reqMock)
		req, _ := http.NewRequest("POST", constants.SignupUserEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.SuccessResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Success, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusOK), resBody.Code)
		assert.Equal(t, "successfully saved user details", resBody.Message)
		assert.Equal(t, nil, resBody.Data)
		mockAuthService.AssertExpectations(t)
	})

	t.Run("should not be able to signup a user when invalid request", func(t *testing.T) {
		invalidReq := &requests.SignupRequest{
			Username: "test",
			Password: "test",
		}

		mockAuthService := new(mocks.IAuthService)

		handler := NewAuthHandler(mockAuthService)

		router := gin.Default()
		router.POST(constants.SignupUserEndpoint, handler.SignupHandler)

		body, _ := json.Marshal(invalidReq)
		req, _ := http.NewRequest("POST", constants.SignupUserEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.ErrorResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), resBody.Code)
		mockAuthService.AssertExpectations(t)
	})

	t.Run("should not be able to signup a user when error from service", func(t *testing.T) {
		serviceErr := errors.New("error from service")
		mockAuthService := new(mocks.IAuthService)
		mockAuthService.On("SaveUser", reqMock).Return(nil, serviceErr)

		handler := NewAuthHandler(mockAuthService)

		router := gin.Default()
		router.POST(constants.SignupUserEndpoint, handler.SignupHandler)

		body, _ := json.Marshal(reqMock)
		req, _ := http.NewRequest("POST", constants.SignupUserEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.ErrorResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), resBody.Code)
		assert.Equal(t, serviceErr.Error(), resBody.Message)
		mockAuthService.AssertExpectations(t)
	})
}

func TestAuthHandler_LoginHandler(t *testing.T) {
	reqMock := &requests.LoginRequest{
		Username: "test",
		Password: "test",
	}
	userMock := &models.User{
		ID:       1,
		Username: reqMock.Username,
		Password: reqMock.Password,
	}

	t.Run("should be able to login a user", func(t *testing.T) {
		mockAuthService := new(mocks.IAuthService)
		mockAuthService.On("LoginUser", reqMock.Username).Return(userMock, nil)

		handler := NewAuthHandler(mockAuthService)

		router := gin.Default()
		router.POST(constants.LoginUserEndpoint, handler.LoginHandler)

		body, _ := json.Marshal(reqMock)
		req, _ := http.NewRequest("POST", constants.LoginUserEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.SuccessResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Success, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusOK), resBody.Code)
		assert.Equal(t, constants.SuccessfullLogin, resBody.Message)
		assert.Equal(t, nil, resBody.Data)
		mockAuthService.AssertExpectations(t)
	})

	t.Run("should not be able to login a user when invalid request", func(t *testing.T) {
		invalidReq := &requests.SignupRequest{
			Username: "test",
		}

		mockAuthService := new(mocks.IAuthService)

		handler := NewAuthHandler(mockAuthService)

		router := gin.Default()
		router.POST(constants.LoginUserEndpoint, handler.LoginHandler)

		body, _ := json.Marshal(invalidReq)
		req, _ := http.NewRequest("POST", constants.LoginUserEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.ErrorResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), resBody.Code)
		mockAuthService.AssertExpectations(t)
	})

	t.Run("should not be able to login a user when error from service", func(t *testing.T) {
		serviceErr := errors.New("error from service")
		mockAuthService := new(mocks.IAuthService)
		mockAuthService.On("LoginUser", reqMock.Username).Return(nil, serviceErr)

		handler := NewAuthHandler(mockAuthService)

		router := gin.Default()
		router.POST(constants.LoginUserEndpoint, handler.LoginHandler)

		body, _ := json.Marshal(reqMock)
		req, _ := http.NewRequest("POST", constants.LoginUserEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.ErrorResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), resBody.Code)
		assert.Equal(t, serviceErr.Error(), resBody.Message)
		mockAuthService.AssertExpectations(t)
	})

	t.Run("should not be able to login a user when password is wrong", func(t *testing.T) {
		mockAuthService := new(mocks.IAuthService)
		mockAuthService.On("LoginUser", reqMock.Username).Return(userMock, nil)

		handler := NewAuthHandler(mockAuthService)

		router := gin.Default()
		router.POST(constants.LoginUserEndpoint, handler.LoginHandler)

		reqMock.Password = "another one"
		body, _ := json.Marshal(reqMock)
		req, _ := http.NewRequest("POST", constants.LoginUserEndpoint, bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.ErrorResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusUnauthorized), resBody.Code)
		assert.Equal(t, constants.ErrPasswordMismatch, resBody.Message)
		mockAuthService.AssertExpectations(t)
	})
}

func TestAuthHandler_Health(t *testing.T) {
	t.Run("should be able to up and running", func(t *testing.T) {
		mockAuthService := new(mocks.IAuthService)
		handler := NewAuthHandler(mockAuthService)

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

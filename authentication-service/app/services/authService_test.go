package services

import (
	"errors"
	"testing"

	"github.com/shaikrasheed99/authentication-service/app/models/requests"
	mocks "github.com/shaikrasheed99/authentication-service/mocks/persistence/repositories"
	"github.com/shaikrasheed99/authentication-service/persistence/models"
	"github.com/shaikrasheed99/authentication-service/utils/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAuthService_SaveUser(t *testing.T) {
	userReq := &requests.SignupRequest{
		Username: "test",
		Password: "test",
		Email:    "test@gmail.com",
	}
	userMock := models.User{
		ID:       1,
		Username: userReq.Username,
		Password: userReq.Password,
		Email:    userReq.Email,
	}

	t.Run("should be able to save user", func(t *testing.T) {
		mockRepo := new(mocks.IAuthRepository)
		mockRepo.On("FindUserByUsername", userReq.Username).Return(models.User{}, gorm.ErrRecordNotFound)
		mockRepo.On("SaveUser", mock.AnythingOfType("*models.User")).Return(userMock, nil)

		authService := NewAuthService(mockRepo)
		user, err := authService.SaveUser(userReq)

		assert.NoError(t, err)
		assert.Equal(t, user.Username, userMock.Username)
		assert.Equal(t, user.Password, userMock.Password)
		assert.Equal(t, user.Email, userMock.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not be able to save user when user is already exists", func(t *testing.T) {
		mockRepo := new(mocks.IAuthRepository)
		mockRepo.On("FindUserByUsername", userReq.Username).Return(userMock, nil)

		authService := NewAuthService(mockRepo)
		_, err := authService.SaveUser(userReq)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrUserAlreadyExists, err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not be able to save user when error from repo", func(t *testing.T) {
		dbError := errors.New("error from db")
		mockRepo := new(mocks.IAuthRepository)
		mockRepo.On("FindUserByUsername", userReq.Username).Return(userMock, gorm.ErrRecordNotFound)
		mockRepo.On("SaveUser", mock.AnythingOfType("*models.User")).Return(userMock, dbError)

		authService := NewAuthService(mockRepo)
		_, err := authService.SaveUser(userReq)

		assert.Error(t, err)
		assert.Equal(t, dbError.Error(), err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestAuthService_LoginUser(t *testing.T) {
	userMock := models.User{
		ID:       1,
		Username: "test",
		Password: "test",
		Email:    "test@gmail.com",
	}

	t.Run("should be able to login user", func(t *testing.T) {
		mockRepo := new(mocks.IAuthRepository)
		mockRepo.On("FindUserByUsername", userMock.Username).Return(userMock, nil)

		authService := NewAuthService(mockRepo)
		user, err := authService.LoginUser(userMock.Username)

		assert.NoError(t, err)
		assert.Equal(t, user.Username, userMock.Username)
		assert.Equal(t, user.Password, userMock.Password)
		assert.Equal(t, user.Email, userMock.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not be able to login user when user is not present", func(t *testing.T) {
		mockRepo := new(mocks.IAuthRepository)
		mockRepo.On("FindUserByUsername", userMock.Username).Return(userMock, gorm.ErrRecordNotFound)

		authService := NewAuthService(mockRepo)
		_, err := authService.LoginUser(userMock.Username)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrUserNotFound, err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("should not be able to login user when error from repo", func(t *testing.T) {
		dbError := errors.New("error from db")
		mockRepo := new(mocks.IAuthRepository)
		mockRepo.On("FindUserByUsername", userMock.Username).Return(userMock, dbError)

		authService := NewAuthService(mockRepo)
		_, err := authService.LoginUser(userMock.Username)

		assert.Error(t, err)
		assert.Equal(t, dbError.Error(), err.Error())
		mockRepo.AssertExpectations(t)
	})
}

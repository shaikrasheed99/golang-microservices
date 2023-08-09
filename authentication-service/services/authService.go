package services

import (
	"errors"
	"log"

	"github.com/shaikrasheed99/authentication-service/constants"
	"github.com/shaikrasheed99/authentication-service/models"
	"github.com/shaikrasheed99/authentication-service/repositories"
	"github.com/shaikrasheed99/authentication-service/requests"
	"gorm.io/gorm"
)

type IAuthService interface {
	SaveUser(*requests.SignupRequest) (*models.User, error)
}

type authService struct {
	ar repositories.IAuthRepository
}

func NewAuthService(ar repositories.IAuthRepository) IAuthService {
	log.Println("[NewAuthService] Creating new auth service")

	return &authService{
		ar: ar,
	}
}

func (as *authService) SaveUser(user *requests.SignupRequest) (*models.User, error) {
	log.Println("[SaveUserService] Hitting save function in auth service")

	_, err := as.ar.FindUserByUsername(user.Username)
	if err != gorm.ErrRecordNotFound {
		errMessage := constants.ErrUserAlreadyExists
		log.Println("[SaveUserService]", errMessage)
		return nil, errors.New(errMessage)
	}

	newUser := &models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}

	savedUser, err := as.ar.SaveUser(newUser)
	if err != nil {
		log.Println("[SaveUserService]", err.Error())
		return nil, err
	}

	log.Println("[SaveUserService] Returned saved user deatils from repository")
	return &savedUser, nil
}

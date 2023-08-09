package repositories

import (
	"log"

	"github.com/shaikrasheed99/authentication-service/models"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	SaveUser(*models.User) (models.User, error)
	FindUserByUsername(string) (models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	log.Println("[NewAuthRepository] Creating new auth repository")

	return &authRepository{
		db: db,
	}
}

func (ar *authRepository) SaveUser(user *models.User) (models.User, error) {
	log.Println("[SaveUserRepository] Hitting save user function in auth repository")

	res := ar.db.Create(&user)
	if res.Error != nil {
		log.Println("[SaveUserRepository]", res.Error.Error())
		return models.User{}, res.Error
	}

	log.Println("[SaveUserRepository] User details have created")
	return *user, nil
}

func (ar *authRepository) FindUserByUsername(username string) (models.User, error) {
	log.Println("[FindUserByUsernamRepository] Hitting find user by username function in auth repository")

	var user models.User
	res := ar.db.Where("username = ?", username).Find(&user)
	if res.Error != nil {
		log.Println("[FindUserByUsernamRepository]", res.Error.Error())
		return models.User{}, res.Error
	}

	if res.RowsAffected == 0 {
		log.Println("[FindUserByUsernamRepository] User is not found with username")
		return models.User{}, gorm.ErrRecordNotFound
	}

	log.Println("[FindUserByUsernamRepository] User deatils has found")
	return user, nil
}

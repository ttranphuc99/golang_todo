package services

import (
	"log"
	"os"
	"time"
	"todoapi/config"
	"todoapi/dtos"
	"todoapi/models"
	"todoapi/repository"

	"github.com/dgrijalva/jwt-go"
)

type UserAccountService interface {
	Login(user models.UserAccount) (dtos.LoginDTO, error)
}

type UserAccountServiceStruct struct {
	repository repository.UserAccountRepository
	config     config.Config
}

func NewUserAccountService(repository repository.UserAccountRepository, config config.Config) *UserAccountServiceStruct {
	service := &UserAccountServiceStruct{}
	service.repository = repository
	service.config = config

	return service
}

func (service *UserAccountServiceStruct) Login(user models.UserAccount) (dtos.LoginDTO, error) {
	userRes, error := service.repository.Login(user)

	if error != nil {
		log.Println(error)
		return dtos.LoginDTO{}, error
	}

	token, error := createToken(user.LoginId, userRes.Role, service.config)

	if error != nil {
		log.Println(error)
		return dtos.LoginDTO{}, error
	}

	return dtos.LoginDTO{Token: token, User: userRes}, nil
}

func createToken(loginId string, role int, config config.Config) (string, error) {
	var err error

	// create access token
	os.Setenv("ACCESS_SECRET", config.SecretKeyJwt)

	atClaims := jwt.MapClaims{}
	atClaims[config.TokenCurrentUserId] = loginId
	atClaims[config.TokenCurrentUserRole] = role
	atClaims[config.TokenExp] = time.Now().Add(time.Minute * 30).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

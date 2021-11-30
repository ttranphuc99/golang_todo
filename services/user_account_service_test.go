package services_test

import (
	"errors"
	"testing"
	"todoapi/models"
	"todoapi/services"
)

type MockLoginSuccessRepository struct{}

func (m MockLoginSuccessRepository) Init() error {
	return nil
}

func (m MockLoginSuccessRepository) Login(user models.UserAccount) (models.UserAccountDTO, error) {
	return models.UserAccountDTO{
		LoginId: user.LoginId,
		Role:    0,
	}, nil
}

func TestLoginSuccess(t *testing.T) {
	repo := MockLoginSuccessRepository{}
	service := services.UserAccountServiceStruct{}
	service.InitWith(repo)

	userAccount := models.UserAccount{
		LoginId: "admin",
	}

	_, error := service.Login(userAccount)

	if error != nil {
		t.Error("Expect error nil but got: ", error)
	}
}

type MockLoginFailRepository struct{}

func (m MockLoginFailRepository) Init() error {
	return nil
}

func (m MockLoginFailRepository) Login(user models.UserAccount) (models.UserAccountDTO, error) {
	return models.UserAccountDTO{}, errors.New("error")
}

func TestLoginFail(t *testing.T) {
	repo := MockLoginFailRepository{}
	service := services.UserAccountServiceStruct{}
	service.InitWith(repo)

	_, error := service.Login(models.UserAccount{})

	if error == nil {
		t.Error("Expect error but got nil")
	} else if error.Error() != "error" {
		t.Error("Expect error 'error' but got: ", error.Error())
	}
}

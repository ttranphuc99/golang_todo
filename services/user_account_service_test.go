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
	service := services.NewUserAccountService(repo, configMock)

	loginId := "admin"

	userAccount := models.UserAccount{
		LoginId: loginId,
	}

	user, error := service.Login(userAccount)

	if error != nil {
		t.Error("Expect error nil but got: ", error)
	} else if user.User.LoginId != loginId {
		t.Errorf("Expect login id = %s but got %s", loginId, user.User.LoginId)
	} else if user.Token == "" {
		t.Error("Login success but get empty token")
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
	service := services.NewUserAccountService(repo, configMock)

	_, error := service.Login(models.UserAccount{})

	if error == nil {
		t.Error("Expect error but got nil")
	} else if error.Error() != "error" {
		t.Error("Expect error 'error' but got: ", error.Error())
	}
}

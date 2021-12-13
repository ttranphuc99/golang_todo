package repository

import (
	"database/sql"
	"log"
	"todoapi/config"
	"todoapi/database"
	"todoapi/models"
)

type UserAccountRepository interface {
	Login(user models.UserAccount) (models.UserAccountDTO, error)
}

type UserAccountRepositoryStruct struct {
	dbHandler database.Database
	config    config.Config
}

func NewUserAccountRepository(dbHandler database.Database, config config.Config) (*UserAccountRepositoryStruct, error) {
	repo := &UserAccountRepositoryStruct{}
	repo.config = config
	repo.dbHandler = dbHandler
	return repo, dbHandler.Open()
}

func (repo *UserAccountRepositoryStruct) Login(user models.UserAccount) (models.UserAccountDTO, error) {
	db := repo.dbHandler.GetDb()
	row := db.QueryRow(loginSql, user.LoginId, user.Password)

	var fullname, createdTime, updatedTime string
	var updatedBy sql.NullString
	var role, status int

	error := row.Scan(&fullname, &role, &status, &createdTime, &updatedTime, &updatedBy)
	if error != nil {
		log.Println(error)
		return models.UserAccountDTO{}, error
	}

	result := models.UserAccountDTO{
		LoginId:     user.LoginId,
		Fullname:    fullname,
		Role:        role,
		Status:      status,
		CreatedTime: createdTime,
		UpdatedTime: updatedTime,
		UpdatedBy:   updatedBy.String}

	if error != nil {
		log.Println(error)
		return models.UserAccountDTO{}, error
	}

	error = db.Close()

	if error != nil {
		log.Println(error)
		return models.UserAccountDTO{}, error
	}

	return result, nil
}

const loginSql = `
SELECT fullname, role, status, created_time, updated_time, updated_by 
FROM user_account 
WHERE login_id = ? AND password = ?`

package repository

import (
	"database/sql"
	"log"
	"todoapi/database"
	"todoapi/models"
)

type UserAccountRepository interface {
	Init() error
	Login(user models.UserAccount) (models.UserAccountDTO, error)
}

type UserAccountRepositoryStruct struct {
	dbHandler database.Database
}

func (repo *UserAccountRepositoryStruct) Init() error {
	tempDb := &database.DatabaseStruct{}
	repo.dbHandler = tempDb
	error := repo.dbHandler.Open()

	return error
}

func (repo *UserAccountRepositoryStruct) Login(user models.UserAccount) (models.UserAccountDTO, error) {
	db := repo.dbHandler.GetDb()
	row, error := db.Query(loginSql, user.LoginId, user.Password)

	if error != nil {
		log.Panicln(error)
		return models.UserAccountDTO{}, error
	}

	var result models.UserAccountDTO

	for row.Next() {
		var fullname, createdTime, updatedTime string
		var updatedBy sql.NullString
		var role, status int

		error = row.Scan(&fullname, &role, &status, &createdTime, &updatedTime, &updatedBy)

		if error != nil {
			log.Panicln(error)
			return models.UserAccountDTO{}, error
		}

		result = models.UserAccountDTO{user.LoginId, fullname, role, status, createdTime, updatedTime, updatedBy.String}
	}

	error = db.Close()

	if error != nil {
		log.Panicln(error)
		return models.UserAccountDTO{}, error
	}

	return result, nil
}

const loginSql = `
SELECT fullname, role, status, created_time, updated_time, updated_by 
FROM user_account 
WHERE login_id = ? AND password = ?`

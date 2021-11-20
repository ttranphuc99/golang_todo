package repository

import (
	"log"
	"todoapi/database"
	"todoapi/models"
)

type TodoRepository interface {
	GetAllTodo() ([]models.Todo, error)
	Init() error
}

type TodoRepositoryStruct struct {
	dbHandler database.Database
}

func (repo *TodoRepositoryStruct) Init() error {
	tempDb := &database.DatabaseStruct{}
	repo.dbHandler = tempDb
	error := repo.dbHandler.Open()
	return error
}

func (repo *TodoRepositoryStruct) GetAllTodo() ([]models.Todo, error) {
	db := repo.dbHandler.GetDb()
	rows, _ := db.Query(getAllTodoSql)

	var todoLst []models.Todo

	for rows.Next() {
		var id int64
		var title, content string
		var status int
		var owner_id string
		var updated_time, created_time string

		error := rows.Scan(&id, &title, &content, &status, &owner_id, &created_time, &updated_time)

		if error != nil {
			log.Panicln(error)
			return nil, error
		}

		todo := models.Todo{id, title, content, status, owner_id, created_time, updated_time}
		todoLst = append(todoLst, todo)
	}

	error := db.Close()

	if error != nil {
		return nil, error
	}

	return todoLst, nil
}

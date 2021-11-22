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
	rows, error := db.Query(getAllTodoSql)

	if error != nil {
		log.Panic(error)
		return nil, error
	}

	var todoLst []models.Todo

	for rows.Next() {
		var id int64
		var title, content string
		var status int
		var owner string
		var updated_time, created_time string

		error := rows.Scan(&id, &title, &content, &status, &owner, &created_time, &updated_time)

		if error != nil {
			log.Panicln(error)
			return nil, error
		}

		todo := models.Todo{id, title, content, status, owner, created_time, updated_time}
		todoLst = append(todoLst, todo)
	}

	error = db.Close()

	if error != nil {
		return nil, error
	}

	return todoLst, nil
}

const insertTodoSql = `
INSERT INTO 
todo(title, content, status, owner_id, created_time, updated_time) 
VALUES (?, ?, ?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())
`
const updateTodoSql = `
UPDATE todo
SET
title = ?,
content = ?,
status = ?,
updated_date = UTC_TIMESTAMP(),
WHERE id = ?
`

const getAllTodoByOwnerSql = `
SELECT id, title, content, status, owner_id, created_time, updated_time 
FROM todo WHERE owner_id = ?
`

const getAllTodoSql = `
SELECT id, title, content, status, owner_id, created_time, updated_time 
FROM todo
`

const getAllTodoByOwnerAndStatusSql = `
SELECT * FROM todo WHERE owner_id = ? AND status = ?
`

package repository

import (
	"log"
	"todoapi/database"
	"todoapi/models"
)

type TodoRepository interface {
	GetAllTodo() ([]models.Todo, error)
	Init() error
	InsertTodo(todo *models.Todo) (models.Todo, error)
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
		var updatedTime, createdTime string

		error := rows.Scan(&id, &title, &content, &status, &owner, &createdTime, &updatedTime)

		if error != nil {
			log.Panicln(error)
			return nil, error
		}

		todo := models.Todo{
			ID:          id,
			Title:       title,
			Content:     content,
			Status:      status,
			OwnerId:     owner,
			CreatedTime: createdTime,
			UpdateTime:  updatedTime,
		}
		todoLst = append(todoLst, todo)
	}

	error = db.Close()

	if error != nil {
		return nil, error
	}

	return todoLst, nil
}

func (repo *TodoRepositoryStruct) InsertTodo(todo *models.Todo) (models.Todo, error) {
	db := repo.dbHandler.GetDb()
	sqlResult, error := db.Exec(insertTodoSql, todo.Title, todo.Content, todo.Status, todo.OwnerId)

	if error != nil {
		log.Panicln(error)
		return models.Todo{}, error
	}

	todo.ID, error = sqlResult.LastInsertId()

	if error != nil {
		log.Panicln(error)
		return models.Todo{}, error
	}

	return repo.GetByID(todo.OwnerId, todo.ID)
}

func (repo *TodoRepositoryStruct) GetByID(ownerId string, id int64) (models.Todo, error) {
	db := repo.dbHandler.GetDb()
	row := db.QueryRow(getByIDAndOwnerSql, ownerId, id)

	var title, content string
	var status int
	var updatedTime, createdTime string

	error := row.Scan(&id, &title, &content, &status, &ownerId, &createdTime, &updatedTime)

	if error != nil {
		log.Panicln(error)
		return models.Todo{}, error
	}

	todo := models.Todo{
		ID:          id,
		Title:       title,
		Content:     content,
		Status:      status,
		OwnerId:     ownerId,
		CreatedTime: createdTime,
		UpdateTime:  updatedTime,
	}
	return todo, nil
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

const getByIDAndOwnerSql = `
SELECT id, title, content, status, owner_id, created_time, updated_time
FROM todo
WHERE owner_id = ? AND id = ?
`

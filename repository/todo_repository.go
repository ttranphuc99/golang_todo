package repository

import (
	"database/sql"
	"log"
	"todoapi/database"
	"todoapi/models"
)

type TodoRepository interface {
	GetAllTodo() ([]models.Todo, error)
	Init() error
	InsertTodo(todo *models.Todo) (models.Todo, error)
	GetTodoByIDAndOwner(id int64, owner string) (models.Todo, error)
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
		var title, content sql.NullString
		var status sql.NullInt16
		var owner sql.NullString
		var updatedTime, createdTime sql.NullString

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
	// db.Close()

	return repo.GetTodoByIDAndOwner(todo.ID, todo.OwnerId.String)
}

func (repo *TodoRepositoryStruct) GetTodoByIDAndOwner(id int64, owner string) (models.Todo, error) {
	db := repo.dbHandler.GetDb()
	row := db.QueryRow(getByIDAndOwnerSql, owner, id)

	var title, content, ownerId sql.NullString
	var status sql.NullInt16
	var updatedTime, createdTime sql.NullString

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

	db.Close()
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

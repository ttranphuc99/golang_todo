package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"todoapi/config"
	"todoapi/database"
	"todoapi/models"
)

type TodoRepository interface {
	GetAllTodo() ([]models.Todo, error)
	GetAllTodoByStatus(status int) ([]models.Todo, error)
	GetAllTodoByOwnerId(ownerId string) ([]models.Todo, error)
	GetAllTodoByOwnerIdAndStatus(ownerId string, statusReq int) ([]models.Todo, error)
	InsertTodo(todo *models.Todo) (models.Todo, error)
	UpdateTodo(todo *models.Todo) (models.Todo, error)
	GetTodoByIDAndOwner(id int64, owner string) (models.Todo, error)
	GetTodoByID(id int64) (models.Todo, error)
	CloseConnection() error
	DeleteTodoById(id int64) error
}

type TodoRepositoryStruct struct {
	dbHandler database.Database
	config    config.Config
}

func NewTodoRepository(dbHandler database.Database, config config.Config) (*TodoRepositoryStruct, error) {
	repo := &TodoRepositoryStruct{
		dbHandler: dbHandler,
		config:    config,
	}
	return repo, dbHandler.Open()
}

func (repo *TodoRepositoryStruct) CloseConnection() error {
	error := repo.dbHandler.GetDb().Close()

	if error != nil {
		log.Println(error)
		return error
	}
	return nil
}

func (repo *TodoRepositoryStruct) GetAllTodo() ([]models.Todo, error) {
	db := repo.dbHandler.GetDb()
	rows, error := db.Query(getAllTodoSql)

	if error != nil {
		log.Panic(error)
		return nil, error
	}

	return getResultTodoList(rows)
}

func (repo *TodoRepositoryStruct) GetAllTodoByStatus(status int) ([]models.Todo, error) {
	db := repo.dbHandler.GetDb()
	rows, error := db.Query(getAllTodoByStatusSql, status)

	if error != nil {
		log.Panic(error)
		return nil, error
	}

	return getResultTodoList(rows)
}

func (repo *TodoRepositoryStruct) GetAllTodoByOwnerId(ownerId string) ([]models.Todo, error) {
	db := repo.dbHandler.GetDb()
	rows, error := db.Query(getAllTodoByOwnerSql, ownerId)

	if error != nil {
		log.Panic(error)
		return nil, error
	}

	return getResultTodoList(rows)
}

func (repo *TodoRepositoryStruct) GetAllTodoByOwnerIdAndStatus(ownerId string, statusReq int) ([]models.Todo, error) {
	db := repo.dbHandler.GetDb()
	rows, error := db.Query(getAllTodoByOwnerAndStatusSql, ownerId, statusReq)

	if error != nil {
		log.Panic(error)
		return nil, error
	}

	return getResultTodoList(rows)
}

func getResultTodoList(rows *sql.Rows) ([]models.Todo, error) {
	todoLst := []models.Todo{}

	for rows.Next() {
		var id int64
		var title, content sql.NullString
		var status sql.NullInt16
		var owner sql.NullString
		var updatedTime, createdTime sql.NullString

		error := rows.Scan(&id, &title, &content, &status, &owner, &createdTime, &updatedTime)

		if error != nil {
			log.Println(error)
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
	return todoLst, nil
}

func (repo *TodoRepositoryStruct) InsertTodo(todo *models.Todo) (models.Todo, error) {
	db := repo.dbHandler.GetDb()
	sqlResult, error := db.Exec(insertTodoSql, todo.Title, todo.Content, todo.Status, todo.OwnerId)

	if error != nil {
		log.Println(error)
		return models.Todo{}, error
	}

	todo.ID, error = sqlResult.LastInsertId()

	if error != nil {
		log.Println(error)
		return models.Todo{}, error
	}

	return repo.GetTodoByIDAndOwner(todo.ID, todo.OwnerId.String)
}

func (repo *TodoRepositoryStruct) UpdateTodo(todo *models.Todo) (models.Todo, error) {
	db := repo.dbHandler.GetDb()
	sqlResult, error := db.Exec(updateTodoSql, todo.Title, todo.Content, todo.Status, todo.ID)

	if error != nil {
		log.Println(error)
		return models.Todo{}, error
	}

	rowEffected, error := sqlResult.RowsAffected()

	if error != nil {
		log.Println(error)
		return models.Todo{}, error
	} else if rowEffected != 1 {
		log.Println("row affected is " + fmt.Sprintf("%d", rowEffected))
		return models.Todo{}, errors.New("row affected is " + fmt.Sprintf("%d", rowEffected))
	}

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
		log.Println(error)
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

func (repo *TodoRepositoryStruct) GetTodoByID(id int64) (models.Todo, error) {
	db := repo.dbHandler.GetDb()
	row := db.QueryRow(getByID, id)

	var title, content, ownerId sql.NullString
	var status sql.NullInt16
	var updatedTime, createdTime sql.NullString

	error := row.Scan(&id, &title, &content, &status, &ownerId, &createdTime, &updatedTime)

	if error != nil {
		log.Println(error)
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

func (repo *TodoRepositoryStruct) DeleteTodoById(id int64) error {
	db := repo.dbHandler.GetDb()
	sqlResult, error := db.Exec(deleteTodo, id)

	if error != nil {
		log.Println(error)
		return error
	}

	rowAffect, error := sqlResult.RowsAffected()

	if error != nil {
		log.Println(error)
		return error
	}

	if rowAffect == 0 {
		msg := fmt.Sprintf("Delete failed with ID = %v.", id)
		log.Println(msg)
		return errors.New(msg)
	}

	return nil
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
updated_time = UTC_TIMESTAMP()
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

const getAllTodoByStatusSql = `
SELECT id, title, content, status, owner_id, created_time, updated_time 
FROM todo 
WHERE status = ?
`

const getAllTodoByOwnerAndStatusSql = `
SELECT id, title, content, status, owner_id, created_time, updated_time
FROM todo 
WHERE owner_id = ? AND status = ?
`

const getByIDAndOwnerSql = `
SELECT id, title, content, status, owner_id, created_time, updated_time
FROM todo
WHERE owner_id = ? AND id = ?
`

const getByID = `
SELECT id, title, content, status, owner_id, created_time, updated_time
FROM todo
WHERE id = ?
`

const deleteTodo = `
DELETE FROM todo
WHERE id = ?
`

package mock

import (
	"database/sql"
	"todoapi/models"
)

var mockTodo = models.Todo{
	ID: 1,
}

type MockTodoRepositorySuccess struct{}

func (repo *MockTodoRepositorySuccess) Init() error {
	return nil
}

func (repo *MockTodoRepositorySuccess) CloseConnection() error {
	return nil
}

func (repo *MockTodoRepositorySuccess) GetAllTodo() ([]models.Todo, error) {
	return []models.Todo{mockTodo}, nil
}

func (repo *MockTodoRepositorySuccess) GetAllTodoByStatus(status int) ([]models.Todo, error) {
	return []models.Todo{mockTodo}, nil
}

func (repo *MockTodoRepositorySuccess) GetAllTodoByOwnerId(ownerId string) ([]models.Todo, error) {
	return []models.Todo{mockTodo}, nil
}

func (repo *MockTodoRepositorySuccess) GetAllTodoByOwnerIdAndStatus(ownerId string, statusReq int) ([]models.Todo, error) {
	return []models.Todo{mockTodo}, nil
}

func (repo *MockTodoRepositorySuccess) InsertTodo(todo *models.Todo) (models.Todo, error) {
	return *todo, nil
}

func (repo *MockTodoRepositorySuccess) UpdateTodo(todo *models.Todo) (models.Todo, error) {
	return *todo, nil
}

func (repo *MockTodoRepositorySuccess) GetTodoByIDAndOwner(id int64, owner string) (models.Todo, error) {
	return models.Todo{ID: id, OwnerId: sql.NullString{String: owner}}, nil
}

func (repo *MockTodoRepositorySuccess) GetTodoByID(id int64) (models.Todo, error) {
	return models.Todo{ID: id}, nil
}

func (repo *MockTodoRepositorySuccess) DeleteTodoById(id int64) error {
	return nil
}

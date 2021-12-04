package mock

import (
	"errors"
	"todoapi/models"
)

type MockTodoRepositoryError struct{}

func (repo *MockTodoRepositoryError) Init() error {
	return nil
}

func (repo *MockTodoRepositoryError) CloseConnection() error {
	return nil
}

func (repo *MockTodoRepositoryError) GetAllTodo() ([]models.Todo, error) {
	return nil, errors.New("Test error")
}

func (repo *MockTodoRepositoryError) GetAllTodoByStatus(status int) ([]models.Todo, error) {
	return nil, errors.New("Test error")
}

func (repo *MockTodoRepositoryError) GetAllTodoByOwnerId(ownerId string) ([]models.Todo, error) {
	return nil, errors.New("Test error")
}

func (repo *MockTodoRepositoryError) GetAllTodoByOwnerIdAndStatus(ownerId string, statusReq int) ([]models.Todo, error) {
	return nil, errors.New("Test error")
}

func (repo *MockTodoRepositoryError) InsertTodo(todo *models.Todo) (models.Todo, error) {
	return models.Todo{}, errors.New("Test error")
}

func (repo *MockTodoRepositoryError) UpdateTodo(todo *models.Todo) (models.Todo, error) {
	return models.Todo{}, errors.New("Test error")
}

func (repo *MockTodoRepositoryError) GetTodoByIDAndOwner(id int64, owner string) (models.Todo, error) {
	return models.Todo{}, errors.New("Test error")
}

func (repo *MockTodoRepositoryError) GetTodoByID(id int64) (models.Todo, error) {
	return models.Todo{}, errors.New("Test error")
}

func (repo *MockTodoRepositoryError) DeleteTodoById(id int64) error {
	return nil
}

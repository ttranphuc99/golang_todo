package services

import (
	"errors"
	"todoapi/models"
	"todoapi/models/constants"
	"todoapi/repository"
)

var todos = []models.Todo{
	{ID: 1, Content: "11111", Status: constants.TodoStatusActive},
	{ID: 2, Content: "2222", Status: constants.TodoStatusCompleted},
	{ID: 3, Content: "3333", Status: constants.TodoStatusActive},
}

type TodoService interface {
	GetAllTodo(status int64) []models.Todo
	InsertTodo(newTodo models.Todo) (models.Todo, error)
	GetTodoByID(id int64) (models.Todo, error)
	UpdateTodo(newTodo models.Todo) (models.Todo, error)
	Init() error
}

type TodoServiceStruct struct {
	repository repository.TodoRepository
}

func (service *TodoServiceStruct) Init() error {
	tempRepo := &repository.TodoRepositoryStruct{}
	service.repository = tempRepo
	return service.repository.Init()
}

func (service *TodoServiceStruct) GetAllTodo(status int64) []models.Todo {
	if status != constants.TodoStatusAll {
		todosRes := []models.Todo{}

		for _, todo := range todos {
			if status == int64(todo.Status) {
				todosRes = append(todosRes, todo)
			}
		}

		return todosRes
	}

	result, _ := service.repository.GetAllTodo()
	return result
}

func (t *TodoServiceStruct) InsertTodo(newTodo models.Todo) (models.Todo, error) {
	_, err := findById(newTodo.ID)

	if err == nil {
		return models.Todo{}, errors.New("duplicate id")
	}

	todos = append(todos, newTodo)
	return newTodo, nil
}

func (t *TodoServiceStruct) GetTodoByID(id int64) (models.Todo, error) {
	todo, err := findById(id)

	if err != nil {
		return models.Todo{}, err
	} else {
		return *todo, nil
	}
}

func (t *TodoServiceStruct) UpdateTodo(newTodo models.Todo) (models.Todo, error) {
	oldTodo, err := findById(newTodo.ID)

	if err != nil {
		return models.Todo{}, err
	}
	oldTodo.Content = newTodo.Content
	oldTodo.Status = newTodo.Status

	return *oldTodo, nil
}

func findById(id int64) (todo *models.Todo, e error) {
	for idx, todo := range todos {
		if id == todo.ID {
			return &todos[idx], nil
		}
	}
	return &models.Todo{}, errors.New("not found with id " + string(id))
}

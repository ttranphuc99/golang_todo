package services

import (
	"errors"
	"log"
	"todoapi/dtos"
	todomapper "todoapi/mapper/todo_mapper"
	"todoapi/models"
	"todoapi/models/constants"
	"todoapi/repository"
)

var todos = []models.Todo{}

type TodoService interface {
	GetAllTodo(status int64) ([]dtos.TodoDTO, error)
	InsertTodo(todoDTO *dtos.TodoDTO) (dtos.TodoDTO, error)
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

func (service *TodoServiceStruct) GetAllTodo(status int64) ([]dtos.TodoDTO, error) {
	if status != constants.TodoStatusAll {
		todosRes := []models.Todo{}

		for _, todo := range todos {
			if status == int64(todo.Status.Int16) {
				todosRes = append(todosRes, todo)
			}
		}

		return todomapper.ToDTOs(todosRes), nil
	}

	result, err := service.repository.GetAllTodo()

	if err != nil {
		return nil, err
	}
	return todomapper.ToDTOs(result), nil
}

func (service *TodoServiceStruct) InsertTodo(todoDTO *dtos.TodoDTO) (dtos.TodoDTO, error) {
	todoModel := todomapper.ToModel(*todoDTO)
	todoRes, error := service.repository.InsertTodo(&todoModel)

	if error != nil {
		log.Panic(error)
		return dtos.TodoDTO{}, error
	}

	return todomapper.ToDTO(todoRes), nil
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

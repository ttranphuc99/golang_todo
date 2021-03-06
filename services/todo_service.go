package services

import (
	"log"
	"todoapi/config"
	"todoapi/dtos"
	todomapper "todoapi/mapper/todo_mapper"
	"todoapi/models"
	"todoapi/models/constants"
	"todoapi/repository"
)

type TodoService interface {
	GetAllTodo(status int, ownerId string, role int) ([]dtos.TodoDTO, error)
	InsertTodo(todoDTO *dtos.TodoDTO) (dtos.TodoDTO, error)
	GetTodoByIDAndOwner(id int64, ownerId string) (dtos.TodoDTO, error)
	GetTodoByID(id int64) (dtos.TodoDTO, error)
	UpdateTodo(newTodo dtos.TodoDTO) (dtos.TodoDTO, error)
	DeleteTodo(id int64, ownerId string) error
}

type TodoServiceStruct struct {
	repository repository.TodoRepository
	config     config.Config
}

func NewTodoServiceStruct(repository repository.TodoRepository, config config.Config) *TodoServiceStruct {
	return &TodoServiceStruct{
		repository: repository,
		config:     config,
	}
}

func (service *TodoServiceStruct) GetAllTodo(status int, ownerId string, role int) ([]dtos.TodoDTO, error) {
	var result []models.Todo
	var err error

	if status != constants.TodoStatusAll {
		if role != constants.RoleAdmin {
			result, err = service.repository.GetAllTodoByOwnerIdAndStatus(ownerId, status)
		} else {
			result, err = service.repository.GetAllTodoByStatus(status)
		}
	} else {
		if role != constants.RoleAdmin {
			result, err = service.repository.GetAllTodoByOwnerId(ownerId)
		} else {
			result, err = service.repository.GetAllTodo()
		}
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return todomapper.ToDTOs(result), service.repository.CloseConnection()
}

func (service *TodoServiceStruct) InsertTodo(todoDTO *dtos.TodoDTO) (dtos.TodoDTO, error) {
	todoModel := todomapper.ToModel(*todoDTO)
	todoRes, error := service.repository.InsertTodo(&todoModel)

	if error != nil {
		log.Println(error)
		return dtos.TodoDTO{}, error
	}

	return todomapper.ToDTO(todoRes), service.repository.CloseConnection()
}

func (service *TodoServiceStruct) GetTodoByIDAndOwner(id int64, ownerId string) (dtos.TodoDTO, error) {
	resultTodo, error := service.repository.GetTodoByIDAndOwner(id, ownerId)

	if error != nil {
		log.Println(error)
		return dtos.TodoDTO{}, error
	}

	return todomapper.ToDTO(resultTodo), service.repository.CloseConnection()
}

func (service *TodoServiceStruct) GetTodoByID(id int64) (dtos.TodoDTO, error) {
	resultTodo, error := service.repository.GetTodoByID(id)

	if error != nil {
		log.Println(error)
		return dtos.TodoDTO{}, error
	}

	return todomapper.ToDTO(resultTodo), service.repository.CloseConnection()
}

func (service *TodoServiceStruct) UpdateTodo(newTodo dtos.TodoDTO) (dtos.TodoDTO, error) {
	oldTodo, err := service.repository.GetTodoByIDAndOwner(newTodo.ID, newTodo.OwnerId)

	if err != nil {
		return dtos.TodoDTO{}, err
	}

	oldTodo.Title.String = newTodo.Title
	oldTodo.Content.String = newTodo.Content
	oldTodo.Status.Int16 = newTodo.Status

	updatedTodo, error := service.repository.UpdateTodo(&oldTodo)

	if error != nil {
		return dtos.TodoDTO{}, error
	}

	return todomapper.ToDTO(updatedTodo), service.repository.CloseConnection()
}

func (service *TodoServiceStruct) DeleteTodo(id int64, ownerId string) error {
	_, error := service.repository.GetTodoByIDAndOwner(id, ownerId)

	if error != nil {
		log.Println(error)
		return error
	}

	error = service.repository.DeleteTodoById(id)

	if error != nil {
		log.Println(error)
		return error
	}

	return service.repository.CloseConnection()
}

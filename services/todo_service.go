package services

import (
	"log"
	"todoapi/dtos"
	todomapper "todoapi/mapper/todo_mapper"
	"todoapi/repository"
)

type TodoService interface {
	GetAllTodo(status int64) ([]dtos.TodoDTO, error)
	InsertTodo(todoDTO *dtos.TodoDTO) (dtos.TodoDTO, error)
	GetTodoByID(id int64, ownerId string) (dtos.TodoDTO, error)
	UpdateTodo(newTodo dtos.TodoDTO) (dtos.TodoDTO, error)
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
	// if status != constants.TodoStatusAll {
	// 	todosRes := []models.Todo{}

	// 	for _, todo := range todos {
	// 		if status == int64(todo.Status.Int16) {
	// 			todosRes = append(todosRes, todo)
	// 		}
	// 	}

	// 	return todomapper.ToDTOs(todosRes), nil
	// }

	result, err := service.repository.GetAllTodo()

	if err != nil {
		return nil, err
	}

	return todomapper.ToDTOs(result), service.repository.CloseConnection()
}

func (service *TodoServiceStruct) InsertTodo(todoDTO *dtos.TodoDTO) (dtos.TodoDTO, error) {
	todoModel := todomapper.ToModel(*todoDTO)
	todoRes, error := service.repository.InsertTodo(&todoModel)

	if error != nil {
		log.Panic(error)
		return dtos.TodoDTO{}, error
	}

	return todomapper.ToDTO(todoRes), service.repository.CloseConnection()
}

func (service *TodoServiceStruct) GetTodoByID(id int64, ownerId string) (dtos.TodoDTO, error) {
	resultTodo, error := service.repository.GetTodoByIDAndOwner(id, ownerId)

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

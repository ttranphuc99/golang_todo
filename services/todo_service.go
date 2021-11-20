package services

import (
	"errors"
	"todoapi/models"
	"todoapi/models/constants"
)

var todos = []models.Todo{
	{ID: 1, Content: "11111", Status: constants.TodoStatusActive},
	{ID: 2, Content: "2222", Status: constants.TodoStatusCompleted},
	{ID: 3, Content: "3333", Status: constants.TodoStatusActive},
}

func GetAllTodo(status int64) []models.Todo {
	if status != constants.TodoStatusAll {
		todosRes := []models.Todo{}

		for _, todo := range todos {
			if status == int64(todo.Status) {
				todosRes = append(todosRes, todo)
			}
		}

		return todosRes
	}

	return todos
}

func InsertTodo(newTodo models.Todo) (models.Todo, error) {
	_, err := findById(newTodo.ID)

	if err == nil {
		return models.Todo{}, errors.New("duplicate id")
	}

	todos = append(todos, newTodo)
	return newTodo, nil
}

func GetTodoByID(id int64) (models.Todo, error) {
	todo, err := findById(id)

	if err != nil {
		return models.Todo{}, err
	} else {
		return *todo, nil
	}
}

func UpdateTodo(newTodo models.Todo) (models.Todo, error) {
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

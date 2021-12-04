package services_test

import (
	"testing"
	"todoapi/dtos"
	"todoapi/models/constants"
	"todoapi/repository/mock"
	"todoapi/services"
)

// get all todo with all status and admin role
func TestGetAllTodoError1(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetAllTodo(constants.TodoStatusAll, "admin", constants.RoleAdmin)

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get all todo with all status and user role
func TestGetAllTodoError2(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetAllTodo(constants.TodoStatusAll, "admin", constants.RoleUser)

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get all todo with 1 status and admin role
func TestGetAllTodoError3(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetAllTodo(constants.TodoStatusActive, "admin", constants.RoleAdmin)

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get all todo with 1 status and user role
func TestGetAllTodoError4(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetAllTodo(constants.TodoStatusActive, "admin", constants.RoleUser)

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// insert todo
func TestInsertTodoError(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.InsertTodo(&dtos.TodoDTO{})

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get todo by id and owner
func TestGetTodoByIDAndOwnerError(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetTodoByIDAndOwner(0, "admin")

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestUpdateTodoError(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.UpdateTodo(dtos.TodoDTO{})

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestDeleteTodoError(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	error := service.DeleteTodo(0, "")

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get all todo with all status and admin role
func TestGetAllTodoSuccess1(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetAllTodo(constants.TodoStatusAll, "admin", constants.RoleAdmin)

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	}
}

// get all todo with all status and user role
func TestGetAllTodoSuccess2(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetAllTodo(constants.TodoStatusAll, "admin", constants.RoleUser)

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	}
}

// get all todo with 1 status and admin role
func TestGetAllTodoSuccess3(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetAllTodo(constants.TodoStatusActive, "admin", constants.RoleAdmin)

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	}
}

// get all todo with 1 status and user role
func TestGetAllTodoSuccess4(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetAllTodo(constants.TodoStatusActive, "admin", constants.RoleUser)

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	}
}

// insert todo
func TestInsertTodoSuccess(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.InsertTodo(&dtos.TodoDTO{})

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	}
}

// get todo by id and owner
func TestGetTodoByIDAndOwnerSuccess(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.GetTodoByIDAndOwner(0, "admin")

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	}
}

func TestUpdateTodoSuccess(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	_, error := service.UpdateTodo(dtos.TodoDTO{})

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	}
}

func TestDeleteTodoSuccess(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	error := service.DeleteTodo(0, "")

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	}
}

func TestInit(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)

	service.Init()
}

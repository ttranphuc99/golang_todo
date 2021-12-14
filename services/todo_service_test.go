package services_test

import (
	"testing"
	"todoapi/config"
	"todoapi/dtos"
	"todoapi/models/constants"
	"todoapi/repository/mock"
	"todoapi/services"
)

var configMock = config.NewMockConfig()

// get all todo with all status and admin role
func TestGetAllTodoError1(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	_, error := service.GetAllTodo(constants.TodoStatusAll, "admin", constants.RoleAdmin)

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get all todo with all status and user role
func TestGetAllTodoError2(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	_, error := service.GetAllTodo(constants.TodoStatusAll, "admin", constants.RoleUser)

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get all todo with 1 status and admin role
func TestGetAllTodoError3(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	_, error := service.GetAllTodo(constants.TodoStatusActive, "admin", constants.RoleAdmin)

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get all todo with 1 status and user role
func TestGetAllTodoError4(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	_, error := service.GetAllTodo(constants.TodoStatusActive, "admin", constants.RoleUser)

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// insert todo
func TestInsertTodoError(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	_, error := service.InsertTodo(&dtos.TodoDTO{})

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get todo by id and owner
func TestGetTodoByIDAndOwnerError(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	_, error := service.GetTodoByIDAndOwner(0, "admin")

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestUpdateTodoError(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	_, error := service.UpdateTodo(dtos.TodoDTO{})

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestDeleteTodoError(t *testing.T) {
	repo := mock.MockTodoRepositoryError{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	error := service.DeleteTodo(0, "")

	if error == nil {
		t.Error("Expect error but got nil")
	}
}

// get all todo with all status and admin role
func TestGetAllTodoSuccess1(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	listTodo, error := service.GetAllTodo(constants.TodoStatusAll, "admin", constants.RoleAdmin)

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	} else if len(listTodo) != 1 {
		t.Errorf("Expect list have 1 element but get %d", len(listTodo))
	}
}

// get all todo with all status and user role
func TestGetAllTodoSuccess2(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	listTodo, error := service.GetAllTodo(constants.TodoStatusAll, "admin", constants.RoleUser)

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	} else if len(listTodo) != 1 {
		t.Errorf("Expect list have 1 element but get %d", len(listTodo))
	}
}

// get all todo with 1 status and admin role
func TestGetAllTodoSuccess3(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	listTodo, error := service.GetAllTodo(constants.TodoStatusActive, "admin", constants.RoleAdmin)

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	} else if len(listTodo) != 1 {
		t.Errorf("Expect list have 1 element but get %d", len(listTodo))
	}
}

// get all todo with 1 status and user role
func TestGetAllTodoSuccess4(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	listTodo, error := service.GetAllTodo(constants.TodoStatusActive, "admin", constants.RoleUser)

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	} else if len(listTodo) != 1 {
		t.Errorf("Expect list have 1 element but get %d", len(listTodo))
	}
}

// insert todo
func TestInsertTodoSuccess(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	var todoId int64 = 1

	todo, error := service.InsertTodo(&dtos.TodoDTO{ID: todoId})

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	} else if todo.ID != todoId {
		t.Errorf("Expect todo with ID %d but get ID %d", todoId, todo.ID)
	}
}

// get todo by id and owner
func TestGetTodoByIDAndOwnerSuccess(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	var id int64 = 1
	owner := "admin"
	todo, error := service.GetTodoByIDAndOwner(id, owner)

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	} else if todo.ID != id || todo.OwnerId != owner {
		t.Errorf("Expect todo with ID %d and Owner %s but get ID %d and Owner %s", id, owner, todo.ID, todo.OwnerId)
	}
}

func TestUpdateTodoSuccess(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	var todoId int64 = 1
	todo, error := service.UpdateTodo(dtos.TodoDTO{ID: todoId})

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	} else if todo.ID != todoId {
		t.Errorf("Expect todo with ID %d but get ID %d", todoId, todo.ID)
	}
}

func TestDeleteTodoSuccess(t *testing.T) {
	repo := mock.MockTodoRepositorySuccess{}
	service := services.NewTodoServiceStruct(&repo, configMock)

	error := service.DeleteTodo(0, "")

	if error != nil {
		t.Error("Expect error nil but got error: ", error)
	}
}

package todomapper

import (
	"database/sql"
	"todoapi/dtos"
	"todoapi/models"
)

func ToModel(dto dtos.TodoDTO) models.Todo {
	return models.Todo{
		ID:          dto.ID,
		Title:       sql.NullString{String: dto.Title, Valid: true},
		Content:     sql.NullString{String: dto.Content, Valid: true},
		Status:      sql.NullInt16{Int16: dto.Status, Valid: true},
		OwnerId:     sql.NullString{String: dto.OwnerId, Valid: true},
		CreatedTime: sql.NullString{String: dto.CreatedTime, Valid: true},
		UpdateTime:  sql.NullString{String: dto.UpdateTime, Valid: true},
	}
}

func ToDTO(model models.Todo) dtos.TodoDTO {
	return dtos.TodoDTO{
		ID:          model.ID,
		Title:       model.Title.String,
		Content:     model.Content.String,
		Status:      model.Status.Int16,
		OwnerId:     model.OwnerId.String,
		CreatedTime: model.CreatedTime.String,
		UpdateTime:  model.UpdateTime.String,
	}
}

func ToDTOs(models []models.Todo) []dtos.TodoDTO {
	var dtos []dtos.TodoDTO
	for _, model := range models {
		dtos = append(dtos, ToDTO(model))
	}
	return dtos
}

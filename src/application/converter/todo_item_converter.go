package converter

import (
	"github.com/go-rest-api/src/application/dto"
	"github.com/go-rest-api/src/application/form"
	"github.com/go-rest-api/src/domain"
)

func ConvertToDtdto(model domain.TodoItem) dto.TodoItemDto {
	return dto.TodoItemDto{
		Id:          model.Id,
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
	}
}

func ConvertCreateToModel(request form.TodoItemCreate) domain.TodoItem {
	return domain.TodoItem{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	}
}

func ConvertUpdateToModel(request form.TodoItemUpdate) domain.TodoItem {
	return domain.TodoItem{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	}
}

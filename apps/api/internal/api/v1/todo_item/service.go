package todo_item

import (
	"fmt"
	"todo_api/internal/repository"
	"todo_api/internal/types"
)

type TodoItemService struct {
	repository *repository.Repository
}

func SetupService(service *types.Service) *TodoItemService {
	return &TodoItemService{
		repository: service.Repository,
	}
}

func (h *TodoItemService) GetByID(id uint64) (*TodoItem, error) {
	model, okModel := h.repository.TodoItem.GetById(id)
	if okModel != nil {
		return nil, okModel
	}

	return &TodoItem{
		ID:          fmt.Sprintf("%d", model.ID),
		ListID:      fmt.Sprintf("%d", model.ListID),
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}, nil
}

func (h *TodoItemService) Create(list_id uint64, title string, description string, status string) (*TodoItem, error) {

	model, okModel := h.repository.TodoItem.Create(list_id, title, description, status)
	if okModel != nil {
		return nil, fmt.Errorf("cannot create item")
	}

	return &TodoItem{
		ID:          fmt.Sprintf("%d", model.ID),
		ListID:      fmt.Sprintf("%d", model.ListID),
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}, nil
}

func (h *TodoItemService) Update(id uint64, fields map[string]interface{}) (*TodoItem, error) {
	model, okModel := h.repository.TodoItem.Update(id, fields)
	if okModel != nil {
		return nil, okModel
	}

	return &TodoItem{
		ID:          fmt.Sprintf("%d", model.ID),
		ListID:      fmt.Sprintf("%d", model.ListID),
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}, nil
}

func (h *TodoItemService) Delete(id uint64) error {
	return h.repository.TodoItem.Delete(id)
}

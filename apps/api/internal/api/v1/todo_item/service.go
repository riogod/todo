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
		ID: fmt.Sprintf("%d", model.ID),
		List: TodoList{
			ID:          fmt.Sprintf("%d", model.List.ID),
			Title:       model.List.Title,
			Description: model.List.Description,
			Status:      model.List.Status,
			CreatedAt:   model.List.CreatedAt,
			UpdatedAt:   model.List.UpdatedAt,
		},
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
		return nil, okModel
	}

	return &TodoItem{
		ID: fmt.Sprintf("%d", model.ID),
		List: TodoList{
			ID:          fmt.Sprintf("%d", model.List.ID),
			Title:       model.List.Title,
			Description: model.List.Description,
			Status:      model.List.Status,
			CreatedAt:   model.List.CreatedAt,
			UpdatedAt:   model.List.UpdatedAt,
		},
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
	fmt.Println(model)
	return &TodoItem{
		ID:    fmt.Sprintf("%d", model.ID),
		Title: model.Title,
		List: TodoList{
			ID:          fmt.Sprintf("%d", model.List.ID),
			Title:       model.List.Title,
			Description: model.List.Description,
			Status:      model.List.Status,
			CreatedAt:   model.List.CreatedAt,
			UpdatedAt:   model.List.UpdatedAt,
		},
		Description: model.Description,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}, nil
}

func (h *TodoItemService) Delete(id uint64) error {
	return h.repository.TodoItem.Delete(id)
}

func (h *TodoItemService) Search(title string) (*[]TodoItem, error) {
	var response []TodoItem
	model, okModel := h.repository.TodoItem.Search(title)
	if okModel != nil {
		return nil, okModel
	}

	for _, item := range *model {
		response = append(response, TodoItem{
			ID: fmt.Sprintf("%d", item.ID),
			List: TodoList{
				ID:          fmt.Sprintf("%d", item.List.ID),
				Title:       item.List.Title,
				Description: item.List.Description,
				Status:      item.List.Status,
				CreatedAt:   item.List.CreatedAt,
				UpdatedAt:   item.List.UpdatedAt,
			},
			Title:       item.Title,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return &response, nil

}

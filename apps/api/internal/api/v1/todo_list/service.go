package todo_list

import (
	"fmt"
	"todo_api/internal/repository"
	"todo_api/internal/types"
)

type TodoListService struct {
	repository *repository.Repository
}

func SetupService(service *types.Service) *TodoListService {
	return &TodoListService{
		repository: service.Repository,
	}
}

func (h *TodoListService) GetByID(id uint64) (*TodoList, error) {
	model, okModel := h.repository.TodoList.GetByID(id)
	if okModel != nil {
		return nil, okModel
	}

	return &TodoList{
		ID:          fmt.Sprintf("%d", model.ID),
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}, nil
}

func (h *TodoListService) GetAll() (*[]TodoList, error) {
	var response []TodoList

	model := h.repository.TodoList.GetAll()

	for _, item := range *model {
		response = append(response, TodoList{
			ID:          fmt.Sprintf("%d", item.ID),
			Title:       item.Title,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	return &response, nil
}

func (h *TodoListService) Create(title string, description string, status string) (*TodoList, error) {

	model, okModel := h.repository.TodoList.Create(title, description, status)
	if okModel != nil {
		return nil, fmt.Errorf("cannot create item")
	}

	return &TodoList{
		ID:          fmt.Sprintf("%d", model.ID),
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}, nil
}

func (h *TodoListService) Update(id uint64, fields map[string]interface{}) (*TodoList, error) {
	model, okModel := h.repository.TodoList.Update(id, fields)
	if okModel != nil {
		return nil, okModel
	}

	return &TodoList{
		ID:          fmt.Sprintf("%d", model.ID),
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}, nil
}

func (h *TodoListService) Delete(id uint64) error {
	return h.repository.TodoList.Delete(id)
}

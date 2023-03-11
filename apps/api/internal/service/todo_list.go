package service

import (
	"fmt"
	"time"
	"todo_api/internal/repository"

	model "github.com/riogod/todo/libs/gomodels"
)

type TodoListService struct {
	repository *repository.Repository
}

type RequestTodoListDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ResponseTodoListItemDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodoListService(repos *repository.Repository) *TodoListService {
	return &TodoListService{
		repository: repos,
	}
}

func (t *TodoListService) GetAllRecords() *[]ResponseTodoListItemDTO {
	var resultMap []model.ToDoItemList
	var responseMap []ResponseTodoListItemDTO
	t.repository.TodoList.GetAll(&resultMap)

	for _, item := range resultMap {
		mappedItem := ResponseTodoListItemDTO{
			ID:          fmt.Sprintf("%d", item.ID),
			Title:       item.Title,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
		responseMap = append(responseMap, mappedItem)
	}
	return &responseMap
}

func (t *TodoListService) GetByID(id string) (*model.ToDoItemList, error) {

	return t.repository.TodoList.GetByID(id)
}

func (t *TodoListService) Create(title string, description string, status string) (*model.ToDoItemList, error) {
	return t.repository.TodoList.Create(&model.ToDoItemList{
		ID:          0,
		Title:       title,
		Description: description,
		Status:      status,
	})
}

func (t *TodoListService) Update(id string, title string, description string, status string) (*model.ToDoItemList, error) {
	updateParams := model.ToDoItemList{
		ID:          0,
		Title:       title,
		Description: description,
		Status:      status,
	}

	return t.repository.TodoList.Update(id, &updateParams)
}

func (t *TodoListService) Delete(id string) error {

	return t.repository.TodoList.Delete(id)
}

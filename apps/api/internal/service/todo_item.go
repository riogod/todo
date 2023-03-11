package service

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"todo_api/internal/repository"

	model "github.com/riogod/todo/libs/gomodels"
)

type TodoItemService struct {
	repository *repository.Repository
}

type RequestTodoItemDTO struct {
	ListID      string `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ResponseTodoItemItemDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodoItemService(repos *repository.Repository) *TodoItemService {
	return &TodoItemService{
		repository: repos,
	}
}

func (t *TodoItemService) GetByID(id uint64) (*model.ToDoItem, error) {
	return t.repository.TodoItem.GetById(id)
}

func (t *TodoItemService) GetAllByListID(list_id uint64) error {
	return fmt.Errorf("not implemented")
}

func (t *TodoItemService) Create(list_id uint64, title string, description string, status string) (*model.ToDoItem, error) {
	getList, err := t.repository.TodoList.GetByID(list_id)
	if err != nil {
		return nil, fmt.Errorf("cannot find list with id=%d", list_id)
	}
	model := &model.ToDoItem{
		ID:          0,
		List:        *getList,
		Title:       title,
		Description: description,
		Status:      status,
	}
	model_err := t.repository.TodoItem.Create(model)
	if model_err != nil {
		return nil, fmt.Errorf("cannot create item")
	}

	return model, nil
}

func (t *TodoItemService) Update(id uint64, list_id string, title string, description string, status string) (*model.ToDoItem, error) {
	model, err := t.repository.TodoItem.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("not found item with id=%d", id)
	}
	fmt.Println(list_id, reflect.TypeOf(list_id))

	if list_id != "" {
		idList, listID_err := strconv.ParseUint(list_id, 10, 64)
		if listID_err != nil {
			return nil, fmt.Errorf("list_id must be an number ")
		}

		listModel, listModelErr := t.repository.TodoList.GetByID(idList)
		if listModelErr != nil {
			return nil, fmt.Errorf("cannot find list with id=%d", idList)
		}
		model.List = *listModel
		model.ListID = idList
	}
	if title != "" {
		model.Title = title
	}
	if description != "" {
		model.Description = description
	}
	if status != "" {
		model.Status = status
	}

	t.repository.TodoItem.Update(model)

	return model, nil
}

func (t *TodoItemService) Delete(id uint64) error {
	return t.repository.TodoItem.Delete(id)
}

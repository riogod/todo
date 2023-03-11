package repository

import (
	"fmt"

	model "github.com/riogod/todo/libs/gomodels"
	"gorm.io/gorm"
)

type TodoItemRepository struct {
	DB *gorm.DB
}

func NewTodoItemRepository(db *gorm.DB) *TodoItemRepository {
	return &TodoItemRepository{
		DB: db,
	}
}

func (r *TodoItemRepository) GetById(id uint64) (*model.ToDoItem, error) {
	var todoItem model.ToDoItem

	r.DB.First(&todoItem, id)
	r.DB.First(&todoItem.List, todoItem.ListID)

	if todoItem.ID == 0 {
		return nil, fmt.Errorf("no item in todo list table")
	}
	return &todoItem, nil
}

func (r *TodoItemRepository) GetAllByListID(id uint64) error {
	return fmt.Errorf("not implemented")
}

func (r *TodoItemRepository) Create(m *model.ToDoItem) error {
	r.DB.Create(m)
	return nil
}

func (r *TodoItemRepository) Update(m *model.ToDoItem) error {
	r.DB.Save(&m)
	return nil
}

func (r *TodoItemRepository) Delete(id uint64) error {
	var todoItem model.ToDoItem
	r.DB.Delete(&todoItem, id)
	return nil
}

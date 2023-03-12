package repository

import (
	"fmt"
	"time"

	model "github.com/riogod/todo/libs/gomodels"
	"gorm.io/gorm"
)

type TodoListRepository struct {
	DB *gorm.DB
}

type ToDoItemList struct {
	ID          uint64
	Items       *[]model.ToDoItem
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodoListRepository(db *gorm.DB) *TodoListRepository {
	return &TodoListRepository{
		DB: db,
	}
}

// get items list from database
func (r *TodoListRepository) GetAll(m *[]model.ToDoItemList) error {

	r.DB.Find(&m)
	return nil
}

// get item by id from database
func (r *TodoListRepository) GetByID(id uint64) (*ToDoItemList, error) {
	var todoItemList model.ToDoItemList

	r.DB.First(&todoItemList, id)
	if todoItemList.ID == 0 {
		return nil, fmt.Errorf("no item in todo list table")
	}

	items := []model.ToDoItem{}
	r.DB.Model(&model.ToDoItem{
		List: todoItemList,
	}).Find(&items)

	return &ToDoItemList{
		ID:          todoItemList.ID,
		Items:       &items,
		Title:       todoItemList.Title,
		Description: todoItemList.Description,
		Status:      todoItemList.Status,
		CreatedAt:   todoItemList.CreatedAt,
		UpdatedAt:   todoItemList.UpdatedAt,
	}, nil
}

// create new item in database
func (r *TodoListRepository) Create(itemModel *model.ToDoItemList) (*model.ToDoItemList, error) {

	r.DB.Create(itemModel)
	return itemModel, nil
}

// update item in database
func (r *TodoListRepository) Update(id uint64, itemModel *model.ToDoItemList) (*model.ToDoItemList, error) {
	var todoItemList model.ToDoItemList

	r.DB.First(&todoItemList, id)
	if todoItemList.ID == 0 {
		return nil, fmt.Errorf("there is no item to update with id=%v", id)
	}

	itemModel.ID = todoItemList.ID

	r.DB.Save(&itemModel)
	return itemModel, nil
}

// mark item as deleted
func (r *TodoListRepository) Delete(id uint64) error {
	var todoItemList model.ToDoItemList
	r.DB.Delete(&todoItemList, id)

	return nil
}

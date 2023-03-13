package repository

import (
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
	panic("not implemented")
}

// create new item in database
func (r *TodoListRepository) Create(itemModel *model.ToDoItemList) (*model.ToDoItemList, error) {
	panic("not implemented")
}

// update item in database
func (r *TodoListRepository) Update(id uint64, itemModel *model.ToDoItemList) (*model.ToDoItemList, error) {
	panic("not implemented")
}

// mark item as deleted
func (r *TodoListRepository) Delete(id uint64) error {
	panic("not implemented")
}

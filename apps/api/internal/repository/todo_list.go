package repository

import (
	service_error "todo_api/internal/service_errors"

	model "github.com/riogod/todo/libs/gomodels"
	"gorm.io/gorm"
)

type TodoListRepository struct {
	DB *gorm.DB
}

func NewTodoListRepository(db *gorm.DB) *TodoListRepository {
	return &TodoListRepository{
		DB: db,
	}
}

// get items list from database
func (r *TodoListRepository) GetAll() *[]model.TodoList {
	var models []model.TodoList

	r.DB.Find(&models)
	return &models
}

// get items list from database
func (r *TodoListRepository) GetAllWithItems() *[]model.TodoList {
	var models []model.TodoList

	r.DB.Preload("Items").Find(&models)
	return &models
}

// get item by id from database
func (r *TodoListRepository) GetByID(id uint64) (*model.TodoList, error) {
	var todoList model.TodoList

	r.DB.Preload("Items").Where("id = ?", id).Find(&todoList)
	if todoList.ID == 0 {
		return nil, service_error.ServiceError("NOT_FOUND", "not found item with this id")
	}
	return &todoList, nil
}

// create new item in database
func (r *TodoListRepository) Create(title string, description string, status string) (*model.TodoList, error) {

	model := model.TodoList{
		ID:          0,
		Title:       title,
		Description: description,
		Status:      status,
	}

	create := r.DB.Create(&model)
	if create.Error != nil {
		return nil, service_error.ServiceError("DB_ERROR", create.Error.Error())
	}

	return &model, nil
}

// update item in database
func (r *TodoListRepository) Update(id uint64, fields map[string]interface{}) (*model.TodoList, error) {
	var updatingListModel model.TodoList

	upd := r.DB.Model(&updatingListModel).Where("id = ?", id).Updates(fields)
	if upd.Error != nil {
		return nil, service_error.ServiceError("DB_ERROR", upd.Error.Error())
	}

	// Обновляем модель актуальными данными из базы
	var updatedModel model.TodoList
	r.DB.First(&updatedModel, id)

	return &updatingListModel, nil
}

// mark item as deleted
func (r *TodoListRepository) Delete(id uint64) error {
	err := r.DB.Delete(&model.TodoList{}, id)
	if err.Error != nil {
		return service_error.ServiceError("DB_ERROR", err.Error.Error())
	}
	return nil
}

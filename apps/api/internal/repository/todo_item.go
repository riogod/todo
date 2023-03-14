package repository

import (
	"fmt"
	"strings"
	service_error "todo_api/internal/service_errors"

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

func (r *TodoItemRepository) GetById(id uint64) (*model.TodoItem, error) {
	var todoItem model.TodoItem

	r.DB.Preload("List").First(&todoItem, id)

	if todoItem.ID == 0 {
		return nil, service_error.ServiceError("NOT_FOUND", "not found item with this id")
	}
	return &todoItem, nil
}

func (r *TodoItemRepository) GetAllBy(key string, value any) (*[]model.TodoItem, error) {
	var items []model.TodoItem

	search := r.DB.Preload("List").Where(fmt.Sprintf("%s = ?", key), value).Find(&items)
	if search.Error != nil {
		return nil, service_error.ServiceError("DB_ERROR", search.Error.Error())
	}
	return &items, nil
}

func (r *TodoItemRepository) Create(list_id uint64, title string, description string, status string) (*model.TodoItem, error) {

	list := model.TodoList{}
	findList := r.DB.Model(&model.TodoList{}).Where("id = ?", list_id).First(&list)
	if findList.Error != nil {
		return nil, service_error.ServiceError("NOT_FOUND", "list not found")
	}

	model := model.TodoItem{
		ID:          0,
		ListID:      list_id,
		Title:       title,
		Description: description,
		Status:      status,
	}

	create := r.DB.Create(&model)
	if create.Error != nil {
		return nil, service_error.ServiceError("DB_ERROR", create.Error.Error())
	}

	model.List = list
	return &model, nil
}

func (r *TodoItemRepository) Update(id uint64, fields map[string]interface{}) (*model.TodoItem, error) {

	var updatingItemModel model.TodoItem
	if r.DB.Where("id = ?", id).Preload("List").Take(&updatingItemModel).Error != nil {
		// Запись с указанным id не найдена
		return nil, service_error.ServiceError("NOT_FOUND", "item not found")
	}
	r.DB.Model(&updatingItemModel).Updates(fields)
	upd := r.DB.Model(&updatingItemModel).Updates(fields)
	if upd.Error != nil {
		return nil, service_error.ServiceError("DB_ERROR", upd.Error.Error())
	}

	// Обновляем модель актуальными данными из базы
	var updatedModel model.TodoItem
	r.DB.Preload("List").First(&updatedModel, id)

	return &updatedModel, nil
}

func (r *TodoItemRepository) Delete(id uint64) error {
	err := r.DB.Delete(&model.TodoItem{}, id)
	if err.Error != nil {
		return service_error.ServiceError("DB_ERROR", err.Error.Error())
	}
	return nil
}

func (r *TodoItemRepository) Search(title string) (*[]model.TodoItem, error) {
	var items []model.TodoItem
	search := r.DB.Preload("List").Where("title like ?", "%"+strings.ToLower(title)+"%").Find(&items)
	if search.Error != nil {
		return nil, service_error.ServiceError("DB_ERROR", search.Error.Error())
	}
	return &items, nil
}

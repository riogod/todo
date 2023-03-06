package todo_list

import (
	"fmt"

	model "github.com/riogod/todo/libs/gomodels"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

// get items list from database
func (s *Service) GetAll() (*[]model.ToDoItemList, error) {
	allList := []model.ToDoItemList{}
	s.db.Find(&allList)
	return &allList, nil
}

// get item by id from database
func (s *Service) GetByID(id string) (*model.ToDoItemList, error) {
	var todoItemList model.ToDoItemList

	s.db.First(&todoItemList, id)
	if todoItemList.ID == 0 {
		return nil, fmt.Errorf("no item in todo list table")
	}
	return &todoItemList, nil
}

// create new item in database
func (s *Service) Create(itemModel *model.ToDoItemList) (*model.ToDoItemList, error) {
	err := s.db.Create(itemModel)
	if err != nil {
		return nil, fmt.Errorf("cannot create item in todo list table")
	}
	return itemModel, nil
}

// update item in database
func (s *Service) Update(id string, itemModel *model.ToDoItemList) (*model.ToDoItemList, error) {
	var todoItemList model.ToDoItemList

	s.db.First(&todoItemList, id)
	if todoItemList.ID == 0 {
		return nil, fmt.Errorf("there is no item to update with id=%v", id)
	}

	itemModel.ID = todoItemList.ID

	s.db.Save(&itemModel)
	return itemModel, nil
}

// mark item as deleted
func (s *Service) Delete(id string) error {
	var todoItemList model.ToDoItemList
	s.db.Delete(&todoItemList, id)

	return nil
}

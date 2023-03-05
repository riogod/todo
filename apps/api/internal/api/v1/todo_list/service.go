package todo_list

import (
	"fmt"

	model "github.com/riogod/todo/libs/gomodels"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func (s *Service) Get(id string) (*model.ToDoItemList, error) {

	var todoItem model.ToDoItemList

	s.db.First(&todoItem, id)
	fmt.Println("Get!", todoItem)
	return &todoItem, nil
}

func (s *Service) Create(model *model.ToDoItemList) error {
	err := s.db.Create(model)
	if err != nil {
		return fmt.Errorf("cannot create item in database")
	}
	return nil
}

package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	TodoList *TodoListRepository
}

func InitRepositories(db *gorm.DB) *Repository {
	return &Repository{
		TodoList: NewTodoListRepository(db),
	}
}

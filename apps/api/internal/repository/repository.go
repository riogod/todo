package repository

import "gorm.io/gorm"

type Repository struct {
	TodoList *TodoListRepository
	TodoItem *TodoItemRepository
}

func Setup(db *gorm.DB) *Repository {
	return &Repository{
		TodoList: NewTodoListRepository(db),
		TodoItem: NewTodoItemRepository(db),
	}
}

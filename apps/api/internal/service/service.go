package service

import "todo_api/internal/repository"

type Service struct {
	TodoList *TodoListService
	TodoItem *TodoItemService
}

func InitServices(repo *repository.Repository) *Service {
	return &Service{
		TodoList: NewTodoListService(repo),
		TodoItem: NewTodoItemService(repo),
	}
}

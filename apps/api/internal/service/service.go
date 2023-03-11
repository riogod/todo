package service

import "todo_api/internal/repository"

type Service struct {
	TodoList *TodoListService
}

func InitServices(repo *repository.Repository) *Service {
	return &Service{
		TodoList: NewTodoListService(repo),
	}
}

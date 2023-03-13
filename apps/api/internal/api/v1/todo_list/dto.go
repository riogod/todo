package todo_list

import "time"

type TodoList struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type RequestTodoListDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ResponoseDTO struct {
	Success bool        `json:"success"`
	Body    interface{} `json:"body"`
}

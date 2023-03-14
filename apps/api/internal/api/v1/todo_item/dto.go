package todo_item

import "time"

type TodoList struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type TodoItem struct {
	ID          string    `json:"id"`
	List        TodoList  `json:"list"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type RequestTodoItemDTO struct {
	Title       string `json:"title"`
	ListID      string `json:"list_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ResponoseDTO struct {
	Success bool        `json:"success"`
	Body    interface{} `json:"body"`
}
